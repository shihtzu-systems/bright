package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/brightview"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	rootBasePath = "/"
)

func RootPath() string {
	return rootBasePath
}

type RootController struct {
	SessionStore sessions.Store
	Tower        tower.Tower
}

func (c RootController) HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", RootPath())

	index := brightview.Index{
		Tower:        c.Tower,
		BnetAuthPath: "/bnet/auth",
		TryPath:      "/try",
		HackPath:     "/hack",
	}
	index.View(w)
}
