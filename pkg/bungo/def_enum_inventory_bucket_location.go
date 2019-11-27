package bungo

import (
	"errors"
	"fmt"
)

type InventoryBucketLocation string

func InventoryBucketLocations(key int) InventoryBucketLocation {
	out, _ := InventoryBucketLocationsE(key)
	return out
}

func InventoryBucketLocationsE(key int) (InventoryBucketLocation, error) {
	switch key {
	case 1:
		return "Inventory", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
