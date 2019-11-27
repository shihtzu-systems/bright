package brightsvc

import (
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func NewTryUser(dad bool, redis tower.Redis) (out ghost.TryUser) {
	redis.Connect()
	defer redis.Disconnect()

	return ghost.TryUser{
		Name:         NewName(dad),
		MembershipId: "0",
		Characters: []bungo.Character{
			newTitan(redis),
			newHunter(redis),
			newWarlock(redis),
		},
	}
}

func newTitan(redis tower.Redis) bungo.Character {
	emblemDef := datax.GetInventoryItem(randomHash(emblemHashes), redis)
	return bungo.Character{
		Id:             "titan",
		MembershipType: 0,
		Class:          "Titan",
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
		Equipped: newOutfit("titan", redis),
		Bag:      newBag("titan", redis),
	}
}

func newHunter(redis tower.Redis) bungo.Character {
	emblemDef := datax.GetInventoryItem(randomHash(emblemHashes), redis)
	return bungo.Character{
		Id:             "hunter",
		MembershipType: 0,
		Class:          "Hunter",
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
		Equipped: newOutfit("hunter", redis),
		Bag:      newBag("hunter", redis),
	}
}

func newWarlock(redis tower.Redis) bungo.Character {
	emblemDef := datax.GetInventoryItem(randomHash(emblemHashes), redis)
	return bungo.Character{
		Id:             "warlock",
		MembershipType: 0,
		Class:          "Warlock",
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
		Equipped: newOutfit("warlock", redis),
		Bag:      newBag("warlock", redis),
	}
}

func newOutfit(characterId string, redis tower.Redis) bungo.Outfit {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(kineticWeaponHashes), func(i, j int) {
		kineticWeaponHashes[i], kineticWeaponHashes[j] = kineticWeaponHashes[j], kineticWeaponHashes[i]
	})

	kineticWeapon := newWeaponInstance(characterId, randomHash(kineticWeaponHashes), redis)
	energyWeapon := newWeaponInstance(characterId, randomHash(energyWeaponHashes), redis)
	powerWeapon := newWeaponInstance(characterId, randomHash(powerWeaponHashes), redis)

	switch characterId {
	case "warlock":
		return bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(characterId, randomHash(warlockHelmetHashes), redis),
			Gauntlets: newArmorInstance(characterId, randomHash(warlockGauntletHashes), redis),
			Chest:     newArmorInstance(characterId, randomHash(warlockChestHashes), redis),
			Leg:       newArmorInstance(characterId, randomHash(warlockLegHashes), redis),
			Class:     newArmorInstance(characterId, randomHash(warlockClassHashes), redis),
		}
	case "titan":
		return bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(characterId, randomHash(titanHelmetHashes), redis),
			Gauntlets: newArmorInstance(characterId, randomHash(titanGauntletHashes), redis),
			Chest:     newArmorInstance(characterId, randomHash(titanChestHashes), redis),
			Leg:       newArmorInstance(characterId, randomHash(titanLegHashes), redis),
			Class:     newArmorInstance(characterId, randomHash(titanClassHashes), redis),
		}
	case "hunter":
		fallthrough
	default:
		return bungo.Outfit{
			KineticWeapon: kineticWeapon,
			EnergyWeapon:  energyWeapon,
			PowerWeapon:   powerWeapon,

			Helmet:    newArmorInstance(characterId, randomHash(hunterHelmetHashes), redis),
			Gauntlets: newArmorInstance(characterId, randomHash(hunterGauntletHashes), redis),
			Chest:     newArmorInstance(characterId, randomHash(hunterChestHashes), redis),
			Leg:       newArmorInstance(characterId, randomHash(hunterLegHashes), redis),
			Class:     newArmorInstance(characterId, randomHash(hunterClassHashes), redis),
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

func newBag(characterId string, redis tower.Redis) (bag bungo.Bag) {

	weapons := bestCrucibleWeapons(characterId, redis)
	log.Debug("Bag Weapons: ", len(weapons))

	armors := hunterArmors(characterId, redis)
	log.Debug("Bag Armors: ", len(armors))

	switch characterId {
	default:
		return bungo.Bag{
			Weapons: weapons,
			Armors:  armors,
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
