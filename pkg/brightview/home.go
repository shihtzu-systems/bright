package brightview

import (
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
)

var (
	homePath = "home.html"
)

type Home struct {
	Tower   tower.Tower
	Ghost   ghost.Ghost
	TryMode bool
}

func (v Home) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, navbarPartPath),
		path.Join(v.Tower.TemplatesPath, homePath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "home", v); err != nil {
		log.Fatal(err)
	}
}
