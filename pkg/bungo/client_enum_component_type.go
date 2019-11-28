package bungo

import (
	"errors"
	"fmt"
)

const (
	ComponentsNoneKey                  = 0
	ComponentsProfilesKey              = 100
	ComponentsVendorReceiptsKey        = 101
	ComponentsProfileInventoriesKey    = 102
	ComponentsProfileCurrenciesKey     = 103
	ComponentsProfileProgressionKey    = 104
	ComponentsPlatformSilverKey        = 105
	ComponentsCharactersKey            = 200
	ComponentsCharacterInventoriesKey  = 201
	ComponentsCharacterProgressionsKey = 202
	CharacterRenderDataKey             = 203
	ComponentsCharacterActivitiesKey   = 204
	ComponentsCharacterEquipmentKey    = 205
	ComponentsItemInstancesKey         = 300
	ComponentsItemObjectivesKey        = 301
	ComponentsItemPerksKey             = 302
	ComponentsItemRenderDataKey        = 303
	ComponentsItemStatsKey             = 304
	ComponentsItemSocketsKey           = 305
	ComponentsItemTalentGridsKey       = 306
	ComponentsItemCommonDataKey        = 307
	ComponentsItemPlugStatesKey        = 308
	ComponentsVendorsKey               = 400
	ComponentsVendorCategoriesKey      = 401
	ComponentsVendorSalesKey           = 402
	ComponentsKiosksKey                = 500
	ComponentsCurrencyLookupsKey       = 600
	ComponentsPresentationNodesKey     = 700
	ComponentsCollectiblesKey          = 800
	ComponentsRecordsKey               = 900
	ComponentsTransitoryKey            = 1000
)

type ComponentType string

func ComponentTypesManyS(keys ...int) (out []string) {
	for _, e := range ComponentTypesMany(keys...) {
		out = append(out, string(e))
	}
	return out
}

func ComponentTypesMany(keys ...int) (out []ComponentType) {
	for _, key := range keys {
		eout, _, _ := ComponentTypesE(key)
		out = append(out, eout)
	}
	return out
}

func ComponentTypes(key int) ComponentType {
	out, _, _ := ComponentTypesE(key)
	return out
}

func ComponentTypesE(key int) (ComponentType, string, error) {
	switch key {
	case ComponentsNoneKey:
		return "None", "", nil
	case ComponentsProfilesKey:
		description := `
Profiles is the most basic component, only relevant when calling GetProfile.
This returns basic information about the profile, which is almost nothing:
- a list of characterIds,
- some information about the last time you logged in, and;
- that most sobering statistic: how long you've played.
`
		return "Profiles", description, nil
	case ComponentsVendorReceiptsKey:
		description := `
Only applicable for GetProfile, this will return information about receipts for refundable vendor items.
`
		return "VendorReceipts", description, nil
	case ComponentsProfileInventoriesKey:
		description := `
Asking for this will get you the profile-level inventories, such as your Vault buckets
(yeah, the Vault is really inventory buckets located on your Profile)
`
		return "ProfileInventories", description, nil
	case ComponentsProfileCurrenciesKey:
		description := `
This will get you a summary of items on your Profile that we consider to be "currencies", such as Glimmer. 
I mean, if there's Glimmer in Destiny 2. I didn't say there was Glimmer.
`
		return "ProfileCurrencies", description, nil
	case ComponentsProfileProgressionKey:
		description := `
This will get you any progression-related information that exists on a Profile-wide level, across all characters.
`
		return "ProfileProgression", description, nil
	case ComponentsPlatformSilverKey:
		description := `
This will get you information about the silver that this profile has on every platform on which it plays.

You may only request this component for the logged in user's Profile, and will not receive it if you request it for another Profile.
`
		return "PlatformSilver", description, nil
	case ComponentsCharactersKey:
		description := `
This will get you summary info about each of the characters in the profile.
`
		return "Characters", description, nil
	case ComponentsCharacterInventoriesKey:
		description := `
This will get you information about any non-equipped items on the character or character(s) in question,
if you're allowed to see it.

You have to either be authenticated as that user, or that user must allow anonymous viewing of their non-equipped items in Bungie.Net settings to actually get results.
`
		return "CharacterInventories", description, nil
	case ComponentsCharacterProgressionsKey:
		description := `
This will get you information about the progression (faction, experience, etc... "levels") relevant to each character.

You have to either be authenticated as that user, or that user must allow anonymous viewing of their progression info in Bungie.Net settings to actually get results.
`
		return "CharacterProgressions", description, nil
	case CharacterRenderDataKey:
		description := `
This will get you just enough information to be able to render the character in 3D if you have written a 3D rendering library for Destiny Characters, or "borrowed" ours.
It's okay, I won't tell anyone if you're using it.
I'm no snitch. (actually, we don't care if you use it - go to town)
`
		return "RenderData", description, nil
	case ComponentsCharacterActivitiesKey:
		description := `
This will return info about activities that a user can see and gating on it, if you are the currently authenticated user or the user has elected to allow anonymous viewing of its progression info. 
Note that the data returned by this can be unfortunately problematic and relatively unreliable in some cases.
We'll eventually work on making it more consistently reliable.
`
		return "CharacterActivities", description, nil
	case ComponentsCharacterEquipmentKey:
		description := `
This will return info about the equipped items on the character(s). Everyone can see this.
`
		return "CharacterEquipment", description, nil
	case ComponentsItemInstancesKey:
		description := `
This will return basic info about instanced items - whether they can be equipped, their tracked status, and some info commonly needed in many places (current damage type, primary stat value, etc)
`
		return "ItemInstances", description, nil
	case ComponentsItemObjectivesKey:
		description := `
Items can have Objectives (DestinyObjectiveDefinition) bound to them.
If they do, this will return info for items that have such bound objectives.
`
		return "ItemObjectives", description, nil
	case ComponentsItemPerksKey:
		description := `
Items can have perks (DestinyPerkDefinition).
If they do, this will return info for what perks are active on items.
`
		return "ItemPerks", description, nil
	case ComponentsItemRenderDataKey:
		description := `
If you just want to render the weapon, this is just enough info to do that rendering.
`
		return "ItemRenderData", description, nil
	case ComponentsItemStatsKey:
		description := `
Items can have stats, like rate of fire. Asking for this component will return requested item's stats if they have stats.
`
		return "ItemStats", description, nil
	case ComponentsItemSocketsKey:
		description := `
Items can have sockets, where plugs can be inserted.
Asking for this component will return all info relevant to the sockets on items that have them.
`
		return "ItemSockets", description, nil
	case ComponentsItemTalentGridsKey:
		description := `
Items can have talent grids, though that matters a lot less frequently than it used to. 
Asking for this component will return all relevant info about activated Nodes and Steps on this talent grid, like the good ol' days.
`
		return "ItemTalentGrids", description, nil
	case ComponentsItemCommonDataKey:
		description := `
Items that *aren't* instanced still have important information you need to know:
- how much of it you have,;
- the itemHash so you can look up their DestinyInventoryItemDefinition,;
- whether they're locked,;
- etc...
Both instanced and non-instanced items will have these properties.
You will get this automatically with Inventory components - you only need to pass this when calling GetItem on a specific item.
`
		return "ItemCommonData", description, nil
	case ComponentsItemPlugStatesKey:
		description := `
Items that are "Plugs" can be inserted into sockets.
This returns statuses about those plugs and why they can/can't be inserted.

I hear you giggling, there's nothing funny about inserting plugs. Get your head out of the gutter and pay attention!
`
		return "ItemPlugStates", description, nil
	case ComponentsVendorsKey:
		description := `
When obtaining vendor information, this will return summary information about the Vendor or Vendors being returned.
`
		return "Vendors", description, nil
	case ComponentsVendorCategoriesKey:
		description := `
When obtaining vendor information, this will return information about the categories of items provided by the Vendor.
`
		return "VendorCategories", description, nil
	case ComponentsVendorSalesKey:
		description := `
When obtaining vendor information, this will return the information about items being sold by the Vendor.
`
		return "VendorSales", description, nil
	case ComponentsKiosksKey:
		description := `
Asking for this component will return you the account's Kiosk statuses: that is, what items have been filled out/acquired.

But only if you are the currently authenticated user or the user has elected to allow anonymous viewing of its progression info.
`
		return "Kiosks", description, nil
	case ComponentsCurrencyLookupsKey:
		description := `
A "shortcut" component that will give you all of the item hashes/quantities of items that the requested character can use to determine if an action (purchasing, socket insertion) has the required currency.

(recall that all currencies are just items, and that some vendor purchases require items that you might not traditionally consider to be a "currency", like plugs/mods!)
`
		return "CurrencyLookups", description, nil
	case ComponentsPresentationNodesKey:
		description := `
Returns summary status information about all "Presentation Nodes".
See DestinyPresentationNodeDefinition for more details, but the gist is that these are entities used by the game UI to bucket Collectibles and Records into a hierarchy of categories.

You may ask for and use this data if you want to perform similar bucketing in your own UI: or you can skip it and roll your own.
`
		return "PresentationNodes", description, nil
	case ComponentsCollectiblesKey:
		description := `
Returns summary status information about all "Collectibles".
These are records of what items you've discovered while playing Destiny, and some other basic information.

For detailed information, you will have to call a separate endpoint devoted to the purpose.
`
		return "Collectibles", description, nil
	case ComponentsRecordsKey:
		description := `
Returns summary status information about all "Records" also known in the game as "Triumphs".

I know, it's confusing because there's also "Moments of Triumph" that will themselves be represented as "Triumphs."
`
		return "Records", description, nil
	case ComponentsTransitoryKey:
		description := `
Returns information that Bungie considers to be "Transitory":
- data that may change too frequently; or
- come from a non-authoritative source 
such that we don't consider the data to be fully trustworthy, but that might prove useful for some limited use cases.

We can provide no guarantee of timeliness nor consistency for this data: buyer beware with the Transitory component.
`
		return "Transitory", description, nil
	default:
		return "", "", errors.New(fmt.Sprintf("unknown key: %d", key))
	}
}
