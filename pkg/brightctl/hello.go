package brightctl

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"path"
)

const (
	helloBasePath = "/hello"
)

func HelloPath(pieces ...string) string {
	return path.Join(append([]string{helloBasePath}, pieces...)...)
}

type HelloController struct{}

func (c HelloController) HandleHello(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", HelloPath())
	w.WriteHeader(http.StatusOK)
}
