package bungo

import (
	"errors"
	"fmt"
)

type BungieMembershipType string

const (
	BungieMembershipTypeNoneKey = 0
	BungieMembershipTypeNone    = "None"

	BungieMembershipTypeTigerXboxKey = 1
	BungieMembershipTypeTigerXbox    = "TigerXbox"

	BungieMembershipTypeTigerPsnKey = 2
	BungieMembershipTypeTigerPsn    = "TigerPsn"

	BungieMembershipTypeTigerSteamKey = 3
	BungieMembershipTypeTigerSteam    = "TigerSteam"

	BungieMembershipTypeTigerBlizzardKey = 4
	BungieMembershipTypeTigerBlizzard    = "TigerBlizzard"

	BungieMembershipTypeTigerStadiaKey = 5
	BungieMembershipTypeTigerStadia    = "TigerStadia"

	BungieMembershipTypeTigerDemonKey = 10
	BungieMembershipTypeTigerDemon    = "TigerDemon"

	BungieMembershipTypeBungieNextKey = 254
	BungieMembershipTypeBungieNext    = "BungieNext"

	BungieMembershipTypeAllKey = -1
	BungieMembershipTypeAll    = "All"
)

func BungieMembershipTypes(key int) BungieMembershipType {
	out, _ := BungieMembershipTypesE(key)
	return out
}

func BungieMembershipTypesE(key int) (BungieMembershipType, error) {
	switch key {
	case BungieMembershipTypeNoneKey:
		return BungieMembershipTypeNone, nil
	case BungieMembershipTypeTigerXboxKey:
		return BungieMembershipTypeTigerXbox, nil
	case BungieMembershipTypeTigerPsnKey:
		return BungieMembershipTypeTigerPsn, nil
	case BungieMembershipTypeTigerSteamKey:
		return BungieMembershipTypeTigerSteam, nil
	case BungieMembershipTypeTigerBlizzardKey:
		return BungieMembershipTypeTigerBlizzard, nil
	case BungieMembershipTypeTigerStadiaKey:
		return BungieMembershipTypeTigerStadia, nil
	case BungieMembershipTypeTigerDemonKey:
		return BungieMembershipTypeTigerDemon, nil
	case BungieMembershipTypeBungieNextKey:
		return BungieMembershipTypeBungieNext, nil
	case BungieMembershipTypeAllKey:
		return BungieMembershipTypeAll, nil
	default:
		return "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
