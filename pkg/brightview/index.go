package brightview

import (
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
)

const (
	bungieUrl    = "https://www.bungie.net"
	bungieApiUri = "/Platform"
)

var (
	// parts
	headerPartPath = "parts/header.html"
	footerPartPath = "parts/footer.html"
	navbarPartPath = "parts/navbar.html"

	indexPath = "index.html"
)

type Index struct {
	Tower tower.Tower
}

func (v Index) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, indexPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "index", struct {
		BnetAuthPath string
		TryPath      string
		HackPath     string
		Tower        tower.Tower
	}{
		"/bnet/auth",
		"/try",
		"/hack",
		v.Tower,
	}); err != nil {
		log.Fatal(err)
	}
}
