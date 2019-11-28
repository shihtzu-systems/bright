package brightview

import (
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
)

var (
	guardianPath = "guardian.html"
)

type Guardian struct {
	Tower          tower.Tower
	Ghost          ghost.Ghost
	Guardian       bungo.Guardian
	OtherGuardians []bungo.Guardian
}

func (v Guardian) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, navbarPartPath),
		path.Join(v.Tower.TemplatesPath, guardianPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "guardian", v); err != nil {
		log.Fatal(err)
	}
}
