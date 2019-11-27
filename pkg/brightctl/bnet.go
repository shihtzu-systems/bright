package brightctl

import (
	"github.com/gorilla/sessions"
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
}

func (c BnetController) HandleBnet(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BnetPath())

	out := "Hello"
	_, _ = w.Write([]byte(out))
}
