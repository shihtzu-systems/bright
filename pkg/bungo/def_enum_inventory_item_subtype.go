package bungo

import (
	"errors"
	"fmt"
)

type InventoryItemSubtype string

func InventoryItemSubtypes(key int) string {
	out, _ := InventoryItemSubtypesE(key)
	return out
}

func InventoryItemSubtypesE(key int) (string, error) {
	switch key {
	case 0:
		return "None", nil
	case 9:
		return "HandCannon", nil
	case 27:
		return "GauntletsArmor", nil
	case 28:
		return "ChestArmor", nil
	case 29:
		return "LegArmor", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
