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
	characterPath = "home.html"
)

type Character struct {
	Tower        tower.Tower
	TryCharacter ghost.TryCharacter
}

func (v Character) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, navbarPartPath),
		path.Join(v.Tower.TemplatesPath, characterPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "character", struct {
		Tower tower.Tower
	}{
		v.Tower,
	}); err != nil {
		log.Fatal(err)
	}
}
