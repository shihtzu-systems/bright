package brightctl

import (
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

func (c TryController) HandleTry(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)
	if g.Try.User.Name == "" {
		tryUser := brightsvc.NewTryUser(c.Tower.Dad)
		g.SetTryUser(tryUser)
		g.SetTryCurrent(tryUser.Characters[0])
		g.Save(c.Tower.Redis)
	}

	home := brightview.Home{
		Tower:   c.Tower,
		Ghost:   g,
		TryMode: true,
	}
	home.View(w)
}

func (c TryController) HandleTryRecycle(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath())

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	g.Materialize(c.Tower.Redis)

	tryUser := brightsvc.NewTryUser(c.Tower.Dad)
	g.SetTryUser(tryUser)
	g.SetTryCurrent(tryUser.Characters[0])
	g.Save(c.Tower.Redis)

	w.Header().Set("Location", TryPath())
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c TryController) HandleTryCharacter(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", TryPath("character"))
	vars := mux.Vars(r)

	g := ghost.NewGhost(c.Tower.System.Id, c.Id(w, r))
	characterId, exists := vars["characterId"]
	if !exists {
		log.Fatal("no character id")
	}
	character := brightview.Character{
		Tower:        c.Tower,
		TryCharacter: g.FindTry(characterId),
	}
	character.View(w)
}
