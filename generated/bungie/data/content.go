package data

type Content struct {
	Achievement                    AchievementDefinitions                    `json:"DestinyAchievementDefinition,omitempty"`
	Activity                       ActivityDefinitions                       `json:"DestinyActivityDefinition,omitempty"`
	ActivityGraph                  ActivityGraphDefinitions                  `json:"DestinyActivityGraphDefinition,omitempty"`
	ActivityInteractable           ActivityInteractableDefinitions           `json:"DestinyActivityInteractableDefinition,omitempty"`
	ActivityMode                   ActivityModeDefinitions                   `json:"DestinyActivityModeDefinition,omitempty"`
	ActivityModifier               ActivityModifierDefinitions               `json:"DestinyActivityModifierDefinition,omitempty"`
	ActivityType                   ActivityTypeDefinitions                   `json:"DestinyActivityTypeDefinition,omitempty"`
	ArtDyeChannel                  ArtDyeChannelDefinitions                  `json:"DestinyArtDyeChannelDefinition,omitempty"`
	ArtDyeReference                ArtDyeReferenceDefinitions                `json:"DestinyArtDyeReferenceDefinition,omitempty"`
	Artifact                       ArtifactDefinitions                       `json:"DestinyArtifactDefinition,omitempty"`
	Bond                           BondDefinitions                           `json:"DestinyBondDefinition,omitempty"`
	BreakerType                    BreakerTypeDefinitions                    `json:"DestinyBreakerTypeDefinition,omitempty"`
	CharacterCustomizationCategory CharacterCustomizationCategoryDefinitions `json:"DestinyCharacterCustomizationCategoryDefinition,omitempty"`
	CharacterCustomizationOption   CharacterCustomizationOptionDefinitions   `json:"DestinyCharacterCustomizationOptionDefinition,omitempty"`
	Checklist                      ChecklistDefinitions                      `json:"DestinyChecklistDefinition,omitempty"`
	Class                          ClassDefinitions                          `json:"DestinyClassDefinition,omitempty"`
	Collectible                    CollectibleDefinitions                    `json:"DestinyCollectibleDefinition,omitempty"`
	DamageType                     DamageTypeDefinitions                     `json:"DestinyDamageTypeDefinition,omitempty"`
	Destination                    DestinationDefinitions                    `json:"DestinyDestinationDefinition,omitempty"`
	EnemyRace                      EnemyRaceDefinitions                      `json:"DestinyEnemyRaceDefinition,omitempty"`
	EnergyType                     EnergyTypeDefinitions                     `json:"DestinyEnergyTypeDefinition,omitempty"`
	EntitlementOffer               EntitlementOfferDefinitions               `json:"DestinyEntitlementOfferDefinition,omitempty"`
	EquipmentSlot                  EquipmentSlotDefinitions                  `json:"DestinyEquipmentSlotDefinition,omitempty"`
	Faction                        FactionDefinitions                        `json:"DestinyFactionDefinition,omitempty"`
	Gender                         GenderDefinitions                         `json:"DestinyGenderDefinition,omitempty"`
	InventoryBucket                InventoryBucketDefinitions                `json:"DestinyInventoryBucketDefinition,omitempty"`
	InventoryItem                  InventoryItemDefinitions                  `json:"DestinyInventoryItemDefinition,omitempty"`
	ItemCategory                   ItemCategoryDefinitions                   `json:"DestinyItemCategoryDefinition,omitempty"`
	ItemTierType                   ItemTierTypeDefinitions                   `json:"DestinyItemTierTypeDefinition,omitempty"`
	Location                       LocationDefinitions                       `json:"DestinyLocationDefinition,omitempty"`
	Lore                           LoreDefinitions                           `json:"DestinyLoreDefinition,omitempty"`
	MaterialRequirementSet         MaterialRequirementSetDefinitions         `json:"DestinyMaterialRequirementSetDefinition,omitempty"`
	MedalTier                      MedalTierDefinitions                      `json:"DestinyMedalTierDefinition,omitempty"`
	Milestone                      MilestoneDefinitions                      `json:"DestinyMilestoneDefinition,omitempty"`
	NodeStepSummary                NodeStepSummaryDefinitions                `json:"DestinyNodeStepSummaryDefinition,omitempty"`
	Objective                      ObjectiveDefinitions                      `json:"DestinyObjectiveDefinition,omitempty"`
	Place                          PlaceDefinitions                          `json:"DestinyPlaceDefinition,omitempty"`
	PlatformBucketMapping          PlatformBucketMappingDefinitions          `json:"DestinyPlatformBucketMappingDefinition,omitempty"`
	PlugSet                        PlugSetDefinitions                        `json:"DestinyPlugSetDefinition,omitempty"`
	PresentationNode               PresentationNodeDefinitions               `json:"DestinyPresentationNodeDefinition,omitempty"`
	Progression                    ProgressionDefinitions                    `json:"DestinyProgressionDefinition,omitempty"`
	ProgressionLevelRequirement    ProgressionLevelRequirementDefinitions    `json:"DestinyProgressionLevelRequirementDefinition,omitempty"`
	ProgressionMapping             ProgressionMappingDefinitions             `json:"DestinyProgressionMappingDefinition,omitempty"`
	Race                           RaceDefinitions                           `json:"DestinyRaceDefinition,omitempty"`
	Record                         RecordDefinitions                         `json:"DestinyRecordDefinition,omitempty"`
	ReportReasonCategory           ReportReasonCategoryDefinitions           `json:"DestinyReportReasonCategoryDefinition,omitempty"`
	RewardAdjusterPointer          RewardAdjusterPointerDefinitions          `json:"DestinyRewardAdjusterPointerDefinition,omitempty"`
	RewardAdjusterProgressionMap   RewardAdjusterProgressionMapDefinitions   `json:"DestinyRewardAdjusterProgressionMapDefinition,omitempty"`
	RewardItemList                 RewardItemListDefinitions                 `json:"DestinyRewardItemListDefinition,omitempty"`
	RewardMapping                  RewardMappingDefinitions                  `json:"DestinyRewardMappingDefinition,omitempty"`
	RewardSheet                    RewardSheetDefinitions                    `json:"DestinyRewardSheetDefinition,omitempty"`
	RewardSource                   RewardSourceDefinitions                   `json:"DestinyRewardSourceDefinition,omitempty"`
	SackRewardItemList             SackRewardItemListDefinitions             `json:"DestinySackRewardItemListDefinition,omitempty"`
	SandboxPattern                 SandboxPatternDefinitions                 `json:"DestinySandboxPatternDefinition,omitempty"`
	SandboxPerk                    SandboxPerkDefinitions                    `json:"DestinySandboxPerkDefinition,omitempty"`
	Season                         SeasonDefinitions                         `json:"DestinySeasonDefinition,omitempty"`
	SeasonPass                     SeasonPassDefinitions                     `json:"DestinySeasonPassDefinition,omitempty"`
	SocketCategory                 SocketCategoryDefinitions                 `json:"DestinySocketCategoryDefinition,omitempty"`
	SocketType                     SocketTypeDefinitions                     `json:"DestinySocketTypeDefinition,omitempty"`
	Stat                           StatDefinitions                           `json:"DestinyStatDefinition,omitempty"`
	StatGroup                      StatGroupDefinitions                      `json:"DestinyStatGroupDefinition,omitempty"`
	TalentGrid                     TalentGridDefinitions                     `json:"DestinyTalentGridDefinition,omitempty"`
	UnlockCountMapping             UnlockCountMappingDefinitions             `json:"DestinyUnlockCountMappingDefinition,omitempty"`
	Unlock                         UnlockDefinitions                         `json:"DestinyUnlockDefinition,omitempty"`
	UnlockEvent                    UnlockEventDefinitions                    `json:"DestinyUnlockEventDefinition,omitempty"`
	UnlockExpressionMapping        UnlockExpressionMappingDefinitions        `json:"DestinyUnlockExpressionMappingDefinition,omitempty"`
	UnlockValue                    UnlockValueDefinitions                    `json:"DestinyUnlockValueDefinition,omitempty"`
	Vendor                         VendorDefinitions                         `json:"DestinyVendorDefinition,omitempty"`
	VendorGroup                    VendorGroupDefinitions                    `json:"DestinyVendorGroupDefinition,omitempty"`
}

type Definition interface {
	Name() string
	Json() ([]byte, error)
	PrettyJson() ([]byte, error)
	Yaml() ([]byte, error)
	Hydrate(content Content)
}
