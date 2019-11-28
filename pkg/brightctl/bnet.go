package brightctl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/brightsvc"
	"github.com/shihtzu-systems/bright/pkg/brightview"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path"
)

const (
	bnetBasePath = "/bnet"
)

func BnetPath(pieces ...string) string {
	return path.Join(append([]string{bnetBasePath}, pieces...)...)
}

type BnetController struct {
	SessionStore sessions.Store
	Tower        tower.Tower
	BungieClient *resty.Client
}

func (c BnetController) Id(w http.ResponseWriter, r *http.Request) string {
	store, err := c.SessionStore.Get(r, c.Tower.SessionKey)
	if err != nil {
		log.Fatal(err)
	}

	name, exists := store.Values["name"]
	if !exists {
		name = generateSessionName()
		store.Values["name"] = name
	}
	if err := store.Save(r, w); err != nil {
		log.Fatal(err)
	}
	log.Debug("id: ", store.Values["name"].(string))
	return store.Values["name"].(string)
}

func (c BnetController) HandleAuth(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("auth"))

	nonce, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	authUrlRaw := fmt.Sprintf(c.Tower.BungieNetUrl+"/en/oauth/authorize?client_id=%s&response_type=code&state=%s", c.Tower.Oauth.ClientId, nonce.String())

	w.Header().Set("Location", authUrlRaw)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c BnetController) HandleCallback(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("callback"))
	vars := mux.Vars(r)

	if vars["code"] == "" {
		log.Fatal("no code found")
	}

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)
	g.Token = c.authorizationCode(vars["code"])
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", BnetPath())
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c BnetController) HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	if g.Bright.Gamer.Name == "" {
		barracks := brightsvc.Barracks{
			Tower:        c.Tower,
			Ghost:        g,
			BungieClient: c.BungieClient,
		}
		gamer := barracks.ReportForDuty()
		g.Bright.Embody(gamer)
		g.Save(c.Tower.Redis)
	}

	v := brightview.Home{
		Tower: c.Tower,
		Ghost: g,
	}
	v.View(w)
}

func (c BnetController) HandleGuardianInit(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("guardian", "init"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	barracks := brightsvc.Barracks{
		Tower:        c.Tower,
		Ghost:        g,
		BungieClient: c.BungieClient,
	}
	gamer := barracks.ReportForDuty()
	g.Bright.Embody(gamer)
	g.Shadow.Embody(gamer)
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", BnetPath("guardian", vars["membershipType"], vars["id"]))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c BnetController) HandleGuardian(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("guardian"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}

	_, exists = g.Bright.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Bright.Possess(id)
	g.Shadow.Possess(id)
	g.Save(c.Tower.Redis)

	v := brightview.Guardian{
		Tower:          c.Tower,
		Ghost:          g,
		Guardian:       g.Bright.Summon(),
		OtherGuardians: g.Bright.SummonOthers(),
	}
	v.View(w)
}

func (c BnetController) HandleSwapWeapon(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("swap", "weapon"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Bright.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Bright.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Bright.Summon()
	weaponHash, exists := vars["weaponHash"]
	if !exists {
		log.Fatal("no weapon hash")
	}

	swapOut, found := soul.Equipped.FindWeapon(weaponHash)
	if !found {
		log.Fatal("equipped weapon not found: ", weaponHash)
	}
	swapIns := soul.Bag.FindWeapons(swapOut.Slot.Name)

	v := brightview.SwapWeapon{
		Tower:          c.Tower,
		Ghost:          g,
		Guardian:       g.Bright.Summon(),
		OtherGuardians: g.Bright.SummonOthers(),
		SwapOut:        swapOut,
		SwapIns:        swapIns,
	}
	v.View(w)
}

func (c BnetController) HandleSwapArmor(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("swap", "armor"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Bright.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Bright.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Bright.Summon()
	armorHash, exists := vars["armorHash"]
	if !exists {
		log.Fatal("no armor hash")
	}

	swapOut, found := soul.Equipped.FindArmor(armorHash)
	if !found {
		log.Fatal("equipped weapon not found: ", armorHash)
	}
	swapIns := soul.Bag.FindArmors(swapOut.Slot.Name)

	v := brightview.SwapArmor{
		Tower:          c.Tower,
		Ghost:          g,
		Guardian:       g.Bright.Summon(),
		OtherGuardians: g.Bright.SummonOthers(),
		SwapOut:        swapOut,
		SwapIns:        swapIns,
	}
	v.View(w)
}

func (c BnetController) HandleSwap(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath("swap"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Bright.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Bright.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Bright.Summon()
	swapOutHash, exists := vars["swapOut"]
	if !exists {
		log.Fatal("no swap out hash")
	}

	swapInHash, exists := vars["swapIn"]
	if !exists {
		log.Fatal("no swap in hash")
	}

	swapOutWeapon, found := soul.Equipped.FindWeapon(swapOutHash)
	if found {
		swapInWeapon, found := soul.Bag.FindWeapon(swapInHash)
		if !found {
			log.Fatalf("%s not found in bag: ", swapOutHash)
		}
		log.Debugf("swapping weapon: %s -> %s", swapOutHash, swapInHash)
		soul.SwapWeapon(swapOutWeapon, swapInWeapon)
	}

	swapOutArmor, found := soul.Equipped.FindArmor(swapOutHash)
	if found {
		swapInArmor, found := soul.Bag.FindArmor(swapInHash)
		if !found {
			log.Fatalf("%s not found in bag: ", swapOutHash)
		}
		log.Debugf("swapping armor: %s -> %s", swapOutHash, swapInHash)
		soul.SwapArmor(swapOutArmor, swapInArmor)
	}

	g.Bright.Charge(soul)
	g.Save(c.Tower.Redis)

	// equip at bungie
	brightsvc.EquipItem(brightsvc.EquipItemArgs{
		InstanceId:   swapInHash,
		Tower:        c.Tower,
		Ghost:        g,
		BungieClient: c.BungieClient,
	})

	w.Header().Set("Location", BnetPath("guardian", fmt.Sprint(soul.MembershipType), soul.Id))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c BnetController) authorizationCode(code string) (out ghost.BungieToken) {

	resp, err := c.BungieClient.R().
		EnableTrace().
		SetBasicAuth(c.Tower.Oauth.ClientId, c.Tower.Oauth.ClientSecret).
		SetFormData(map[string]string{
			"grant_type": "authorization_code",
			"code":       code,
		}).
		SetContentLength(true).
		Post("/Platform/App/OAuth/Token/")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(resp.Body(), &out); err != nil {
		log.Fatal(err)
	}
	return out
}
