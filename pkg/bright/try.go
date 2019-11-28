package bright

import (
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	log "github.com/sirupsen/logrus"
)

type TryUserArgs struct {
	Destiny data.Content
	Dad     bool
}

type TryUserOut struct {
	CurrentUser bungo.Gamer
}

func TryUser(args TryUserArgs) (out bungo.Gamer, err error) {
	return bungo.Gamer{
		Name:         GenerateName(args.Dad),
		MembershipId: "0",
		Guardians: []bungo.Guardian{
			newTitan(args.Destiny),
			newHunter(args.Destiny),
			newWarlock(args.Destiny),
		},
	}, nil
}

func TryUserX(dad bool) (out bungo.Gamer) {
	return bungo.Gamer{
		Name:         GenerateName(dad),
		MembershipId: "0",
		Guardians: []bungo.Guardian{
			newTitanX(),
			newHunterX(),
			newWarlockX(),
		},
	}
}

func newTitanX() bungo.Guardian {
	return bungo.Guardian{
		Id:             tryTitanId,
		MembershipType: 0,
		Class:          "Titan",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}

func newHunterX() bungo.Guardian {
	return bungo.Guardian{
		Id:             tryHunterId,
		MembershipType: 0,
		Class:          "Hunter",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}

func newWarlockX() bungo.Guardian {
	return bungo.Guardian{
		Id:             tryWarlockId,
		MembershipType: 0,
		Class:          "Warlock",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}

const (
	// emblems
	thePyramidionEmblem      = 10493725
	chasingPerfectionEmblem  = 19962737
	forValorEmblem           = 29194593
	ignitionUnsparkedEmblem  = 313685909
	wishGranterEmblem        = 338612265
	zeroFutureEmblem         = 414672658
	ludomaniacalEmblem       = 540603118
	mentorOfLightEmblem      = 481345527
	futureWarCultistEmblem   = 745759695
	operationPiccoloEmblem   = 769740915
	aGardenWorldEmblem       = 893502917
	victoriousVeteranEmblem  = 969863968
	spireStarEmblem          = 1057119308
	vanguardTerminus         = 1075647353
	damageIncorporatedEmblem = 1131244817
	theGoldenStandardEmblem  = 1162929425
	rasmussensGiftEmblem     = 1347384483
	arcologyInterfaceEmblem  = 1409726989
	jadeRabbitReduxEmblem    = 1409726987
	tryTheSpecialEmblem      = 1409726988
	kadi5530sBlessingEmblem  = 1940590818
	sonicSimulationEmblem    = 1940590825
	darkOrbitEmblem          = 2526736328
	exodusAccessEmblem       = 2484637938
	splishSplashEmblem       = 2984066626
	sunsetEchoesEmblem       = 3605490915
	firstToTheForgeEmblem    = 3961503937
	systemOfPeaceEmblem      = 4011836824
	trueColorsEmblem         = 3961503936
	prismaticInfernoEmblem   = 4159550313
	emperorsEnvyEmblem       = 4261480751
)

var emblems = []int{
	thePyramidionEmblem,
	chasingPerfectionEmblem,
	forValorEmblem,
	ignitionUnsparkedEmblem,
	wishGranterEmblem,
	zeroFutureEmblem,
	ludomaniacalEmblem,
	mentorOfLightEmblem,
	futureWarCultistEmblem,
	operationPiccoloEmblem,
	aGardenWorldEmblem,
	victoriousVeteranEmblem,
	spireStarEmblem,
	vanguardTerminus,
	damageIncorporatedEmblem,
	theGoldenStandardEmblem,
	rasmussensGiftEmblem,
	arcologyInterfaceEmblem,
	jadeRabbitReduxEmblem,
	tryTheSpecialEmblem,
	kadi5530sBlessingEmblem,
	sonicSimulationEmblem,
	darkOrbitEmblem,
	exodusAccessEmblem,
	splishSplashEmblem,
	sunsetEchoesEmblem,
	firstToTheForgeEmblem,
	systemOfPeaceEmblem,
	trueColorsEmblem,
	prismaticInfernoEmblem,
	emperorsEnvyEmblem,
}

const tryTitanId = "titan"

func newTitan(destiny data.Content) bungo.Guardian {
	emblemDef := destiny.InventoryItem.Find(emblems[randomGen.Intn(len(emblems))])
	return bungo.Guardian{
		Id:             tryTitanId,
		MembershipType: 0,
		Class:          "Titan",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
		Emblem: bungo.Emblem{
			Name:            emblemDef.DisplayProperties.Name,
			InventoryItemId: emblemDef.Hash,
			IconUri:         emblemDef.DisplayProperties.Icon,
			BannerUri:       emblemDef.SecondarySpecial,
			IconOverlayUri:  emblemDef.SecondaryOverlay,
			Background: bungo.Color{
				Red:   emblemDef.BackgroundColor.Red,
				Green: emblemDef.BackgroundColor.Green,
				Blue:  emblemDef.BackgroundColor.Blue,
				Alpha: emblemDef.BackgroundColor.Alpha,
			},
		},
		Equipped: newOutfit(tryTitanId, destiny),
		Bag:      newBag(tryTitanId, destiny),
	}
}

const tryHunterId = "hunter"

func newHunter(destiny data.Content) bungo.Guardian {
	emblemDef := destiny.InventoryItem.Find(emblems[randomGen.Intn(len(emblems))])
	outfit := newOutfit(tryHunterId, destiny)
	bag := newBag(tryHunterId, destiny)
	return bungo.Guardian{
		Id:             tryHunterId,
		MembershipType: 0,
		Class:          "Hunter",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
		Emblem: bungo.Emblem{
			Name:            emblemDef.DisplayProperties.Name,
			InventoryItemId: emblemDef.Hash,
			IconUri:         emblemDef.DisplayProperties.Icon,
			BannerUri:       emblemDef.SecondarySpecial,
			IconOverlayUri:  emblemDef.SecondaryOverlay,
			Background: bungo.Color{
				Red:   emblemDef.BackgroundColor.Red,
				Green: emblemDef.BackgroundColor.Green,
				Blue:  emblemDef.BackgroundColor.Blue,
				Alpha: emblemDef.BackgroundColor.Alpha,
			},
		},
		Equipped: outfit,
		Bag:      bag,
	}
}

const tryWarlockId = "warlock"

func newWarlock(destiny data.Content) bungo.Guardian {
	emblemDef := destiny.InventoryItem.Find(emblems[randomGen.Intn(len(emblems))])
	outfit := newOutfit(tryWarlockId, destiny)
	bag := newBag(tryWarlockId, destiny)
	return bungo.Guardian{
		Id:             tryWarlockId,
		MembershipType: 0,
		Class:          "Warlock",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
		Emblem: bungo.Emblem{
			Name:            emblemDef.DisplayProperties.Name,
			InventoryItemId: emblemDef.Hash,
			IconUri:         emblemDef.DisplayProperties.Icon,
			BannerUri:       emblemDef.SecondarySpecial,
			IconOverlayUri:  emblemDef.SecondaryOverlay,
			Background: bungo.Color{
				Red:   emblemDef.BackgroundColor.Red,
				Green: emblemDef.BackgroundColor.Green,
				Blue:  emblemDef.BackgroundColor.Blue,
				Alpha: emblemDef.BackgroundColor.Alpha,
			},
		},
		Equipped: outfit,
		Bag:      bag,
	}
}

func newOutfit(characterId string, destiny data.Content) bungo.Outfit {
	switch characterId {
	case "warlock":
		return bungo.Outfit{
			KineticWeapon: NewWeapon(characterId, kinetics[randomGen.Intn(len(kinetics))], destiny),
			EnergyWeapon:  NewWeapon(characterId, energies[randomGen.Intn(len(energies))], destiny),
			PowerWeapon:   NewWeapon(characterId, powers[randomGen.Intn(len(powers))], destiny),

			Helmet:    NewArmor(characterId, warlockHelmets[randomGen.Intn(len(warlockHelmets))], destiny),
			Gauntlets: NewArmor(characterId, warlockGauntlets[randomGen.Intn(len(warlockGauntlets))], destiny),
			Chest:     NewArmor(characterId, warlockChests[randomGen.Intn(len(warlockChests))], destiny),
			Leg:       NewArmor(characterId, warlockLegs[randomGen.Intn(len(warlockLegs))], destiny),
			Class:     NewArmor(characterId, warlockClasses[randomGen.Intn(len(warlockClasses))], destiny),
		}
	case "titan":
		return bungo.Outfit{
			KineticWeapon: NewWeapon(characterId, kinetics[randomGen.Intn(len(kinetics))], destiny),
			EnergyWeapon:  NewWeapon(characterId, energies[randomGen.Intn(len(energies))], destiny),
			PowerWeapon:   NewWeapon(characterId, powers[randomGen.Intn(len(powers))], destiny),

			Helmet:    NewArmor(characterId, titanHelmets[randomGen.Intn(len(titanHelmets))], destiny),
			Gauntlets: NewArmor(characterId, titanGauntlets[randomGen.Intn(len(titanGauntlets))], destiny),
			Chest:     NewArmor(characterId, titanChests[randomGen.Intn(len(titanChests))], destiny),
			Leg:       NewArmor(characterId, titanLegs[randomGen.Intn(len(titanLegs))], destiny),
			Class:     NewArmor(characterId, titanClasses[randomGen.Intn(len(titanClasses))], destiny),
		}
	case "hunter":
		fallthrough
	default:
		return bungo.Outfit{
			KineticWeapon: NewWeapon(characterId, kinetics[randomGen.Intn(len(kinetics))], destiny),
			EnergyWeapon:  NewWeapon(characterId, energies[randomGen.Intn(len(energies))], destiny),
			PowerWeapon:   NewWeapon(characterId, powers[randomGen.Intn(len(powers))], destiny),

			Helmet:    NewArmor(characterId, hunterHelmets[randomGen.Intn(len(hunterHelmets))], destiny),
			Gauntlets: NewArmor(characterId, hunterGauntlets[randomGen.Intn(len(hunterGauntlets))], destiny),
			Chest:     NewArmor(characterId, hunterChests[randomGen.Intn(len(hunterChests))], destiny),
			Leg:       NewArmor(characterId, hunterLegs[randomGen.Intn(len(hunterLegs))], destiny),
			Class:     NewArmor(characterId, hunterClasses[randomGen.Intn(len(hunterClasses))], destiny),
		}
	}
}

const (
	// kinetic
	dukeMk44HandCannon         = 2112909414
	brayTechWerewolfAutoRifle  = 528834068
	sacredProvenancePulseRifle = 2408405461
	guidingStarAutoRifle       = 472847207
	peaceByConsensusSidearm    = 806021398
	frontierJusticeScoutRifle  = 273396910
	theSpitefulFangBow         = 421573768

	// energy
	borealisSniperRifle      = 3141979347
	drangBaroqueSideArm      = 79075821
	nationOfBeastsHandcannon = 654370424
	impromptu49PulseRifle    = 228424224

	// power
	nightTerrorSword          = 3870811754
	curtainCallRocketLauncher = 137879537
	iAmAliveGrenadeLauncher   = 3886263130
	zephyrSword               = 4203034886
)

var kinetics = []int{
	dukeMk44HandCannon,
	brayTechWerewolfAutoRifle,
	sacredProvenancePulseRifle,
	guidingStarAutoRifle,
	peaceByConsensusSidearm,
	frontierJusticeScoutRifle,
	theSpitefulFangBow,
}

var energies = []int{
	borealisSniperRifle,
	drangBaroqueSideArm,
	nationOfBeastsHandcannon,
	impromptu49PulseRifle,
}

var powers = []int{
	nightTerrorSword,
	curtainCallRocketLauncher,
	iAmAliveGrenadeLauncher,
	zephyrSword,
}

const (
	// hunter
	masqueradersCowlHelmet  = 2352138838
	dreambaneGripsGauntlets = 3571441640
	dreambaneVestChest      = 883769696
	dreambaneStridesLegs    = 377813570
	dreambaneCloakClass     = 193805725
)

var hunterHelmets = []int{
	masqueradersCowlHelmet,
}

var hunterGauntlets = []int{
	dreambaneGripsGauntlets,
}

var hunterChests = []int{
	dreambaneVestChest,
}

var hunterLegs = []int{
	dreambaneStridesLegs,
}

var hunterClasses = []int{
	dreambaneCloakClass,
}

const (
	// titan
	shadowsHelmHelmet         = 1129634130
	shadowsGauntletsGauntlets = 1595987387
	shadowsPlateChect         = 1862963733
	shadowsGreavesLegs        = 4450861
	shadowsMarkClass          = 311429764
)

var titanHelmets = []int{
	shadowsHelmHelmet,
}

var titanGauntlets = []int{
	shadowsGauntletsGauntlets,
}

var titanChests = []int{
	shadowsPlateChect,
}

var titanLegs = []int{
	shadowsGreavesLegs,
}

var titanClasses = []int{
	shadowsMarkClass,
}

const (
	// warlock
	solsticeHoodMajesticHelmet      = 692234472
	solsticeGlovesMajesticGauntlets = 3571441640
	solsticeRobesMajesticChest      = 450844637
	solsticeBootsMajesticLegs       = 1724587366
	solsticeBondMajesticClass       = 855792702
)

var warlockHelmets = []int{
	solsticeHoodMajesticHelmet,
}

var warlockGauntlets = []int{
	solsticeGlovesMajesticGauntlets,
}

var warlockChests = []int{
	solsticeRobesMajesticChest,
}

var warlockLegs = []int{
	solsticeBootsMajesticLegs,
}

var warlockClasses = []int{
	solsticeBondMajesticClass,
}

const (
	kineticCategory = 2
	energyCategory  = 3
	powerCategory   = 4

	autoRifleCategory       = 5
	handCannonCategory      = 6
	pulseRifleCategory      = 7
	scoutRifleCategory      = 8
	fusionRifleCategory     = 9
	sniperRifleCategory     = 10
	shotgunCategory         = 11
	machineGunCategory      = 12
	rocketLauncherCategory  = 13
	sidearmCategory         = 14
	swordCategory           = 54
	grenadeLauncherCategory = 153950757
	tracerRifleCategory     = 2489664120
	bowCategory             = 3317538576
)

func NewWeaponInstance(characterId, instanceId string, inventoryItemId int, destiny data.Content) (out bungo.Weapon) {
	out = NewWeapon(characterId, inventoryItemId, destiny)
	out.InstanceId = instanceId
	return out
}

func NewWeapon(characterId string, inventoryItemId int, destiny data.Content) bungo.Weapon {
	var weaponBucket string
	var weaponSlot string
	var weaponType string
	itemDef := destiny.InventoryItem.Find(inventoryItemId)
	for _, itemCategoryHash := range itemDef.ItemCategoryHashes {
		itemCategoryDef := destiny.ItemCategory.Find(itemCategoryHash)
		if itemCategoryHash == kineticCategory ||
			itemCategoryHash == energyCategory ||
			itemCategoryHash == powerCategory {
			weaponBucket = itemCategoryDef.ShortTitle
			weaponSlot = itemCategoryDef.ShortTitle
		}
		if itemCategoryHash == autoRifleCategory ||
			itemCategoryHash == handCannonCategory ||
			itemCategoryHash == pulseRifleCategory ||
			itemCategoryHash == scoutRifleCategory ||
			itemCategoryHash == fusionRifleCategory ||
			itemCategoryHash == sniperRifleCategory ||
			itemCategoryHash == shotgunCategory ||
			itemCategoryHash == machineGunCategory ||
			itemCategoryHash == rocketLauncherCategory ||
			itemCategoryHash == sidearmCategory ||
			itemCategoryHash == swordCategory ||
			itemCategoryHash == grenadeLauncherCategory ||
			itemCategoryHash == tracerRifleCategory ||
			itemCategoryHash == bowCategory {
			weaponType = itemCategoryDef.ShortTitle
		}
	}
	return bungo.Weapon{
		MembershipType:  0,
		CharacterId:     characterId,
		InstanceId:      fmt.Sprintf("instance-of-%d", inventoryItemId),
		InventoryItemId: inventoryItemId,

		Tier:        itemDef.Inventory.TierTypeName,
		Bucket:      weaponBucket,
		Name:        itemDef.DisplayProperties.Name,
		Type:        weaponType,
		IconEnabled: itemDef.DisplayProperties.HasIcon,
		IconUri:     itemDef.DisplayProperties.Icon,
		Slot: bungo.Slot{
			Name: weaponSlot,
		},
	}
}

const (
	warlockCategory = 21
	titanCategory   = 22
	hunterCategory  = 23
)

func NewArmorInstance(characterId, instanceId string, inventoryItemId int, destiny data.Content) (out bungo.Armor) {
	out = NewArmor(characterId, inventoryItemId, destiny)
	out.InstanceId = instanceId
	return out
}

func NewArmor(characterId string, inventoryItemId int, destiny data.Content) bungo.Armor {
	var armorClass string
	itemDef := destiny.InventoryItem.Find(inventoryItemId)
	itemBucketDef := destiny.InventoryBucket.Find(itemDef.Inventory.BucketTypeHash)
	for _, itemCategoryHash := range itemDef.ItemCategoryHashes {
		itemCategoryDef := destiny.ItemCategory.Find(itemCategoryHash)
		if itemCategoryHash == warlockCategory ||
			itemCategoryHash == titanCategory ||
			itemCategoryHash == hunterCategory {
			armorClass = itemCategoryDef.ShortTitle
		}
	}
	armorBucket := itemBucketDef.DisplayProperties.Name
	armorSlot := itemBucketDef.DisplayProperties.Name
	return bungo.Armor{
		MembershipType:  0,
		CharacterId:     characterId,
		InstanceId:      fmt.Sprintf("instance-of-%d", inventoryItemId),
		InventoryItemId: inventoryItemId,

		Tier:        itemDef.Inventory.TierTypeName,
		Bucket:      armorBucket,
		Name:        itemDef.DisplayProperties.Name,
		Type:        armorClass,
		IconEnabled: itemDef.DisplayProperties.HasIcon,
		IconUri:     itemDef.DisplayProperties.Icon,
		Slot: bungo.Slot{
			Name: armorSlot,
		},
	}
}

// https://www.fanbyte.com/guides/destiny-2-best-crucible-weapons/
var bestCrucibleWeaponsIds = []int{
	aceOfSpadesHandCannon,
	theLastWordHandCannon,
	thornHandCannon,
	spareRationsHandCannon,
	revokerSniperRifle,
	midaMultiToolScoutRifle,
	jadeRabbitScoutRifle,
	bygonesPulseRifle,
	dustRockBluesShotgun,
	theRecluseSubmachineGun,
	riskrunnerSubmachineGun,
	erentilFr4FusionRifle,
	loadedQuestionFusionRifle,
	mindbendersAmbitionShotgun,
	lunasHowlHandCannon,
	notForgottenHandCannon,
	trustHandCannon,
	erianasVowHandCannon,
	belovedSniperRifle,
	lastPerditionPulseRifle,
	hammerheadMachinegun,
	thunderlordMachinegun,
	n21deliriumMachinegun,
	aFineMemorialMachinegun,
	wendigoGl3GrenadeLauncher,
	playOfTheGameGrenadeLauncher,
	edgeTransitGrenadeLauncher,
	theColonyGrenadeLauncher,
	throughFireAndFloodGrenadeLauncher,
	truthRocketLauncher,
	theWardCliffCoilRocketLauncher,
	deathbringerRocketLauncher,
	zenobiaDRockerLauncher,
}

const (
	aceOfSpadesHandCannon              = 347366834
	theLastWordHandCannon              = 1364093401
	thornHandCannon                    = 3973202132
	spareRationsHandCannon             = 3116356268
	revokerSniperRifle                 = 654608616
	midaMultiToolScoutRifle            = 1331482397
	jadeRabbitScoutRifle               = 546372301
	bygonesPulseRifle                  = 2712244741
	dustRockBluesShotgun               = 636912560
	theRecluseSubmachineGun            = 3354242550
	riskrunnerSubmachineGun            = 3089417789
	erentilFr4FusionRifle              = 3027844941
	loadedQuestionFusionRifle          = 580961571
	mindbendersAmbitionShotgun         = 4117693024
	lunasHowlHandCannon                = 153979396
	notForgottenHandCannon             = 153979399
	trustHandCannon                    = 695814388
	erianasVowHandCannon               = 3524313097
	belovedSniperRifle                 = 4190932264
	lastPerditionPulseRifle            = 188882152
	hammerheadMachinegun               = 603242241
	thunderlordMachinegun              = 3325463374
	n21deliriumMachinegun              = 1600633250
	aFineMemorialMachinegun            = 3325778512
	wendigoGl3GrenadeLauncher          = 578459533
	playOfTheGameGrenadeLauncher       = 195440257
	edgeTransitGrenadeLauncher         = 218335759
	theColonyGrenadeLauncher           = 3899270607
	throughFireAndFloodGrenadeLauncher = 301362381
	truthRocketLauncher                = 1201830623
	theWardCliffCoilRocketLauncher     = 1508896098
	deathbringerRocketLauncher         = 2232171099
	zenobiaDRockerLauncher             = 3776129137

	imperativeScoutRifle              = 2314999489
	theHuckleBerrySubmachineGun       = 2286143274
	smugglersWordSidearm              = 1843044399
	tranquilitySniperRifle            = 1645386487
	theCutAndRunScoutRifle            = 3501969491
	optativeHandCannon                = 2138599001
	badlanderShotgun                  = 1327264046
	subjunctiveSubmachineGun          = 2663204025
	hardLightAutoRifle                = 4124984448
	theQueensbreakerLinearFusionRifle = 2044500762
	loveAndDeathGrenadeLauncher       = 3690523502
)

func bestCrucibleWeapons(characterId string, destiny data.Content) (out []bungo.Weapon) {
	for _, inventoryItemId := range bestCrucibleWeaponsIds {
		weapon := NewWeapon(characterId, inventoryItemId, destiny)
		log.Trace("adding to bag: ", weapon.Name)
		out = append(out, weapon)
	}
	return out
}

var hunterBagIds = []int{
	dreambaneCowlHelmet,
	substitutionalAlloyMaskHelmet,
	substitutionalAlloyGripsGauntlets,
	theSixthCoyoteChest,
	substitutionalAlloyVestChest,
	swordflight41Chest,
	substitutionalAlloyStridesLegs,
	prodigalCloakClass,
	substitutionalAlloyCloak,
}

const (

	// bag
	dreambaneCowlHelmet               = 659922705
	substitutionalAlloyMaskHelmet     = 4078925540
	substitutionalAlloyGripsGauntlets = 4026120124
	theSixthCoyoteChest               = 1474735277
	substitutionalAlloyVestChest      = 1855720514
	swordflight41Chest                = 328902054
	substitutionalAlloyStridesLegs    = 2096778462
	prodigalCloakClass                = 1693706589
	substitutionalAlloyCloak          = 1137424312
)

func hunterBag(characterId string, destiny data.Content) (out []bungo.Armor) {
	for _, inventoryItemId := range hunterBagIds {
		armor := NewArmor(characterId, inventoryItemId, destiny)
		log.Trace("adding to bag: ", armor.Name)
		out = append(out, armor)
	}
	return out
}

func newBag(characterId string, destiny data.Content) (bag bungo.Bag) {

	weapons := bestCrucibleWeapons(characterId, destiny)
	log.Debug("Bag Weapons: ", len(weapons))

	armors := hunterBag(characterId, destiny)
	log.Debug("Bag Armors: ", len(armors))

	switch characterId {
	default:
		return bungo.Bag{
			Weapons: weapons,
			Armors:  armors,
		}
	}
}
