package bright

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	log "github.com/sirupsen/logrus"
	"strings"
)

type CurrentUserArgs struct {
	OathClientId     string
	OathClientSecret string
	Auth             BungieToken
	BungieClient     *resty.Client
	Destiny          data.Content
}

func CurrentUser(args CurrentUserArgs) (out bungo.CurrentUser, err error) {
	freshToken, err := refresh(args.BungieClient, args.OathClientId, args.OathClientSecret, args.Auth)
	if err != nil {
		return bungo.CurrentUser{}, err
	}
	log.Debug(freshToken.MembershipId)

	getMembershipDataForCurrentUserResponse, err := args.BungieClient.R().
		EnableTrace().
		SetAuthToken(freshToken.AccessToken).
		Get(bungieApiUri + "/User/GetMembershipsForCurrentUser/")
	if err != nil {
		return bungo.CurrentUser{}, err
	}

	var getMembershipDataForCurrentUser bungo.GetMembershipDataForCurrentUser
	if err := json.Unmarshal(getMembershipDataForCurrentUserResponse.Body(), &getMembershipDataForCurrentUser); err != nil {
		return bungo.CurrentUser{}, err
	}

	bungieNetUser := getMembershipDataForCurrentUser.Response.BungieNetUser
	destinyMemberships := getMembershipDataForCurrentUser.Response.DestinyMemberships

	destiny := args.Destiny
	log.Debug("/User/GetMembershipsForCurrentUser -> ", bungieNetUser.DisplayName)
	var characters []bungo.Character
	for _, destinyMembership := range destinyMemberships {

		getProfileResp, err := args.BungieClient.R().
			EnableTrace().
			SetAuthToken(freshToken.AccessToken).
			SetQueryParams(map[string]string{
				"components": strings.Join(bungo.ComponentTypesManyS(
					bungo.ComponentsProfilesKey,
				), ","),
			}).
			SetPathParams(map[string]string{
				"membershipType":      fmt.Sprint(destinyMembership.MembershipType),
				"destinyMembershipId": destinyMembership.MembershipID,
			}).
			Get(bungieApiUri + "/Destiny2/{membershipType}/Profile/{destinyMembershipId}/")
		if err != nil {
			return bungo.CurrentUser{}, err
		}

		getProfileRespRaw := getProfileResp.Body()

		var getProfile bungo.GetProfile
		if err := json.Unmarshal(getProfileRespRaw, &getProfile); err != nil {
			return bungo.CurrentUser{}, err
		}

		profileData := getProfile.Response.Profile.Data

		for _, characterId := range profileData.CharacterIds {
			getCharacterResp, err := args.BungieClient.R().
				EnableTrace().
				SetAuthToken(freshToken.AccessToken).
				SetQueryParams(map[string]string{
					"components": strings.Join(bungo.ComponentTypesManyS(
						bungo.ComponentsCharactersKey,
						bungo.ComponentsCharacterEquipmentKey,
						bungo.ComponentsCharacterInventoriesKey,
					), ","),
				}).
				SetPathParams(map[string]string{
					"membershipType":      fmt.Sprint(destinyMembership.MembershipType),
					"destinyMembershipId": destinyMembership.MembershipID,
					"characterId":         characterId,
				}).
				Get(bungieApiUri + "/Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/")
			if err != nil {
				return bungo.CurrentUser{}, err
			}

			getCharacterRespRaw := getCharacterResp.Body()

			var getCharacter bungo.GetCharacter
			if err := json.Unmarshal(getCharacterRespRaw, &getCharacter); err != nil {
				return bungo.CurrentUser{}, err
			}

			characterData := getCharacter.Response.Character.Data
			equipmentItemsData := getCharacter.Response.Equipment.Data
			inventoryItemsData := getCharacter.Response.Inventory.Data

			log.Debug("/Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/ -> ", characterData.CharacterID)
			var equipped bungo.Outfit
			for _, item := range equipmentItemsData.Items {

				itemDef := destiny.InventoryItem.Find(int(item.ItemHash))
				bucketDef := destiny.InventoryBucket.Find(itemDef.Inventory.BucketTypeHash)

				switch bucketDef.DisplayProperties.Name {
				case "Kinetic Weapons":
					log.Debug("Kinetic: ", itemDef.DisplayProperties.Name)
					equipped.KineticWeapon = NewWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Energy Weapons":
					log.Debug("Energy: ", itemDef.DisplayProperties.Name)
					equipped.EnergyWeapon = NewWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Power Weapons":
					log.Debug("Power: ", itemDef.DisplayProperties.Name)
					equipped.PowerWeapon = NewWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Helmet":
					log.Debug("Helmet: ", itemDef.DisplayProperties.Name)
					equipped.Helmet = NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Gauntlets":
					log.Debug("Gauntlets: ", itemDef.DisplayProperties.Name)
					equipped.Gauntlets = NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Chest Armor":
					log.Debug("Chest: ", itemDef.DisplayProperties.Name)
					equipped.Chest = NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Leg Armor":
					log.Debug("Legs: ", itemDef.DisplayProperties.Name)
					equipped.Leg = NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				case "Class Armor":
					log.Debug("Class: ", itemDef.DisplayProperties.Name)
					equipped.Class = NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny)
				default:
					log.Trace("Ignoring: ", bucketDef.DisplayProperties.Name)
				}
			}

			var bag bungo.Bag
			for _, item := range inventoryItemsData.Items {

				itemDef := destiny.InventoryItem.Find(int(item.ItemHash))
				bucketDef := destiny.InventoryBucket.Find(itemDef.Inventory.BucketTypeHash)

				switch bucketDef.DisplayProperties.Name {
				case "Kinetic Weapons":
					fallthrough
				case "Energy Weapons":
					fallthrough
				case "Power Weapons":
					log.Trace("Bag Weapon: ", itemDef.DisplayProperties.Name)
					bag.Weapons = append(bag.Weapons, NewWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny))
				case "Helmet":
					fallthrough
				case "Gauntlets":
					fallthrough
				case "Chest Armor":
					fallthrough
				case "Leg Armor":
					fallthrough
				case "Class Armor":
					log.Trace("Bag Armor: ", itemDef.DisplayProperties.Name)
					bag.Armors = append(bag.Armors, NewArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, destiny))
				default:
					log.Trace("Ignoring: ", bucketDef.DisplayProperties.Name)
				}
			}

			var class string
			if characterData.ClassHash > 0 {
				classDef := destiny.Class.Find(characterData.ClassHash)
				class = classDef.DisplayProperties.Name
			}

			var emblem bungo.Emblem
			if characterData.EmblemHash > 0 {
				itemDef := destiny.InventoryItem.Find(characterData.EmblemHash)

				emblem.Name = itemDef.DisplayProperties.Name
				emblem.InventoryItemId = itemDef.Hash
				emblem.IconUri = itemDef.DisplayProperties.Icon
				emblem.BannerUri = itemDef.SecondarySpecial
				emblem.IconOverlayUri = itemDef.SecondaryOverlay
				emblem.Background = bungo.Color{
					Red:   itemDef.BackgroundColor.Red,
					Green: itemDef.BackgroundColor.Green,
					Blue:  itemDef.BackgroundColor.Blue,
					Alpha: itemDef.BackgroundColor.Alpha,
				}
			}

			characters = append(characters, bungo.Character{
				Id:             characterData.CharacterID,
				MembershipType: characterData.MembershipType,

				Class:    class,
				Level:    characterData.BaseCharacterLevel,
				Light:    characterData.Light,
				Emblem:   emblem,
				Equipped: equipped,
				Bag:      bag,
			})
		}
	}

	return bungo.CurrentUser{
		Name:         bungieNetUser.DisplayName,
		MembershipId: bungieNetUser.MembershipID,
		Characters:   characters,
	}, nil
}

type CurrentCharacterArgs struct {
	MembershipType int
	Id             string

	OathClientId     string
	OathClientSecret string
	Auth             BungieToken
	BungieClient     *resty.Client
	Destiny          data.Content
}

func CurrentCharacter(args CurrentCharacterArgs) (out bungo.Character, err error) {
	currentUser, err := CurrentUser(CurrentUserArgs{
		OathClientId:     args.OathClientId,
		OathClientSecret: args.OathClientSecret,
		Auth:             args.Auth,
		BungieClient:     args.BungieClient,
		Destiny:          args.Destiny,
	})
	if err != nil {
		return bungo.Character{}, err
	}

	for _, character := range currentUser.Characters {
		if character.Id == args.Id && character.MembershipType == args.MembershipType {
			return character, nil
		}
	}
	return bungo.Character{}, errors.New("not found")
}
