package brightsvc

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"

	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Barracks struct {
	Tower        tower.Tower
	Ghost        ghost.Ghost
	BungieClient *resty.Client
}

func (b Barracks) ReportForDuty() (out bungo.User) {
	b.Tower.Redis.Connect()
	defer b.Tower.Redis.Disconnect()

	freshToken := Refresh(b.BungieClient, b.Ghost, b.Tower)
	getMembershipDataForCurrentUserResponse, err := b.BungieClient.R().
		EnableTrace().
		SetAuthToken(freshToken.AccessToken).
		Get("/Platform/User/GetMembershipsForCurrentUser/")
	if err != nil {
		log.Fatal(err)
	}

	var getMembershipDataForCurrentUser bungo.GetMembershipDataForCurrentUser
	if err := json.Unmarshal(getMembershipDataForCurrentUserResponse.Body(), &getMembershipDataForCurrentUser); err != nil {
		log.Warn(string(getMembershipDataForCurrentUserResponse.Body()))
		log.Fatal(err)
	}

	bungieNetUser := getMembershipDataForCurrentUser.Response.BungieNetUser
	destinyMemberships := getMembershipDataForCurrentUser.Response.DestinyMemberships

	log.Debug("/Platform/User/GetMembershipsForCurrentUser -> ", bungieNetUser.DisplayName)
	var characters []bungo.Character
	for _, destinyMembership := range destinyMemberships {

		getProfileResp, err := b.BungieClient.R().
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
			Get("/Platform/Destiny2/{membershipType}/Profile/{destinyMembershipId}/")
		if err != nil {
			log.Fatal(err)
		}

		var getProfile bungo.GetProfile
		if err := json.Unmarshal(getProfileResp.Body(), &getProfile); err != nil {
			log.Warn(string(getProfileResp.Body()))
			log.Fatal(err)
		}

		profileData := getProfile.Response.Profile.Data

		for _, characterId := range profileData.CharacterIds {

			getCharacterResp, err := b.BungieClient.R().
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
				Get("/Platform/Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/")
			if err != nil {
				log.Fatal(err)
			}

			log.Warn(string(getCharacterResp.Body()))

			var getCharacter bungo.GetCharacter
			if err := json.Unmarshal(getCharacterResp.Body(), &getCharacter); err != nil {
				log.Warn(string(getCharacterResp.Body()))
				log.Fatal(err)
			}

			characterData := getCharacter.Response.Character.Data
			equipmentItemsData := getCharacter.Response.Equipment.Data
			inventoryItemsData := getCharacter.Response.Inventory.Data

			var equipped bungo.Outfit
			for _, item := range equipmentItemsData.Items {

				itemDef := datax.GetInventoryItem(int(item.ItemHash), b.Tower.Redis)
				bucketDef := datax.GetInventoryBucket(itemDef.Inventory.BucketTypeHash, b.Tower.Redis)

				switch bucketDef.DisplayProperties.Name {
				case "Kinetic Weapons":
					log.Debug("Kinetic: ", itemDef.DisplayProperties.Name)
					equipped.KineticWeapon = newWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Energy Weapons":
					log.Debug("Energy: ", itemDef.DisplayProperties.Name)
					equipped.EnergyWeapon = newWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Power Weapons":
					log.Debug("Power: ", itemDef.DisplayProperties.Name)
					equipped.PowerWeapon = newWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Helmet":
					log.Debug("Helmet: ", itemDef.DisplayProperties.Name)
					equipped.Helmet = newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Gauntlets":
					log.Debug("Gauntlets: ", itemDef.DisplayProperties.Name)
					equipped.Gauntlets = newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Chest Armor":
					log.Debug("Chest: ", itemDef.DisplayProperties.Name)
					equipped.Chest = newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Leg Armor":
					log.Debug("Legs: ", itemDef.DisplayProperties.Name)
					equipped.Leg = newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				case "Class Armor":
					log.Debug("Class: ", itemDef.DisplayProperties.Name)
					equipped.Class = newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis)
				default:
					log.Trace("Ignoring: ", bucketDef.DisplayProperties.Name)
				}
			}

			var bag bungo.Bag
			for _, item := range inventoryItemsData.Items {

				itemDef := datax.GetInventoryItem(int(item.ItemHash), b.Tower.Redis)
				bucketDef := datax.GetInventoryBucket(itemDef.Inventory.BucketTypeHash, b.Tower.Redis)

				switch bucketDef.DisplayProperties.Name {
				case "Kinetic Weapons":
					fallthrough
				case "Energy Weapons":
					fallthrough
				case "Power Weapons":
					log.Trace("Bag Weapon: ", itemDef.DisplayProperties.Name)
					bag.Weapons = append(bag.Weapons, newWeaponInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis))
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
					bag.Armors = append(bag.Armors, newArmorInstance(characterId, item.ItemInstanceID, itemDef.Hash, b.Tower.Redis))
				default:
					log.Trace("Ignoring: ", bucketDef.DisplayProperties.Name)
				}
			}

			var class string
			if characterData.ClassHash > 0 {
				classDef := datax.GetClass(characterData.ClassHash, b.Tower.Redis)
				class = classDef.DisplayProperties.Name
			}

			var emblem bungo.Emblem
			if characterData.EmblemHash > 0 {
				itemDef := datax.GetInventoryItem(characterData.EmblemHash, b.Tower.Redis)

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

	return bungo.User{
		Name:         bungieNetUser.DisplayName,
		MembershipId: bungieNetUser.MembershipID,
		Characters:   characters,
	}

}
