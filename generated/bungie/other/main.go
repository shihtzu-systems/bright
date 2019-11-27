package other

var AllComponents = []int32{
	ComponentsProfiles, ComponentsVendorReceipts, ComponentsProfileInventories, ComponentsProfileCurrencies, ComponentsProfileProgression, ComponentsPlatformSilver, ComponentsCharacters, ComponentsCharacterInventories, ComponentsCharacterProgressions, CharacterRenderData, ComponentsCharacterActivities, ComponentsCharacterEquipment, ComponentsItemInstances, ComponentsItemObjectives, ComponentsItemPerks, ComponentsItemRenderData, ComponentsItemStats, ComponentsItemSockets, ComponentsItemTalentGrids, ComponentsItemCommonData, ComponentsItemPlugStates, ComponentsVendors, ComponentsVendorCategories, ComponentsVendorSales, ComponentsKiosks, ComponentsCurrencyLookups, ComponentsPresentationNodes, ComponentsCollectibles, ComponentsRecords, ComponentsTransitory,
}

var BasicComponents = []int32{
	ComponentsProfiles, ComponentsCharacters, ComponentsCharacterActivities,
}

var CharacterComponents = []int32{
	ComponentsCharacters, ComponentsCharacterInventories, ComponentsCharacterProgressions,
}

const (
	MembershipTypeNone          = 0
	MembershipTypeTigerXbox     = 1
	MembershipTypeTigerPsn      = 2
	MembershipTypeTigerSteam    = 3
	MembershipTypeTigerBlizzard = 4
	MembershipTypeTigerStadia   = 5
	MembershipTigerDemon        = 10
	MembershipBungieNext        = 254
	// "All" is only valid for searching capabilities:
	// you need to pass the actual matching BungieMembershipType for any query where you pass a known membershipId.
	MembershipTypeAll = -1

	ComponentsNone = 0

	// Profiles is the most basic component, only relevant when calling GetProfile.
	// This returns basic information about the profile, which is almost nothing:
	// - a list of characterIds,
	// - some information about the last time you logged in, and;
	// - that most sobering statistic: how long you've played.
	ComponentsProfiles = 100

	// Only applicable for GetProfile, this will return information about receipts
	// for refundable vendor items.
	ComponentsVendorReceipts = 101

	// Asking for this will get you the profile-level inventories, such as your Vault buckets
	// (yeah, the Vault is really inventory buckets located on your Profile)
	ComponentsProfileInventories = 102

	// This will get you a summary of items on your Profile that we consider to be "currencies",
	// such as Glimmer. I mean, if there's Glimmer in Destiny 2. I didn't say there was Glimmer.
	ComponentsProfileCurrencies = 103

	// This will get you any progression-related information that exists
	// on a Profile-wide level, across all characters.
	ComponentsProfileProgression = 104

	// This will get you information about the silver that this profile has on
	// every platform on which it plays.
	//
	// You may only request this component for the logged in user's Profile,
	// and will not recieve it if you request it for another Profile.
	ComponentsPlatformSilver = 105

	// This will get you summary info about each of the characters in the profile.
	ComponentsCharacters = 200

	// This will get you information about any non-equipped items on the character or character(s) in question,
	// if you're allowed to see it.
	// You have to either be authenticated as that user, or that user must allow anonymous
	// viewing of their non-equipped items in Bungie.Net settings to actually get results.
	ComponentsCharacterInventories = 201

	//This will get you information about the progression (faction, experience, etc... "levels")
	// relevant to each character, if you are the currently authenticated user or
	// the user has elected to allow anonymous viewing of its progression info.
	ComponentsCharacterProgressions = 202

	// This will get you just enough information to be able to render the character in 3D
	// if you have written a 3D rendering library for Destiny Characters, or "borrowed" ours.
	// It's okay, I won't tell anyone if you're using it.
	// I'm no snitch. (actually, we don't care if you use it - go to town)
	CharacterRenderData = 203

	// This will return info about activities that a user can see and gating on it,
	// if you are the currently authenticated user or the user has elected to
	// allow anonymous viewing of its progression info. Note that the data returned by
	// this can be unfortunately problematic and relatively unreliable in some cases.
	// We'll eventually work on making it more consistently reliable.
	ComponentsCharacterActivities = 204

	// This will return info about the equipped items on the character(s). Everyone can see this.
	ComponentsCharacterEquipment = 205

	// This will return basic info about instanced items - whether they can be equipped,
	// their tracked status, and some info commonly needed in many places
	// (current damage type, primary stat value, etc)
	ComponentsItemInstances = 300

	// Items can have Objectives (DestinyObjectiveDefinition) bound to them.
	// If they do, this will return info for items that have such bound objectives.
	ComponentsItemObjectives = 301

	// Items can have perks (DestinyPerkDefinition).
	// If they do, this will return info for what perks are active on items.
	ComponentsItemPerks = 302

	// If you just want to render the weapon, this is just enough info to do that rendering.
	ComponentsItemRenderData = 303

	// Items can have stats, like rate of fire. Asking for this component
	// will return requested item's stats if they have stats.
	ComponentsItemStats = 304

	// Items can have sockets, where plugs can be inserted.
	// Asking for this component will return all info relevant
	// to the sockets on items that have them.
	ComponentsItemSockets = 305

	// Items can have talent grids, though that matters a lot less frequently
	// than it used to. Asking for this component will return all relevant info
	// about activated Nodes and Steps on this talent grid, like the good ol' days.
	ComponentsItemTalentGrids = 306

	// Items that *aren't* instanced still have important information you need to know:
	// - how much of it you have,;
	// - the itemHash so you can look up their DestinyInventoryItemDefinition,;
	// - whether they're locked,;
	// - etc...
	//
	// Both instanced and non-instanced items will have these properties.
	// You will get this automatically with Inventory components - you only need to pass
	// this when calling GetItem on a specific item.
	ComponentsItemCommonData = 307

	// Items that are "Plugs" can be inserted into sockets.
	// This returns statuses about those plugs and
	// why they can/can't be inserted.
	//
	// I hear you giggling, there's nothing funny about inserting plugs.
	// Get your head out of the gutter and pay attention!
	ComponentsItemPlugStates = 308

	// When obtaining vendor information, this will return summary information
	// about the Vendor or Vendors being returned.
	ComponentsVendors = 400

	// When obtaining vendor information, this will return information
	// about the categories of items provided by the Vendor.
	ComponentsVendorCategories = 401

	// When obtaining vendor information, this will return the
	// information about items being sold by the Vendor.
	ComponentsVendorSales = 402

	// Asking for this component will return you the account's Kiosk statuses:
	// that is, what items have been filled out/acquired.
	// But only if you are the currently authenticated user
	// or the user has elected to allow anonymous viewing of its progression info.
	ComponentsKiosks = 500

	// A "shortcut" component that will give you all of the item hashes/quantities of
	// items that the requested character can use to determine
	// if an action (purchasing, socket insertion) has the required currency.
	//
	// (recall that all currencies are just items, and that some vendor purchases
	// require items that you might not traditionally consider to be a "currency", like plugs/mods!)
	ComponentsCurrencyLookups = 600

	// Returns summary status information about all "Presentation Nodes".
	// See DestinyPresentationNodeDefinition for more details,
	// but the gist is that these are entities used by the game UI to bucket
	// Collectibles and Records into a hierarchy of categories.
	//
	// You may ask for and use this data if you want to perform similar bucketing in your own UI:
	// or you can skip it and roll your own.
	ComponentsPresentationNodes = 700

	// Returns summary status information about all "Collectibles".
	// These are records of what items you've discovered while playing Destiny,
	// and some other basic information.
	//
	// For detailed information, you will have to call a separate endpoint devoted to the purpose.
	ComponentsCollectibles = 800

	// Returns summary status information about all "Records" also known in the game as "Triumphs".
	// I know, it's confusing because there's also "Moments of Triumph"
	// that will themselves be represented as "Triumphs."
	ComponentsRecords = 900

	// Returns information that Bungie considers to be "Transitory":
	// data that may change too frequently or
	// come from a non-authoritative source such that
	// we don't consider the data to be fully trustworthy,
	// but that might prove useful for some limited use cases.
	//
	// We can provide no guarantee of timeliness nor consistency for this data:
	// buyer beware with the Transitory component.
	ComponentsTransitory = 1000

	// collectible state

	// None
	CollectibleStateNone = 0

	// Not Acquired
	// If this flag is set, you have not yet obtained this collectible.
	CollectibleStateNotAcquired = 1

	// Obscured
	// If this flag is set, the item is "obscured" to you:
	// you can/should use the alternate item hash found in
	// DestinyCollectibleDefinition.stateInfo.obscuredOverrideItemHash when displaying
	// this collectible instead of the default display info.
	CollectibleStateObscured = 2

	// Invisible
	// If this flag is set, the collectible should not be shown to the user.
	//
	// Please do consider honoring this flag. It is used - for example - to hide items that
	// a person didn't get from the Eververse. I can't prevent these from being returned in definitions,
	// because some people may have acquired them and thus they should show up:
	// but I would hate for people to start feeling some variant of a Collector's Remorse about these items,
	// and thus increasing their purchasing based on that compulsion.
	//
	// That would be a very unfortunate outcome, and one that I wouldn't like to see happen.
	// So please, whether or not I'm your mom, consider honoring this flag
	// and don't show people invisible collectibles.
	CollectibleStateInvisible = 4

	// CannotAffordMaterialRequirements
	// If this flag is set, the collectible requires payment for creating an instance of the item,
	// and you are lacking in currency. Bring the benjamins next time. Or spinmetal. Whatever.
	CollectibleStateCannotAffordMaterialRequirements = 8

	// InventorySpaceUnavailable
	// If this flag is set, you can't pull this item out of your collection because there's no room left in your inventory.
	CollectibleStateInventorySpaceUnavailable = 16

	// UniquenessViolation
	// If this flag is set, you already have one of these items and can't have a second one.
	CollectibleStateUniquenessViolation = 32

	// PurchaseDisabled: 64
	// If this flag is set, the ability to pull this item out of your collection has been disabled.
	CollectibleStatePurchaseDisabled = 64
)

var CollectibleState = map[int]struct {
	Name        string
	Description string
}{
	CollectibleStateNone: {
		"none", "none",
	},
	CollectibleStateNotAcquired: {
		"not acquired", "you have not yet obtained this collectible",
	},
	CollectibleStateObscured: {
		"obscured", "the item is \"obscured\" to you",
	},
	CollectibleStateInvisible: {
		"invisible", "the collectible should not be shown to the user",
	},
	CollectibleStateCannotAffordMaterialRequirements: {
		"unaffordable", "the collectible requires payment for creating an instance of the item and you are lacking in currency",
	},
	CollectibleStateInventorySpaceUnavailable: {
		"no room", "you can't pull this item out of your collection because there's no room left in your inventory",
	},
	CollectibleStateUniquenessViolation: {
		"duplicate", "you already have one of these items and can't have a second one",
	},
	CollectibleStatePurchaseDisabled: {
		"disabled", "the ability to pull this item out of your collection has been disabled",
	},
}
