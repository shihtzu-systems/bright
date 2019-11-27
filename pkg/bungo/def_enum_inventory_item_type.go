package bungo

import (
	"errors"
	"fmt"
)

type InventoryItemType string

func InventoryItemTypes(key int) string {
	out, _ := InventoryItemTypesE(key)
	return out
}

func InventoryItemTypesE(key int) (string, error) {
	switch key {
	case 2:
		return "Armor", nil
	case 3:
		return "Weapon", nil
	case 20:
		return "Dummy", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
