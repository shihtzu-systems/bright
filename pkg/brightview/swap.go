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
	swapPath = "swap.html"
)

type SwapWeapon struct {
	Tower          tower.Tower
	Ghost          ghost.Ghost
	Guardian       bungo.Guardian
	OtherGuardians []bungo.Guardian
	SwapOut        bungo.Weapon
	SwapIns        []bungo.Weapon
}

func (v SwapWeapon) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, navbarPartPath),
		path.Join(v.Tower.TemplatesPath, swapPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "swap", v); err != nil {
		log.Fatal(err)
	}
}

type SwapArmor struct {
	Tower          tower.Tower
	Ghost          ghost.Ghost
	Guardian       bungo.Guardian
	OtherGuardians []bungo.Guardian
	SwapOut        bungo.Armor
	SwapIns        []bungo.Armor
}

func (v SwapArmor) View(w http.ResponseWriter) {
	tpl, err := template.ParseFiles(
		path.Join(v.Tower.TemplatesPath, headerPartPath),
		path.Join(v.Tower.TemplatesPath, navbarPartPath),
		path.Join(v.Tower.TemplatesPath, swapPath),
		path.Join(v.Tower.TemplatesPath, footerPartPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := tpl.ExecuteTemplate(w, "swap", v); err != nil {
		log.Fatal(err)
	}
}
