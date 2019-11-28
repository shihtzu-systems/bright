package bright

import (
	"bytes"
	"errors"
	"fmt"
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"

	"github.com/shihtzu-systems/bright/generated/bungie/data"
)

type ServerArgs struct {
	Serial        string
	SessionKey    string
	SessionSecret []byte

	OathClientId     string
	OathClientSecret string
	OathRedirectUrl  string

	BungieClient *resty.Client
	Destiny      data.Content

	RedisAddress string
	RedisPort    string

	SystemName string
	HackMode   bool
	BnetMode   bool
	TryMode    bool

	Trace          bool
	Debug          bool
	Terminator     bool
	DadModifier    bool
	MayhemModifier bool

	HackToken BungieToken

	GoogleAnalyticsId string

	WorkingBasePath string
}

var myName string
var sessionStore sessions.Store
var randomGen = rand.New(rand.NewSource(time.Now().UnixNano()))
var redisAddress = "localhost"
var redisPort = "6379"
var serial = "unknown"
var systemName = "unknown"
var sessionKey = "unknown"

const (
	bungieUrl    = "https://www.bungie.net"
	bungieApiUri = "/Platform"

	// parts
	actionPath = "templates/parts/action.html"
	headerPath = "templates/parts/header.html"
	footerPath = "templates/parts/footer.html"
	navPath    = "templates/parts/nav.html"

	characterPath = "templates/character.html"
	homePath      = "templates/home.html"
	indexPath     = "templates/index.html"
	swapPath      = "templates/swap.html"
)

var templates = []string{
	// parts
	actionPath,
	headerPath,
	footerPath,
	navPath,

	characterPath,
	homePath,
	indexPath,
	swapPath,
}

func Server(args ServerArgs) error {

	r := mux.NewRouter()

	sessionKey = args.SessionKey
	sessionStore = sessions.NewCookieStore(args.SessionSecret)

	redisAddress = args.RedisAddress
	redisPort = args.RedisPort

	systemName = args.SystemName
	serial = args.Serial

	destinyContentPath, err := DestinyContentPath(args.WorkingBasePath)
	if err != nil {
		return err
	}

	tower := Tower{
		Hack: args.HackMode,
		Bnet: args.BnetMode,
		Try:  args.TryMode,

		Trace:      args.Trace,
		Debug:      args.Debug,
		Terminator: args.Terminator,
		Dad:        args.DadModifier,
		Mayhem:     args.MayhemModifier,

		RedisAddress: redisAddress,
		RedisPort:    redisPort,

		GoogleAnalyticsId: args.GoogleAnalyticsId,

		DestinyContentPath: destinyContentPath,
	}
	if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
		return err
	}

	if tower.Trace {
		log.Print("setting log level to trace")
		log.SetLevel(log.TraceLevel)
	} else if tower.Debug {
		log.Print("setting log level to debug")
		log.SetLevel(log.DebugLevel)
	}

	// hack mode
	if tower.Hack {
		r.HandleFunc("/hack", func(w http.ResponseWriter, r *http.Request) {

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost.Token = args.HackToken

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/bnet", w)
		})

		r.HandleFunc("/hack/bnet", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Bnet = !tower.Bnet
			log.Info("bnet = ", tower.Bnet)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})
		r.HandleFunc("/hack/try", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Try = !tower.Try
			log.Info("try = ", tower.Try)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})
		r.HandleFunc("/hack/dad", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Dad = !tower.Dad
			log.Info("dad = ", tower.Dad)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})
		r.HandleFunc("/hack/mayhem", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Mayhem = !tower.Mayhem
			log.Info("mayhem = ", tower.Mayhem)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})
		r.HandleFunc("/hack/debug", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Debug = !tower.Debug
			log.Info("debug = ", tower.Debug)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})
		r.HandleFunc("/hack/trace", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}
			tower.Trace = !tower.Trace
			log.Info("trace = ", tower.Trace)
			if err := saveTower(redisAddress+":"+redisPort, systemName, serial, tower); err != nil {
				errRedirect(err, "saveTower", w)
				return
			}
			redirect("/", w)
		})

		r.HandleFunc("/hack/ghost", func(w http.ResponseWriter, r *http.Request) {

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			spew.Dump(ghost)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/", w)
		})

		r.HandleFunc("/hack/tower", func(w http.ResponseWriter, r *http.Request) {
			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}

			spew.Dump(tower)

			redirect("/", w)
		})
	}

	// bnet mode - auth
	if tower.Bnet {

		r.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
			nonce, err := uuid.NewRandom()
			if err != nil {
				errRedirect(err, "uuid.NewRandom", w)
				return
			}
			authUrlRaw := fmt.Sprintf(bungieUrl+"/en/oauth/authorize?client_id=%s&response_type=code&state=%s", args.OathClientId, nonce.String())

			redirect(authUrlRaw, w)
		})

		r.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost.Token, err = bungieToken(args.BungieClient, args.OathClientId, args.OathClientSecret, vars["code"])
			if err != nil {
				errRedirect(err, "bungieToken", w)
				return
			}

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/bnet", w)

		}).Queries("code", "{code}", "state", "{state}")

		// home

		r.HandleFunc("/bnet", func(w http.ResponseWriter, r *http.Request) {

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			currentUser, err := CurrentUser(CurrentUserArgs{
				OathClientId:     args.OathClientId,
				OathClientSecret: args.OathClientSecret,
				Auth:             ghost.Token,
				BungieClient:     args.BungieClient,
				Destiny:          args.Destiny,
			})
			if err != nil {
				errRedirect(err, "Gamer", w)
				return
			}

			ghost.User = currentUser
			ghost.BnetOne = bungo.Guardian{}
			ghost.BnetTwo = bungo.Guardian{}
			ghost.BnetThree = bungo.Guardian{}
			ghost.BrightOne = bungo.Guardian{}
			ghost.BrightTwo = bungo.Guardian{}
			ghost.BrightThree = bungo.Guardian{}

			if len(currentUser.Guardians) >= 1 {
				ghost.BnetOne = currentUser.Guardians[0]
				if ghost.BrightOne.Id == "" {
					ghost.BrightOne = currentUser.Guardians[0]
				}
			}
			if len(currentUser.Guardians) >= 2 {
				ghost.BnetTwo = currentUser.Guardians[1]
				if ghost.BrightTwo.Id == "" {
					ghost.BrightTwo = currentUser.Guardians[1]
				}
			}
			if len(currentUser.Guardians) == 3 {
				ghost.BnetThree = currentUser.Guardians[2]
				if ghost.BrightThree.Id == "" {
					ghost.BrightThree = currentUser.Guardians[2]
				}
			}

			ghost = ghost.MakeCurrent(currentUser.Guardians[0])

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "home", struct {
				DisplayName       string
				BungieUrl         string
				CharacterBaseUri  string
				Characters        []bungo.Guardian
				AppVersion        string
				GoogleAnalyticsId string
			}{
				ghost.User.Name,
				bungieUrl,
				"/bnet",
				ghost.RetrieveAllBrightCharacters(),
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		// character

		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}
			if ghost.User.Name == "" {
				redirect("/bnet", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "character", struct {
				Mayhem                bool
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				TakenDiffs            int
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				tower.Mayhem,
				bungieUrl,
				"/bnet",
				ghost.User.Name,
				ghost.RetrieveCurrentBright(),
				ghost.RetrieveOtherBrightCharacters(),
				ghost.RetrieveCurrentBnet(),
				ghost.RetrieveOtherBnetCharacters(),
				len(ghost.RetrieveCurrentBright().Differences(ghost.RetrieveCurrentBnet())),
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}/swap/armor/{instanceId:[a-z-0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			character := ghost.RetrieveCurrentBright()
			if ghost.User.Name == "" && character.Id != "" {
				redirect("/bnet", w)
				return
			}

			swapOut, _ := character.Equipped.FindArmor(vars["instanceId"])
			if swapOut.InstanceId != "" {
				log.Debug("Found: ", swapOut.InstanceId)
			} else {
				log.Debug("Not found: ", vars["instanceId"])
				errRedirect(errors.New("unknown instance of inventory item"), "FindWeapon", w)
				return
			}

			swapIns := character.Bag.FindArmors(swapOut.Slot.Name)

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "swap", struct {
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				SwapType              string
				SwapOut               bungo.Armor
				SwapIns               []bungo.Armor
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				bungieUrl,
				"/bnet",
				ghost.User.Name,
				ghost.RetrieveCurrentBright(),
				ghost.RetrieveOtherBrightCharacters(),
				ghost.RetrieveCurrentBnet(),
				ghost.RetrieveOtherBnetCharacters(),
				swapOut.SwapType(),
				swapOut,
				swapIns,
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}/swap/weapon/{instanceId:[a-z-0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			currentCharacter := ghost.RetrieveCurrentBright()
			if ghost.User.Name == "" && currentCharacter.Id != "" {
				redirect("/bnet", w)
				return
			}

			swapOut, _ := currentCharacter.Equipped.FindWeapon(vars["instanceId"])
			if swapOut.InstanceId != "" {
				log.Debug("Found: ", swapOut.InstanceId)
			} else {
				log.Debug("Not found: ", vars["instanceId"])
				errRedirect(errors.New("unknown instance of inventory item"), "FindWeapon", w)
				return
			}

			swapIns := currentCharacter.Bag.FindWeapons(swapOut.Slot.Name)

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "swap", struct {
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				SwapType              string
				SwapOut               bungo.Weapon
				SwapIns               []bungo.Weapon
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				bungieUrl,
				"/bnet",
				ghost.User.Name,
				ghost.RetrieveCurrentBright(),
				ghost.RetrieveOtherBrightCharacters(),
				ghost.RetrieveCurrentBnet(),
				ghost.RetrieveOtherBnetCharacters(),
				swapOut.SwapType(),
				swapOut,
				swapIns,
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}/swap/{swapType}/{swapIn:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			membershipType, err := strconv.Atoi(vars["membershipType"])
			if err != nil {
				errRedirect(err, "strconv.Atoi", w)
				return
			}

			currentCharacter := ghost.RetrieveCurrentBright()
			if ghost.User.Name == "" && currentCharacter.Id != "" {
				redirect("/bnet", w)
				return
			}

			//switch vars["swapType"] {
			//case "kinetic":
			//	swapOut := currentCharacter.Equipped.KineticWeapon
			//	swapIn, bag := currentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	currentCharacter.Equipped.KineticWeapon = swapIn
			//	currentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "energy":
			//	swapOut := currentCharacter.Equipped.EnergyWeapon
			//	swapIn, bag := currentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	currentCharacter.Equipped.EnergyWeapon = swapIn
			//	currentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "power":
			//	swapOut := currentCharacter.Equipped.PowerWeapon
			//	swapIn, bag := currentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	currentCharacter.Equipped.PowerWeapon = swapIn
			//	currentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "helmet":
			//	swapOut := currentCharacter.Equipped.Helmet
			//	swapIn, bag := currentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	currentCharacter.Equipped.Helmet = swapIn
			//	currentCharacter.Bag = bag.StowArmor(swapOut)
			//case "gauntlets":
			//	swapOut := currentCharacter.Equipped.Gauntlets
			//	swapIn, bag := currentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	currentCharacter.Equipped.Gauntlets = swapIn
			//	currentCharacter.Bag = bag.StowArmor(swapOut)
			//case "chest":
			//	swapOut := currentCharacter.Equipped.Chest
			//	swapIn, bag := currentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	currentCharacter.Equipped.Chest = swapIn
			//	currentCharacter.Bag = bag.StowArmor(swapOut)
			//case "legs":
			//	swapOut := currentCharacter.Equipped.Leg
			//	swapIn, bag := currentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	currentCharacter.Equipped.Leg = swapIn
			//	currentCharacter.Bag = bag.StowArmor(swapOut)
			//case "class":
			//	swapOut := currentCharacter.Equipped.Class
			//	swapIn, bag := currentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	currentCharacter.Equipped.Class = swapIn
			//	currentCharacter.Bag = bag.StowArmor(swapOut)
			//default:
			//	log.Warn("unknown swapType: ", vars["swapType"])
			//}

			ghost = ghost.BrightSave(currentCharacter)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			// equip at bnet
			if err := EquipItem(EquipItemArgs{
				MembershipType: membershipType,
				CharacterId:    vars["characterId"],
				InstanceId:     vars["swapIn"],

				OathClientId:     args.OathClientId,
				OathClientSecret: args.OathClientSecret,
				Auth:             ghost.Token,
				BungieClient:     args.BungieClient,
				Destiny:          args.Destiny,
			}); err != nil {
				errRedirect(err, "EquipItem", w)
				return
			}

			redirect("/bnet/character/"+vars["membershipType"]+"/"+vars["characterId"], w)
		})

		// taken refresh
		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}/taken/refresh", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			membershipType, err := strconv.Atoi(vars["membershipType"])
			if err != nil {
				errRedirect(err, "strconv.Atoi", w)
				return
			}

			currentCharacter, err := CurrentCharacter(CurrentCharacterArgs{
				MembershipType: membershipType,
				Id:             vars["characterId"],

				OathClientId:     args.OathClientId,
				OathClientSecret: args.OathClientSecret,
				Auth:             ghost.Token,
				BungieClient:     args.BungieClient,
				Destiny:          args.Destiny,
			})
			if err != nil {
				errRedirect(err, "CurrentCharacter", w)
				return
			}
			ghost = ghost.BnetSave(currentCharacter)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/bnet/character/"+vars["membershipType"]+"/"+vars["characterId"], w)
		})

		r.HandleFunc("/bnet/character/{membershipType:[0-9]+}/{characterId:[0-9]+}/init", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost = ghost.MakeCurrentId(vars["characterId"])

			membershipType, err := strconv.Atoi(vars["membershipType"])
			if err != nil {
				errRedirect(err, "strconv.Atoi", w)
				return
			}

			currentCharacter, err := CurrentCharacter(CurrentCharacterArgs{
				MembershipType: membershipType,
				Id:             vars["characterId"],

				OathClientId:     args.OathClientId,
				OathClientSecret: args.OathClientSecret,
				Auth:             ghost.Token,
				BungieClient:     args.BungieClient,
				Destiny:          args.Destiny,
			})
			if err != nil {
				errRedirect(err, "CurrentCharacter", w)
				return
			}
			ghost = ghost.BnetSave(currentCharacter)
			ghost = ghost.BrightSave(currentCharacter)
			ghost = ghost.MakeCurrent(currentCharacter)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/bnet/character/"+vars["membershipType"]+"/"+vars["characterId"], w)
		})
	}

	// try mode
	if tower.Try {

		r.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}

			if ghost.TryUser.Name == "" {
				tryUser, err := TryUser(TryUserArgs{
					Destiny: args.Destiny,
					Dad:     tower.Dad,
				})
				if err != nil {
					errRedirect(err, "TryUser", w)
					return
				}

				ghost.TryUser = bungo.Gamer{
					Name:         tryUser.Name,
					MembershipId: tryUser.MembershipId,
				}

				if len(tryUser.Guardians) >= 1 {
					ghost.TryOne = tryUser.Guardians[0]
				}
				if len(tryUser.Guardians) >= 2 {
					ghost.TryTwo = tryUser.Guardians[1]
				}
				if len(tryUser.Guardians) == 3 {
					ghost.TryThree = tryUser.Guardians[2]
				}
				ghost = ghost.MakeTryCurrent(ghost.TryOne)
			}

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "home", struct {
				BungieUrl         string
				CharacterBaseUri  string
				DisplayName       string
				Characters        []bungo.Guardian
				AppVersion        string
				GoogleAnalyticsId string
			}{
				bungieUrl,
				"/try",
				ghost.TryUser.Name,
				ghost.RetrieveAllTryCharacters(),
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/try/recycle", func(w http.ResponseWriter, r *http.Request) {

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}

			ghost.TryUser = bungo.Gamer{}
			ghost.TryCurrentId = ""
			ghost.TryOne = bungo.Guardian{}
			ghost.TryTwo = bungo.Guardian{}
			ghost.TryThree = bungo.Guardian{}

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/try", w)
		})

		r.HandleFunc("/try/character/0/{characterId:[a-z]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}
			if ghost.TryUser.Name == "" {
				redirect("/try", w)
				return
			}

			ghostCharacter := ghost.FindTryCharacter(vars["characterId"])
			ghost = ghost.MakeTryCurrent(ghostCharacter)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
			if err != nil {
				errRedirect(err, "retrieveTower", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "character", struct {
				Mayhem                bool
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				TakenDiffs            int
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				tower.Mayhem,
				bungieUrl,
				"/try",
				ghost.TryUser.Name,
				ghost.RetrieveTryCurrentCharacter(),
				ghost.RetrieveTryOtherCharacters(),
				ghost.RetrieveTryCurrentCharacter(),
				ghost.RetrieveTryOtherCharacters(),
				0,
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/try/character/0/{characterId:[a-z]+}/swap/armor/{instanceId:[a-z-0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}
			tryCurrentCharacter := ghost.RetrieveTryCurrentCharacter()
			if ghost.TryUser.Name == "" && tryCurrentCharacter.Id != "" {
				redirect("/try", w)
				return
			}

			swapOut, _ := tryCurrentCharacter.Equipped.FindArmor(vars["instanceId"])
			if swapOut.InstanceId != "" {
				log.Debug("Found: ", swapOut.InstanceId)
			} else {
				log.Debug("Not found: ", vars["instanceId"])
				errRedirect(errors.New("unknown instance of inventory item"), "FindWeapon", w)
				return
			}

			swapIns := tryCurrentCharacter.Bag.FindArmors(swapOut.Slot.Name)

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "swap", struct {
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				SwapType              string
				SwapOut               bungo.Armor
				SwapIns               []bungo.Armor
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				bungieUrl,
				"/try",
				ghost.TryUser.Name,
				tryCurrentCharacter,
				ghost.RetrieveTryOtherCharacters(),
				tryCurrentCharacter,
				ghost.RetrieveTryOtherCharacters(),
				swapOut.SwapType(),
				swapOut,
				swapIns,
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/try/character/0/{characterId:[a-z]+}/swap/weapon/{instanceId:[a-z-0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}
			tryCurrentCharacter := ghost.RetrieveTryCurrentCharacter()
			if ghost.TryUser.Name == "" && tryCurrentCharacter.Id != "" {
				redirect("/try", w)
				return
			}

			swapOut, _ := tryCurrentCharacter.Equipped.FindWeapon(vars["instanceId"])
			if swapOut.InstanceId != "" {
				log.Debug("Found: ", swapOut.InstanceId)
			} else {
				log.Debug("Not found: ", vars["instanceId"])
				errRedirect(errors.New("unknown instance of inventory item"), "FindWeapon", w)
				return
			}

			swapIns := tryCurrentCharacter.Bag.FindWeapons(swapOut.Slot.Name)

			tpl, err := template.ParseFiles(templates...)
			if err != nil {
				errRedirect(err, "template.ParseFiles", w)
				return
			}

			if err := tpl.ExecuteTemplate(w, "swap", struct {
				BungieUrl             string
				CharacterBaseUri      string
				DisplayName           string
				BrightCharacter       bungo.Guardian
				BrightOtherCharacters []bungo.Guardian
				TakenCharacter        bungo.Guardian
				TakenOtherCharacters  []bungo.Guardian
				SwapType              string
				SwapOut               bungo.Weapon
				SwapIns               []bungo.Weapon
				AppVersion            string
				GoogleAnalyticsId     string
			}{
				bungieUrl,
				"/try",
				ghost.TryUser.Name,
				tryCurrentCharacter,
				ghost.RetrieveTryOtherCharacters(),
				tryCurrentCharacter,
				ghost.RetrieveTryOtherCharacters(),
				swapOut.SwapType(),
				swapOut,
				swapIns,
				args.Serial,
				args.GoogleAnalyticsId,
			}); err != nil {
				errRedirect(err, "tpl.ExecuteTemplate", w)
				return
			}
		})

		r.HandleFunc("/try/character/0/{characterId:[a-z]+}/swap/{swapType}/{swapIn:[a-z-0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
			if err != nil {
				errRedirect(err, "retrieveGhost", w)
				return
			}
			tryCurrentCharacter := ghost.RetrieveTryCurrentCharacter()
			if ghost.TryUser.Name == "" && tryCurrentCharacter.Id != "" {
				redirect("/try", w)
				return
			}
			//
			//switch vars["swapType"] {
			//case "kinetic":
			//	swapOut := tryCurrentCharacter.Equipped.KineticWeapon
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.KineticWeapon = swapIn
			//	tryCurrentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "energy":
			//	swapOut := tryCurrentCharacter.Equipped.EnergyWeapon
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.EnergyWeapon = swapIn
			//	tryCurrentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "power":
			//	swapOut := tryCurrentCharacter.Equipped.PowerWeapon
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeWeapon(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.PowerWeapon = swapIn
			//	tryCurrentCharacter.Bag = bag.StowWeapon(swapOut)
			//case "helmet":
			//	swapOut := tryCurrentCharacter.Equipped.Helmet
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.Helmet = swapIn
			//	tryCurrentCharacter.Bag = bag.StowArmor(swapOut)
			//case "gauntlets":
			//	swapOut := tryCurrentCharacter.Equipped.Gauntlets
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.Gauntlets = swapIn
			//	tryCurrentCharacter.Bag = bag.StowArmor(swapOut)
			//case "chest":
			//	swapOut := tryCurrentCharacter.Equipped.Chest
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.Chest = swapIn
			//	tryCurrentCharacter.Bag = bag.StowArmor(swapOut)
			//case "legs":
			//	swapOut := tryCurrentCharacter.Equipped.Leg
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.Leg = swapIn
			//	tryCurrentCharacter.Bag = bag.StowArmor(swapOut)
			//case "class":
			//	swapOut := tryCurrentCharacter.Equipped.Class
			//	swapIn, bag := tryCurrentCharacter.Bag.TakeArmor(vars["swapIn"])
			//
			//	tryCurrentCharacter.Equipped.Class = swapIn
			//	tryCurrentCharacter.Bag = bag.StowArmor(swapOut)
			//default:
			//	log.Warn("unknown swapType: ", vars["swapType"])
			//}

			ghost = ghost.TrySave(tryCurrentCharacter)
			ghost = ghost.MakeTryCurrent(tryCurrentCharacter)

			if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
				errRedirect(err, "saveGhost", w)
				return
			}

			redirect("/try/character/0/"+vars["characterId"], w)
		})

		r.HandleFunc("/try/character/0/{characterId:[a-z]+}/init", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			log.Debug("try init")
			redirect("/try/character/0/"+vars["characterId"], w)
		})
	}

	// health

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		if err := redisPing(args.RedisAddress); err != nil {
			errRedirect(err, "redisPing", w)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	// err

	r.HandleFunc("/err", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	// logout

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {

		ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
		if err != nil {
			errRedirect(err, "retrieveGhost", w)
			return
		}

		ghost = Ghost{}

		if err := saveGhost(redisAddress+":"+redisPort, systemName, id(w, r), ghost); err != nil {
			errRedirect(err, "saveGhost", w)
			return
		}
		redirect("/", w)
	})

	// root

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sess, err := sessionStore.Get(r, sessionKey)
		if err != nil {
			errRedirect(err, "sessionStore.Get", w)
			return
		}

		if sess.Values["name"] == "" {
			sess.Values["name"] = GenerateName(args.DadModifier)
		}

		if err := sess.Save(r, w); err != nil {
			errRedirect(err, "sessions.Save", w)
			return
		}

		tower, err := retrieveTower(redisAddress+":"+redisPort, systemName, serial)
		if err != nil {
			errRedirect(err, "retrieveTower", w)
			return
		}

		tpl, err := template.ParseFiles(templates...)
		if err != nil {
			errRedirect(err, "template.ParseFiles", w)
			return
		}
		if err := tpl.ExecuteTemplate(w, "index", struct {
			HackMode          bool
			BnetMode          bool
			TryMode           bool
			AppVersion        string
			GoogleAnalyticsId string
		}{
			tower.Hack && args.HackToken.AccessToken != "",
			tower.Bnet,
			tower.Try,
			args.Serial,
			args.GoogleAnalyticsId,
		}); err != nil {
			errRedirect(err, "tpl.ExecuteTemplate", w)
			return
		}
	})

	// static
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static/"))))

	r.Use(loggingMiddleware)
	r.Use(authenticationMiddleware)
	name := haikunator.New()
	name.Delimiter = " "
	name.TokenLength = 0
	myName = name.Haikunate()
	log.Printf("starting v%s as %s", args.Serial, myName)
	log.Printf("listening on localhost:8080")
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down ", myName)
	os.Exit(0)
	return nil
}

type DisplayCharacter struct {
	Class      string
	Level      string
	Light      string
	EmblemIcon string
}

type BungieToken struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	MembershipId     string `json:"membership_id"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// don't log the health checker
		if !strings.HasPrefix(r.URL.Path, "/static") &&
			r.UserAgent() != "ELB-HealthChecker/2.0" {
			log.Debug(r.URL.Path)
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// only apply to paths that start with /bnet/
		if !strings.HasPrefix(r.RequestURI, "/bnet/") {
			next.ServeHTTP(w, r)
			return
		}

		ghost, err := retrieveGhost(redisAddress+":"+redisPort, systemName, id(w, r))
		if err != nil {
			errRedirect(err, "retrieveGhost", w)
			return
		}

		if ghost.Token.MembershipId != "" {
			log.Printf("Authenticated user %s", ghost.Token.MembershipId)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

type Tower struct {
	Hack bool
	Bnet bool
	Try  bool

	Trace      bool
	Debug      bool
	Terminator bool
	Dad        bool
	Mayhem     bool

	RedisAddress string
	RedisPort    string

	GoogleAnalyticsId string

	DestinyContentPath string
}

type Ghost struct {
	User bungo.Gamer

	CurrentId string

	BnetOne   bungo.Guardian
	BnetTwo   bungo.Guardian
	BnetThree bungo.Guardian

	BrightOne   bungo.Guardian
	BrightTwo   bungo.Guardian
	BrightThree bungo.Guardian

	TryUser      bungo.Gamer
	TryCurrentId string
	TryOne       bungo.Guardian
	TryTwo       bungo.Guardian
	TryThree     bungo.Guardian

	Token BungieToken
}

func (g Ghost) CurrentDiff() int {
	log.Debug("search for current: ", g.CurrentId)
	bnet := bungo.Guardian{}
	switch g.CurrentId {
	case g.BnetOne.Id:
		bnet = g.BnetOne
	case g.BnetTwo.Id:
		bnet = g.BnetTwo
	case g.BnetThree.Id:
		bnet = g.BnetThree
	}

	bright := bungo.Guardian{}
	switch g.CurrentId {
	case g.BrightOne.Id:
		bright = g.BrightOne
	case g.BrightTwo.Id:
		bnet = g.BrightTwo
	case g.BrightThree.Id:
		bnet = g.BrightThree
	}
	return len(bright.Differences(bnet))
}

func (g Ghost) MakeCurrent(character bungo.Guardian) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	if character.Id == "" {
		log.Fatal("character must have value")
	}

	out.CurrentId = character.Id
	return out
}

func (g Ghost) MakeCurrentId(id string) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	if id == "" {
		log.Fatal("id must have value")
	}

	out.CurrentId = id
	return out
}

func (g Ghost) BrightSave(character bungo.Guardian) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	switch character.Id {
	case g.BrightOne.Id:
		log.Debugf("bright one save: %s [%s]", character.Id, character.Class)
		out.BrightOne = character
	case g.BrightTwo.Id:
		log.Debugf("bright two save: %s [%s]", character.Id, character.Class)
		out.BrightTwo = character
	case g.BrightThree.Id:
		log.Debugf("bright three save: %s [%s]", character.Id, character.Class)
		out.BrightThree = character
	default:
		log.Warnf("bright not found: %s [%s] in %s [%s], %s [%s], %s [%s]",
			character.Id, character.Class, g.BrightOne.Id, g.BrightOne.Class, g.BrightTwo.Id, g.BrightTwo.Class, g.BrightThree.Id, g.BrightThree.Class)
	}
	return out
}

func (g Ghost) BnetSave(character bungo.Guardian) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	switch character.Id {
	case g.BnetOne.Id:
		log.Debugf("bnet one save: %s [%s]", character.Id, character.Class)
		out.BnetOne = character
	case g.BnetTwo.Id:
		log.Debugf("bnet two save: %s [%s]", character.Id, character.Class)
		out.BnetTwo = character
	case g.BnetThree.Id:
		log.Debugf("bnet three save: %s [%s]", character.Id, character.Class)
		out.BnetThree = character
	default:
		log.Warnf("bnet not found: %s [%s] in  %s [%s], %s [%s], %s [%s]",
			character.Id, character.Class, g.BnetOne.Id, g.BnetOne.Class, g.BnetTwo.Id, g.BnetTwo.Class, g.BnetThree.Id, g.BnetThree.Class)
	}
	return out
}

func (g Ghost) RetrieveCurrentBnet() bungo.Guardian {
	log.Debug("search for current: ", g.CurrentId)
	switch g.CurrentId {
	case g.BnetOne.Id:
		log.Debug("found: ", g.BnetOne.Id)
		return g.BnetOne
	case g.BnetTwo.Id:
		log.Debug("found: ", g.BnetTwo.Id)
		return g.BnetTwo
	case g.BnetThree.Id:
		log.Debug("found: ", g.BnetThree.Id)
		return g.BnetThree
	default:
		log.Debug("not found: ", g.CurrentId)
		return bungo.Guardian{}
	}
}

func (g Ghost) RetrieveCurrentBright() bungo.Guardian {
	log.Debug("search for current: ", g.CurrentId)
	switch g.CurrentId {
	case g.BrightOne.Id:
		log.Debug("found: ", g.BrightOne.Id)
		return g.BrightOne
	case g.BrightTwo.Id:
		log.Debug("found: ", g.BrightTwo.Id)
		return g.BrightTwo
	case g.BrightThree.Id:
		log.Debug("found: ", g.BrightThree.Id)
		return g.BrightThree
	default:
		log.Debug("not found: ", g.CurrentId)
		return bungo.Guardian{}
	}
}

func (g Ghost) FindBnetCharacter(characterId string) (out bungo.Guardian) {
	log.Debug("search for: ", characterId)
	switch characterId {
	case g.BnetOne.Id:
		log.Debug("found: ", characterId)
		return g.BnetOne
	case g.BnetTwo.Id:
		log.Debug("found: ", characterId)
		return g.TryTwo
	case g.BnetThree.Id:
		log.Debug("found: ", characterId)
		return g.BnetThree
	default:
		return bungo.Guardian{}
	}
}

func (g Ghost) FindBrightCharacter(characterId string) (out bungo.Guardian) {
	log.Debug("search for: ", characterId)
	switch characterId {
	case g.BrightOne.Id:
		log.Debug("found: ", characterId)
		return g.BrightOne
	case g.BrightTwo.Id:
		log.Debug("found: ", characterId)
		return g.TryTwo
	case g.BrightThree.Id:
		log.Debug("found: ", characterId)
		return g.BrightThree
	default:
		return bungo.Guardian{}
	}
}

func (g Ghost) RetrieveAllBnetCharacters() (out []bungo.Guardian) {
	return []bungo.Guardian{
		g.BnetOne,
		g.BnetTwo,
		g.BnetThree,
	}
}

func (g Ghost) RetrieveAllBrightCharacters() (out []bungo.Guardian) {
	return []bungo.Guardian{
		g.BrightOne,
		g.BrightTwo,
		g.BrightThree,
	}
}

func (g Ghost) RetrieveOtherBnetCharacters() (out []bungo.Guardian) {
	log.Trace("current character is: ", g.CurrentId)
	if g.BnetOne.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BnetOne.Id)
		out = append(out, g.BnetOne)
	}
	if g.BnetTwo.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BnetTwo.Id)
		out = append(out, g.BnetTwo)
	}
	if g.BnetThree.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BnetThree.Id)
		out = append(out, g.BnetThree)
	}
	log.Trace("other characters: ", len(out))
	return out
}

func (g Ghost) RetrieveOtherBrightCharacters() (out []bungo.Guardian) {
	log.Trace("current character is: ", g.CurrentId)
	if g.BrightOne.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BrightOne.Id)
		out = append(out, g.BrightOne)
	}
	if g.BrightTwo.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BrightTwo.Id)
		out = append(out, g.BrightTwo)
	}
	if g.BrightThree.Id != g.CurrentId {
		log.Trace("adding other character: ", g.BrightThree.Id)
		out = append(out, g.BrightThree)
	}
	log.Trace("other characters: ", len(out))
	return out
}

func (g Ghost) MakeTryCurrent(character bungo.Guardian) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	out.TryCurrentId = character.Id
	return out
}

func (g Ghost) TrySave(character bungo.Guardian) (out Ghost) {
	// blind copy
	_ = copier.Copy(&out, &g)

	switch character.Id {
	case g.TryOne.Id:
		out.TryOne = character
	case g.TryTwo.Id:
		out.TryTwo = character
	case g.TryThree.Id:
		out.TryThree = character
	}
	return out
}

func (g Ghost) RetrieveTryCurrentCharacter() bungo.Guardian {
	log.Trace("search for current: ", g.TryCurrentId)
	switch g.TryCurrentId {
	case g.TryOne.Id:
		log.Trace("found: ", g.TryOne.Id)
		return g.TryOne
	case g.TryTwo.Id:
		log.Trace("found: ", g.TryTwo.Id)
		return g.TryTwo
	case g.TryThree.Id:
		log.Trace("found: ", g.TryThree.Id)
		return g.TryThree
	default:
		log.Trace("not found: ", g.TryCurrentId)
		return bungo.Guardian{}
	}
}

func (g Ghost) FindTryCharacter(characterId string) (out bungo.Guardian) {
	log.Trace("search for: ", characterId)
	switch characterId {
	case g.TryOne.Id:
		log.Trace("found: ", characterId)
		return g.TryOne
	case g.TryTwo.Id:
		log.Trace("found: ", characterId)
		return g.TryTwo
	case g.TryThree.Id:
		log.Trace("found: ", characterId)
		return g.TryThree
	default:
		return bungo.Guardian{}
	}
}

func (g Ghost) RetrieveAllTryCharacters() (out []bungo.Guardian) {
	return []bungo.Guardian{
		g.TryOne,
		g.TryTwo,
		g.TryThree,
	}
}

func (g Ghost) RetrieveTryOtherCharacters() (out []bungo.Guardian) {
	log.Trace("current character is: ", g.TryCurrentId)
	if g.TryOne.Id != g.TryCurrentId {
		log.Trace("adding other character: ", g.TryOne.Id)
		out = append(out, g.TryOne)
	}
	if g.TryTwo.Id != g.TryCurrentId {
		log.Trace("adding other character: ", g.TryTwo.Id)
		out = append(out, g.TryTwo)
	}
	if g.TryThree.Id != g.TryCurrentId {
		log.Trace("adding other character: ", g.TryThree.Id)
		out = append(out, g.TryThree)
	}
	log.Trace("other characters: ", len(out))
	return out
}

func redirect(location string, w http.ResponseWriter) {
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func errRedirect(err error, functionName string, w http.ResponseWriter) {
	log.Error(err)
	w.Header().Set("Location", "/err?function="+functionName)
	w.WriteHeader(http.StatusInternalServerError)
}

func bungieToken(rClient *resty.Client, oauthClientId, oauthSecret, code string) (out BungieToken, err error) {

	resp, err := rClient.R().
		EnableTrace().
		SetBasicAuth(oauthClientId, oauthSecret).
		SetFormData(map[string]string{
			"grant_type": "authorization_code",
			"code":       code,
		}).
		SetContentLength(true).
		Post(bungieUrl + "/Platform/App/OAuth/Token/")
	if err != nil {
		return BungieToken{}, err
	}
	if err := json.Unmarshal(resp.Body(), &out); err != nil {
		return BungieToken{}, err
	}
	return out, nil
}

func refresh(rClient *resty.Client, oauthClientId, oauthSecret string, token BungieToken) (out BungieToken, err error) {
	log.Debug("refresh token: ", token.RefreshToken)
	resp, err := rClient.R().
		EnableTrace().
		SetBasicAuth(oauthClientId, oauthSecret).
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": token.RefreshToken,
		}).
		SetContentLength(true).
		Post(bungieUrl + "/Platform/App/OAuth/Token/")
	if err != nil {
		return BungieToken{}, err
	}
	log.Debug("refresh response: ", string(resp.Body()))
	if err := json.Unmarshal(resp.Body(), &out); err != nil {
		return BungieToken{}, err
	}
	return out, nil
}

func redisConnect(address string) (out redis.Conn, err error) {
	log.Tracef("dialing redis %s", address)
	return redis.Dial("tcp", address)
}

func redisPing(address string) (err error) {
	log.Tracef("ping redis at %s", address)
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		return err
	}
	reply, err := conn.Do("ping")
	if err != nil {
		return err
	}
	if reply.(string) != "PONG" {
		return errors.New("unexpected response from redis: " + reply.(string))
	}
	log.Trace("ping ok")
	return nil
}

func redisDisconnect(out redis.Conn) {
	log.Trace("closing redis")
	_ = out.Close()
}

func ghostExists(address, systemName, seessionKey string) (out bool, err error) {
	cache, err := redisConnect(address)
	if err != nil {
		return false, err
	}
	defer redisDisconnect(cache)

	log.Debugf("[get] | %s", ghostKey(systemName, seessionKey))
	ghost, err := cache.Do("GET", ghostKey(systemName, seessionKey))
	found := ghost != nil

	log.Debugf("[get] | found -> %v", found)
	return found, nil
}

func retrieveGhost(address, systemName, seessionKey string) (out Ghost, err error) {
	cache, err := redisConnect(address)
	if err != nil {
		return Ghost{}, err
	}
	defer redisDisconnect(cache)

	log.Debugf("[get] | %s", ghostKey(systemName, seessionKey))

	ghost, err := cache.Do("GET", ghostKey(systemName, seessionKey))
	if err != nil {
		return Ghost{}, err
	}
	if ghost != nil {
		if err := json.Unmarshal(ghost.([]byte), &out); err != nil {
			return Ghost{}, err
		}
	}

	log.Trace("[get] !!! ghost.TryCurrentCharacter.Equipped.KineticWeapon")
	log.Trace("[get] !!! ", ghostKey(systemName, seessionKey))
	log.Trace("[get] !!! try -> current.equipped.kinetic = ", out.RetrieveTryCurrentCharacter().Equipped.KineticWeapon.Name)
	return out, nil
}

func saveGhost(address, systemName, sessionKey string, ghost Ghost) (err error) {
	cache, err := redisConnect(address)
	if err != nil {
		return err
	}
	defer redisDisconnect(cache)

	gout, err := json.Marshal(ghost)
	if err != nil {
		return err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, gout, "", "\t"); err != nil {
		return err
	}

	// ghost
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey), prettyJSON.Bytes()); err != nil {
		return err
	}
	// user
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "user"), ghost.User.Name); err != nil {
		return err
	}
	// current
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "current"), ghost.CurrentId); err != nil {
		return err
	}
	// bnet one
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bnet", "one"), ghost.BnetOne.String()); err != nil {
		return err
	}
	// bnet two
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bnet", "two"), ghost.BnetTwo.String()); err != nil {
		return err
	}
	// bnet three
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bnet", "three"), ghost.BnetThree.String()); err != nil {
		return err
	}
	// bright one
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bright", "one"), ghost.BrightOne.String()); err != nil {
		return err
	}
	// bright two
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bright", "two"), ghost.BrightTwo.String()); err != nil {
		return err
	}
	// bright three
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "bright", "three"), ghost.BrightThree.String()); err != nil {
		return err
	}

	// try user
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "try", "user"), ghost.TryUser.Name); err != nil {
		return err
	}
	// try current
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "try", "current"), ghost.TryCurrentId); err != nil {
		return err
	}
	// try one
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "try", "one"), ghost.TryOne.String()); err != nil {
		return err
	}
	// try two
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "try", "two"), ghost.TryTwo.String()); err != nil {
		return err
	}
	// try three
	if _, err = cache.Do("SET", ghostKey(systemName, sessionKey, "try", "three"), ghost.TryThree.String()); err != nil {
		return err
	}
	return nil
}

func ghostKey(systemName, sessionId string, pieces ...string) string {
	if len(pieces) > 0 {
		return fmt.Sprintf("%s:%s:ghost:%s", systemName, sessionId, strings.Join(pieces, ":"))
	} else {
		return fmt.Sprintf("%s:%s:ghost", systemName, sessionId)
	}
}

func retrieveTower(address, systemName, serial string) (out Tower, err error) {
	cache, err := redisConnect(address)
	if err != nil {
		return Tower{}, err
	}
	defer redisDisconnect(cache)

	log.Debugf("[get] | %s", towerKey(serial, systemName))

	tower, err := cache.Do("GET", towerKey(serial, systemName))
	if err != nil {
		return Tower{}, err
	}
	if tower != nil {
		if err := json.Unmarshal(tower.([]byte), &out); err != nil {
			return Tower{}, err
		}
	}

	log.Trace("[get] !!! ", towerKey(serial, systemName))
	return out, nil
}

func saveTower(address, systemName, serial string, tower Tower) (err error) {
	cache, err := redisConnect(address)
	if err != nil {
		return err
	}
	defer redisDisconnect(cache)

	tout, err := json.Marshal(tower)
	if err != nil {
		return err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, tout, "", "\t"); err != nil {
		return err
	}
	// tower
	if _, err = cache.Do("SET", towerKey(serial, systemName), prettyJSON.Bytes()); err != nil {
		return err
	}
	// debug
	if _, err = cache.Do("SET", towerKey(serial, systemName, "debug"), tower.Debug); err != nil {
		return err
	}
	// trace
	if _, err = cache.Do("SET", towerKey(serial, systemName, "trace"), tower.Trace); err != nil {
		return err
	}
	// terminator
	if _, err = cache.Do("SET", towerKey(serial, systemName, "terminator"), tower.Terminator); err != nil {
		return err
	}
	// hack
	if _, err = cache.Do("SET", towerKey(serial, systemName, "hack"), tower.Hack); err != nil {
		return err
	}
	// bnet
	if _, err = cache.Do("SET", towerKey(serial, systemName, "bnet"), tower.Bnet); err != nil {
		return err
	}
	// try
	if _, err = cache.Do("SET", towerKey(serial, systemName, "try"), tower.Try); err != nil {
		return err
	}
	// dad
	if _, err = cache.Do("SET", towerKey(serial, systemName, "dad"), tower.Dad); err != nil {
		return err
	}
	// mayhem
	if _, err = cache.Do("SET", towerKey(serial, systemName, "mayhem"), tower.Dad); err != nil {
		return err
	}
	return nil
}

func towerKey(serial, systemName string, pieces ...string) string {
	if len(pieces) > 0 {
		return fmt.Sprintf("%s:%s:tower:%s", systemName, serial, strings.Join(pieces, ":"))
	} else {
		return fmt.Sprintf("%s:%s:tower", systemName, serial)
	}
}

func id(w http.ResponseWriter, r *http.Request) string {
	sess, err := sessionStore.Get(r, sessionKey)
	if err != nil {
		log.Fatal(err)
	}

	if sess.Values["name"] == nil {
		name := GenerateName(false)
		log.Debug("setting name to ", name)
		sess.Values["name"] = name
	}

	if err := sess.Save(r, w); err != nil {
		log.Fatal(err)
	}
	log.Debug("id: ", sess.Values["name"].(string))
	return sess.Values["name"].(string)
}
