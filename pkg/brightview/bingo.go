package brightview

import (
	"github.com/shihtzu-systems/bright/pkg/bingo"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
)

var (
	bingoPath = "bingo.html"
)

type Bingo struct {
	Board bingo.Board
	Tower tower.Tower
}

func (v Bingo) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, bingoPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "bingo", v); err != nil {
		log.Fatal(err)
	}
}
