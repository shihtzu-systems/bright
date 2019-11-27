package brightsvc

const (
	// enums
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

	warlockCategory = 21
	titanCategory   = 22
	hunterCategory  = 23
)

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

	// best crucible weapons
	// https://www.fanbyte.com/guides/destiny-2-best-crucible-weapons/
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

var emblemHashes = []int{
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

var bestCrucibleWeaponHashes = []int{
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

var kineticWeaponHashes = []int{
	dukeMk44HandCannon,
	brayTechWerewolfAutoRifle,
	sacredProvenancePulseRifle,
	guidingStarAutoRifle,
	peaceByConsensusSidearm,
	frontierJusticeScoutRifle,
	theSpitefulFangBow,
}

var energyWeaponHashes = []int{
	borealisSniperRifle,
	drangBaroqueSideArm,
	nationOfBeastsHandcannon,
	impromptu49PulseRifle,
}

var powerWeaponHashes = []int{
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

var hunterHelmetHashes = []int{
	masqueradersCowlHelmet,
}

var hunterGauntletHashes = []int{
	dreambaneGripsGauntlets,
}

var hunterChestHashes = []int{
	dreambaneVestChest,
}

var hunterLegHashes = []int{
	dreambaneStridesLegs,
}

var hunterClassHashes = []int{
	dreambaneCloakClass,
}

var hunterArmorHashes = []int{
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
	// titan
	shadowsHelmHelmet         = 1129634130
	shadowsGauntletsGauntlets = 1595987387
	shadowsPlateChect         = 1862963733
	shadowsGreavesLegs        = 4450861
	shadowsMarkClass          = 311429764
)

var titanHelmetHashes = []int{
	shadowsHelmHelmet,
}

var titanGauntletHashes = []int{
	shadowsGauntletsGauntlets,
}

var titanChestHashes = []int{
	shadowsPlateChect,
}

var titanLegHashes = []int{
	shadowsGreavesLegs,
}

var titanClassHashes = []int{
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

var warlockHelmetHashes = []int{
	solsticeHoodMajesticHelmet,
}

var warlockGauntletHashes = []int{
	solsticeGlovesMajesticGauntlets,
}

var warlockChestHashes = []int{
	solsticeRobesMajesticChest,
}

var warlockLegHashes = []int{
	solsticeBootsMajesticLegs,
}

var warlockClassHashes = []int{
	solsticeBondMajesticClass,
}
