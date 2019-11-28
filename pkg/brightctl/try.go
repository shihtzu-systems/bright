package brightctl

import (
	"fmt"
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
	tryBasePath = "/try"
)

func TryPath(pieces ...string) string {
	return path.Join(append([]string{tryBasePath}, pieces...)...)
}

type TryController struct {
	SessionStore sessions.Store
	Tower        tower.Tower
}

func (c TryController) Id(w http.ResponseWriter, r *http.Request) string {
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

func (c TryController) HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	if g.Try.Gamer.Name == "" {
		cantina := brightsvc.Cantina{Tower: c.Tower}
		tryUser := cantina.ReportForDuty()
		g.Try.Embody(tryUser)
		g.Save(c.Tower.Redis)
	}

	v := brightview.Home{
		Tower: c.Tower,
		Ghost: g,
	}
	v.TryView(w)
}

func (c TryController) HandleRecycle(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	cantina := brightsvc.Cantina{Tower: c.Tower}
	tryUser := cantina.ReportForDuty()
	g.Try.Embody(tryUser)
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", TryPath())
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c TryController) HandleGuardian(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath("guardian"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Try.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Try.Possess(id)
	g.Save(c.Tower.Redis)

	v := brightview.Guardian{
		Tower:          c.Tower,
		Ghost:          g,
		Guardian:       g.Try.Summon(),
		OtherGuardians: g.Try.SummonOthers(),
	}
	v.TryView(w)
}

func (c TryController) HandleSwapWeapon(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath("swap", "weapon"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Try.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Try.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Try.Summon()
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
		Guardian:       g.Try.Summon(),
		OtherGuardians: g.Try.SummonOthers(),
		SwapOut:        swapOut,
		SwapIns:        swapIns,
	}
	v.TryView(w)
}

func (c TryController) HandleSwapArmor(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath("swap", "armor"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Try.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Try.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Try.Summon()
	armorHash, exists := vars["armorHash"]
	if !exists {
		log.Fatal("no weapon hash")
	}

	swapOut, found := soul.Equipped.FindArmor(armorHash)
	if !found {
		log.Fatal("equipped armor not found: ", armorHash)
	}
	swapIns := soul.Bag.FindArmors(swapOut.Slot.Name)

	v := brightview.SwapArmor{
		Tower:          c.Tower,
		Ghost:          g,
		Guardian:       g.Try.Summon(),
		OtherGuardians: g.Try.SummonOthers(),
		SwapOut:        swapOut,
		SwapIns:        swapIns,
	}
	v.TryView(w)
}

func (c TryController) HandleSwap(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath("swap"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	id, exists := vars["id"]
	if !exists {
		log.Fatal("no id")
	}
	g.Materialize(c.Tower.Redis)

	_, exists = g.Try.Call(id)
	if !exists {
		log.Fatal("not found: " + id)
	}
	g.Try.Possess(id)
	g.Save(c.Tower.Redis)

	soul := g.Try.Summon()
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

	g.Try.Charge(soul)
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", TryPath("guardian", fmt.Sprint(soul.MembershipType), soul.Id))
	w.WriteHeader(http.StatusTemporaryRedirect)
}
