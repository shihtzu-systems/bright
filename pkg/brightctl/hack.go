package brightctl

import (
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path"
)

const (
	hackBasePath = "/hack"
)

func HackPath(pieces ...string) string {
	return path.Join(append([]string{hackBasePath}, pieces...)...)
}

type HackController struct {
	SessionStore sessions.Store
	Tower        tower.Tower
	HackToken    tower.HackToken
}

func (c HackController) Id(w http.ResponseWriter, r *http.Request) string {
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

func (c HackController) HandleHack(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", HackPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)
	g.Token = ghost.BungieToken(c.HackToken)
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", path.Join(bnetBasePath))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func generateSessionName() string {
	namer := haikunator.New()
	namer.TokenLength = 6
	return namer.Haikunate()
}
