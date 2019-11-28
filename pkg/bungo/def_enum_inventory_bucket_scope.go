package bungo

import (
	"errors"
	"fmt"
)

type InventoryBucketScope string

func InventoryBucketScopes(key int) InventoryBucketScope {
	out, _ := InventoryBucketScopesE(key)
	return out
}

func InventoryBucketScopesE(key int) (InventoryBucketScope, error) {
	switch key {
	case 0:
		return "Guardian", nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
