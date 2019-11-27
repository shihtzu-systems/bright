package bungo

import (
	"errors"
	"fmt"
)

type InventoryBucketCategory string

func InventoryBucketCategories(key int) InventoryBucketCategory {
	out, _ := InventoryBucketCategoriesE(key)
	return out
}

func InventoryBucketCategoriesE(key int) (InventoryBucketCategory, error) {
	switch key {
	case 3:
		return "Equippable", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
