package brightsvc

import (
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"math/rand"
	"strings"
	"time"
)

type Cantina struct {
	Tower tower.Tower
}

func (c Cantina) ReportForDuty() (out bungo.User) {
	c.Tower.Redis.Connect()
	defer c.Tower.Redis.Disconnect()

	return bungo.User{
		Name:         NewName(c.Tower.Dad),
		MembershipId: "0",
		Characters: []bungo.Character{
			c.Deploy("titan"),
			c.Deploy("hunter"),
			c.Deploy("warlock"),
		},
	}
}

func (c Cantina) Deploy(class string) bungo.Character {
	emblemDef := datax.GetInventoryItem(randomHash(emblemHashes), c.Tower.Redis)
	guardian := bungo.Character{
		Id:             strings.ToLower(class),
		MembershipType: 0,
		Class:          strings.Title(class),
		Level:          50,
		Light:          750 + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(170),
		Emblem: bungo.Emblem{
			Name:            emblemDef.DisplayProperties.Name,
			InventoryItemId: emblemDef.Hash,
			IconUri:         emblemDef.DisplayProperties.Icon,
			BannerUri:       emblemDef.SecondarySpecial,
			IconOverlayUri:  emblemDef.SecondaryOverlay,
			Background: bungo.Color{
				Red:   emblemDef.BackgroundColor.Red,
				Green: emblemDef.BackgroundColor.Green,
				Blue:  emblemDef.BackgroundColor.Blue,
				Alpha: emblemDef.BackgroundColor.Alpha,
			},
		},
	}
	c.Dress(&guardian)
	c.Provision(&guardian)
	return guardian
}

func (c Cantina) Dress(guardian *bungo.Character) {

	kineticWeaponHash := randomHash(kineticWeaponHashes)
	kineticWeapon := newWeaponInstance(guardian.Id, fmt.Sprintf("instance-of-%d", kineticWeaponHash), kineticWeaponHash, c.Tower.Redis)

	energyWeaponHash := randomHash(energyWeaponHashes)
	energyWeapon := newWeaponInstance(guardian.Id, fmt.Sprintf("instance-of-%d", energyWeaponHash), energyWeaponHash, c.Tower.Redis)

	powerWeaponnHash := randomHash(powerWeaponHashes)
	powerWeapon := newWeaponInstance(guardian.Id, fmt.Sprintf("instance-of-%d", powerWeaponnHash), powerWeaponnHash, c.Tower.Redis)

	switch guardian.Id {
	case "warlock":
		helmetHash := randomHash(warlockHelmetHashes)
		gauntletsHash := randomHash(warlockGauntletHashes)
		chestHash := randomHash(warlockChestHashes)
		legHash := randomHash(warlockLegHashes)
		classHash := randomHash(warlockClassHashes)
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", helmetHash), helmetHash, c.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", gauntletsHash), gauntletsHash, c.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", chestHash), chestHash, c.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", legHash), legHash, c.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", classHash), classHash, c.Tower.Redis),
		}
	case "titan":
		helmetHash := randomHash(titanHelmetHashes)
		gauntletsHash := randomHash(titanGauntletHashes)
		chestHash := randomHash(titanChestHashes)
		legHash := randomHash(titanLegHashes)
		classHash := randomHash(titanClassHashes)
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", helmetHash), helmetHash, c.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", gauntletsHash), gauntletsHash, c.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", chestHash), chestHash, c.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", legHash), legHash, c.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", classHash), classHash, c.Tower.Redis),
		}
	case "hunter":
		fallthrough
	default:
		helmetHash := randomHash(hunterHelmetHashes)
		gauntletsHash := randomHash(hunterGauntletHashes)
		chestHash := randomHash(hunterChestHashes)
		legHash := randomHash(hunterLegHashes)
		classHash := randomHash(hunterClassHashes)
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", helmetHash), helmetHash, c.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", gauntletsHash), gauntletsHash, c.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", chestHash), chestHash, c.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", legHash), legHash, c.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, fmt.Sprintf("instance-of-%d", classHash), classHash, c.Tower.Redis),
		}
	}
}

func (c Cantina) Provision(guardian *bungo.Character) {
	switch guardian.Id {
	case "warlock":
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, c.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, c.Tower.Redis),
		}
	case "titan":
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, c.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, c.Tower.Redis),
		}
	case "hunter":
		fallthrough
	default:
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, c.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, c.Tower.Redis),
		}
	}
}
