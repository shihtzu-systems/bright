package brightsvc

import (
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/tower"
)

func newArmorInstance(characterId, instanceId string, armorHash int, redis tower.Redis) bungo.Armor {
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
		InstanceId:      instanceId,
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

func hunterArmors(characterId string, redis tower.Redis) (out []bungo.Armor) {
	for _, hash := range hunterArmorHashes {
		instance := newArmorInstance(characterId, fmt.Sprintf("instance-of-%d", hash), hash, redis)
		out = append(out, instance)
	}
	return out
}
