package bungo

import (
	"errors"
	"fmt"
)

type InventoryItemTierType string

func InventoryItemTierTypeTypes(key int) InventoryItemTierType {
	out, _ := InventoryItemTierTypesE(key)
	return out
}

func InventoryItemTierTypesE(key int) (InventoryItemTierType, error) {
	switch key {
	case 0:
		return "Unknown", nil
	case 1:
		return "??????", nil
	case 2:
		return "Basic", nil
	case 3:
		return "Common", nil
	case 4:
		return "Rare", nil
	case 5:
		return "Superior", nil
	case 6:
		return "Exotic", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
