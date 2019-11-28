package brightsvc

import (
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/tower"
)

func newWeaponInstance(characterId, instanceId string, weaponHash int, redis tower.Redis) bungo.Weapon {
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
		InstanceId:      instanceId,
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

func bestCrucibleWeapons(characterId string, redis tower.Redis) (out []bungo.Weapon) {
	for _, hash := range bestCrucibleWeaponHashes {
		instance := newWeaponInstance(characterId, fmt.Sprintf("instance-of-%d", hash), hash, redis)
		out = append(out, instance)
	}
	return out
}
