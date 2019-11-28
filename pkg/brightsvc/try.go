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

type Barracks struct {
	Tower tower.Tower
}

func (b Barracks) ReportForDuty() (out bungo.Gamer) {
	b.Tower.Redis.Connect()
	defer b.Tower.Redis.Disconnect()

	return bungo.Gamer{
		Name:         NewName(b.Tower.Dad),
		MembershipId: "0",
		Guardians: []bungo.Guardian{
			b.Deploy("titan"),
			b.Deploy("hunter"),
			b.Deploy("warlock"),
		},
	}
}

func (b Barracks) Deploy(class string) bungo.Guardian {
	emblemDef := datax.GetInventoryItem(randomHash(emblemHashes), b.Tower.Redis)
	guardian := bungo.Guardian{
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
	b.Dress(&guardian)
	b.Provision(&guardian)
	return guardian
}

func (b Barracks) Dress(guardian *bungo.Guardian) {

	kineticWeapon := newWeaponInstance(guardian.Id, randomHash(kineticWeaponHashes), b.Tower.Redis)
	energyWeapon := newWeaponInstance(guardian.Id, randomHash(energyWeaponHashes), b.Tower.Redis)
	powerWeapon := newWeaponInstance(guardian.Id, randomHash(powerWeaponHashes), b.Tower.Redis)

	switch guardian.Id {
	case "warlock":
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, randomHash(warlockHelmetHashes), b.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, randomHash(warlockGauntletHashes), b.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, randomHash(warlockChestHashes), b.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, randomHash(warlockLegHashes), b.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, randomHash(warlockClassHashes), b.Tower.Redis),
		}
	case "titan":
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, randomHash(titanHelmetHashes), b.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, randomHash(titanGauntletHashes), b.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, randomHash(titanChestHashes), b.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, randomHash(titanLegHashes), b.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, randomHash(titanClassHashes), b.Tower.Redis),
		}
	case "hunter":
		fallthrough
	default:
		guardian.Equipped = bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(guardian.Id, randomHash(hunterHelmetHashes), b.Tower.Redis),
			Gauntlets: newArmorInstance(guardian.Id, randomHash(hunterGauntletHashes), b.Tower.Redis),
			Chest:     newArmorInstance(guardian.Id, randomHash(hunterChestHashes), b.Tower.Redis),
			Leg:       newArmorInstance(guardian.Id, randomHash(hunterLegHashes), b.Tower.Redis),
			Class:     newArmorInstance(guardian.Id, randomHash(hunterClassHashes), b.Tower.Redis),
		}
	}
}

func randomHash(hashes []int) int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hashes), func(i, j int) {
		hashes[i], hashes[j] = hashes[j], hashes[i]
	})
	return hashes[0]
}

func (b Barracks) Provision(guardian *bungo.Guardian) {
	switch guardian.Id {
	case "warlock":
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, b.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, b.Tower.Redis),
		}
	case "titan":
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, b.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, b.Tower.Redis),
		}
	case "hunter":
		fallthrough
	default:
		guardian.Bag = bungo.Bag{
			Weapons: bestCrucibleWeapons(guardian.Id, b.Tower.Redis),
			Armors:  hunterArmors(guardian.Id, b.Tower.Redis),
		}
	}
}

func newWeaponInstance(characterId string, weaponHash int, redis tower.Redis) bungo.Weapon {
	var weaponBucket string
	var weaponSlot string
	var weaponType string
	itemDef := datax.GetInventoryItem(weaponHash, redis)
	for _, itemCategoryHash := range itemDef.ItemCategoryHashes {
		itemCategoryDef := datax.GetItemCategory(itemCategoryHash, redis)
		if itemCategoryHash == kineticCategory ||
			itemCategoryHash == energyCategory ||
			itemCategoryHash == powerCategory {
			weaponBucket = itemCategoryDef.ShortTitle
			weaponSlot = itemCategoryDef.ShortTitle
		}
		if itemCategoryHash == autoRifleCategory ||
			itemCategoryHash == handCannonCategory ||
			itemCategoryHash == pulseRifleCategory ||
			itemCategoryHash == scoutRifleCategory ||
			itemCategoryHash == fusionRifleCategory ||
			itemCategoryHash == sniperRifleCategory ||
			itemCategoryHash == shotgunCategory ||
			itemCategoryHash == machineGunCategory ||
			itemCategoryHash == rocketLauncherCategory ||
			itemCategoryHash == sidearmCategory ||
			itemCategoryHash == swordCategory ||
			itemCategoryHash == grenadeLauncherCategory ||
			itemCategoryHash == tracerRifleCategory ||
			itemCategoryHash == bowCategory {
			weaponType = itemCategoryDef.ShortTitle
		}
	}
	return bungo.Weapon{
		MembershipType:  0,
		CharacterId:     characterId,
		InstanceId:      fmt.Sprintf("instance-of-%d", weaponHash),
		InventoryItemId: weaponHash,

		Tier:        itemDef.Inventory.TierTypeName,
		Bucket:      weaponBucket,
		Name:        itemDef.DisplayProperties.Name,
		Type:        weaponType,
		IconEnabled: itemDef.DisplayProperties.HasIcon,
		IconUri:     itemDef.DisplayProperties.Icon,
		Slot: bungo.Slot{
			Name: weaponSlot,
		},
	}
}

func newArmorInstance(characterId string, armorHash int, redis tower.Redis) bungo.Armor {
	var armorClass string
	itemDef := datax.GetInventoryItem(armorHash, redis)
	itemBucketDef := datax.GetInventoryBucket(itemDef.Inventory.BucketTypeHash, redis)
	for _, itemCategoryHash := range itemDef.ItemCategoryHashes {
		itemCategoryDef := datax.GetItemCategory(itemCategoryHash, redis)
		if itemCategoryHash == warlockCategory ||
			itemCategoryHash == titanCategory ||
			itemCategoryHash == hunterCategory {
			armorClass = itemCategoryDef.ShortTitle
		}
	}
	armorBucket := itemBucketDef.DisplayProperties.Name
	armorSlot := itemBucketDef.DisplayProperties.Name
	return bungo.Armor{
		MembershipType:  0,
		CharacterId:     characterId,
		InstanceId:      fmt.Sprintf("instance-of-%d", armorHash),
		InventoryItemId: armorHash,

		Tier:        itemDef.Inventory.TierTypeName,
		Bucket:      armorBucket,
		Name:        itemDef.DisplayProperties.Name,
		Type:        armorClass,
		IconEnabled: itemDef.DisplayProperties.HasIcon,
		IconUri:     itemDef.DisplayProperties.Icon,
		Slot: bungo.Slot{
			Name: armorSlot,
		},
	}
}

func bestCrucibleWeapons(characterId string, redis tower.Redis) (out []bungo.Weapon) {
	for _, hash := range bestCrucibleWeaponHashes {
		instance := newWeaponInstance(characterId, hash, redis)
		out = append(out, instance)
	}
	return out
}

func hunterArmors(characterId string, redis tower.Redis) (out []bungo.Armor) {
	for _, hash := range hunterArmorHashes {
		instance := newArmorInstance(characterId, hash, redis)
		out = append(out, instance)
	}
	return out
}
