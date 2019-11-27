package bungo

import (
	"errors"
	"fmt"
)

type InventoryItemClassType string

func InventoryItemClassTypes(key int) InventoryItemClassType {
	out, _ := InventoryItemClassTypesE(key)
	return out
}

func InventoryItemClassTypesE(key int) (InventoryItemClassType, error) {
	switch key {
	case 0:
		return "Titan", nil
	case 1:
		return "Hunter", nil
	case 2:
		return "Warlock", nil
	case 3:
		return "Unknown", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
