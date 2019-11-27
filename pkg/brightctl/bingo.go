package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/brightsvc"
	"github.com/shihtzu-systems/bright/pkg/brightview"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path"
)

const (
	bingoBasePath = "/bingo"
)

func BingoPath(pieces ...string) string {
	return path.Join(append([]string{bingoBasePath}, pieces...)...)
}

type BingoController struct {
	SessionStore sessions.Store
	Tower        tower.Tower
}

func (c BingoController) HandleBingo(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling ", BingoPath())
	board := brightsvc.NewBingo("christmas")
	bingo := brightview.Bingo{
		Board: board,
		Tower: c.Tower,
	}
	bingo.View(w)
}
