// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const InventoryItemDefinitionName = "inventory-item"

type InventoryItemDefinitions map[string]InventoryItemDefinition

func (d InventoryItemDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d InventoryItemDefinitions) Values() (out []InventoryItemDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d InventoryItemDefinitions) Find(id int) (out InventoryItemDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return InventoryItemDefinition{}
}

func (d InventoryItemDefinitions) Name() string {
	return InventoryItemDefinitionName
}

type InventoryItemDefinition struct {
	AcquireRewardSiteHash             int                 `json:"acquireRewardSiteHash" yaml:"acquireRewardSiteHash,omitempty"`
	AcquireUnlockHash                 int                 `json:"acquireUnlockHash" yaml:"acquireUnlockHash,omitempty"`
	Action                            InventoryItemAction `json:"action" yaml:"action,omitempty"`
	AllowActions                      bool                `json:"allowActions" yaml:"allowActions,omitempty"`
	BackgroundColor                   BackgroundColor     `json:"backgroundColor" yaml:"backgroundColor,omitempty"`
	Blacklisted                       bool                `json:"blacklisted" yaml:"blacklisted,omitempty"`
	BreakerType                       int                 `json:"breakerType" yaml:"breakerType,omitempty"`
	ClassType                         int                 `json:"classType" yaml:"classType,omitempty"`
	DefaultDamageType                 int                 `json:"defaultDamageType" yaml:"defaultDamageType,omitempty"`
	DisplayProperties                 DisplayProperties   `json:"displayProperties" yaml:"displayProperties,omitempty"`
	DisplaySource                     string              `json:"displaySource" yaml:"displaySource,omitempty"`
	DoesPostmasterPullHaveSideEffects bool                `json:"doesPostmasterPullHaveSideEffects" yaml:"doesPostmasterPullHaveSideEffects,omitempty"`
	Equippable                        bool                `json:"equippable" yaml:"equippable,omitempty"`
	EquippingBlock                    EquippingBlock      `json:"equippingBlock" yaml:"equippingBlock,omitempty"`
	Hash                              int                 `json:"hash" yaml:"hash,omitempty"`
	Index                             int                 `json:"index" yaml:"index,omitempty"`
	Inventory                         Inventory           `json:"inventory" yaml:"inventory,omitempty"`
	InvestmentStats                   []InvestmentStats   `json:"investmentStats" yaml:"investmentStats,omitempty"`
	IsWrapper                         bool                `json:"isWrapper" yaml:"isWrapper,omitempty"`
	ItemCategoryHashes                []int               `json:"itemCategoryHashes" yaml:"itemCategoryHashes,omitempty"`
	ItemSubType                       int                 `json:"itemSubType" yaml:"itemSubType,omitempty"`
	ItemType                          int                 `json:"itemType" yaml:"itemType,omitempty"`
	ItemTypeAndTierDisplayName        string              `json:"itemTypeAndTierDisplayName" yaml:"itemTypeAndTierDisplayName,omitempty"`
	ItemTypeDisplayName               string              `json:"itemTypeDisplayName" yaml:"itemTypeDisplayName,omitempty"`
	NonTransferrable                  bool                `json:"nonTransferrable" yaml:"nonTransferrable,omitempty"`
	Perks                             []interface{}       `json:"perks" yaml:"perks,omitempty"`
	Preview                           Preview             `json:"preview" yaml:"preview,omitempty"`
	Quality                           Quality             `json:"quality" yaml:"quality,omitempty"`
	Redacted                          bool                `json:"redacted" yaml:"redacted,omitempty"`
	SecondaryIcon                     string              `json:"secondaryIcon" yaml:"secondaryIcon,omitempty"`
	SecondaryOverlay                  string              `json:"secondaryOverlay" yaml:"secondaryOverlay,omitempty"`
	SecondarySpecial                  string              `json:"secondarySpecial" yaml:"secondarySpecial,omitempty"`
	Screenshot                        string              `json:"screenshot" yaml:"screenshot,omitempty"`
	Sockets                           Sockets             `json:"sockets" yaml:"sockets,omitempty"`
	SourceData                        SourceData          `json:"sourceData" yaml:"sourceData,omitempty"`
	SpecialItemType                   int                 `json:"specialItemType" yaml:"specialItemType,omitempty"`
	Stats                             Stats               `json:"stats" yaml:"stats,omitempty"`
	SummaryItemHash                   int64               `json:"summaryItemHash" yaml:"summaryItemHash,omitempty"`
	TalentGrid                        TalentGrid          `json:"talentGrid" yaml:"talentGrid,omitempty"`
	TooltipNotifications              []interface{}       `json:"tooltipNotifications" yaml:"tooltipNotifications,omitempty"`
	TranslationBlock                  TranslationBlock    `json:"translationBlock" yaml:"translationBlock,omitempty"`
	UiItemDisplayStyle                string              `json:"uiItemDisplayStyle" yaml:"uiItemDisplayStyle,omitempty"`
}
type InventoryItemAction struct {
	ActionTypeLabel         string        `json:"actionTypeLabel" yaml:"actionTypeLabel,omitempty"`
	ConsumeEntireStack      bool          `json:"consumeEntireStack" yaml:"consumeEntireStack,omitempty"`
	DeleteOnAction          bool          `json:"deleteOnAction" yaml:"deleteOnAction,omitempty"`
	IsPositive              bool          `json:"isPositive" yaml:"isPositive,omitempty"`
	ProgressionRewards      []interface{} `json:"progressionRewards" yaml:"progressionRewards,omitempty"`
	RequiredCooldownHash    int           `json:"requiredCooldownHash" yaml:"requiredCooldownHash,omitempty"`
	RequiredCooldownSeconds int           `json:"requiredCooldownSeconds" yaml:"requiredCooldownSeconds,omitempty"`
	RequiredItems           []interface{} `json:"requiredItems" yaml:"requiredItems,omitempty"`
	RewardItemHash          int           `json:"rewardItemHash" yaml:"rewardItemHash,omitempty"`
	RewardSheetHash         int           `json:"rewardSheetHash" yaml:"rewardSheetHash,omitempty"`
	RewardSiteHash          int           `json:"rewardSiteHash" yaml:"rewardSiteHash,omitempty"`
	UseOnAcquire            bool          `json:"useOnAcquire" yaml:"useOnAcquire,omitempty"`
	VerbDescription         string        `json:"verbDescription" yaml:"verbDescription,omitempty"`
	VerbName                string        `json:"verbName" yaml:"verbName,omitempty"`
}
type BackgroundColor struct {
	Alpha     int `json:"alpha" yaml:"alpha,omitempty"`
	Blue      int `json:"blue" yaml:"blue,omitempty"`
	ColorHash int `json:"colorHash" yaml:"colorHash,omitempty"`
	Green     int `json:"green" yaml:"green,omitempty"`
	Red       int `json:"red" yaml:"red,omitempty"`
}
type EquippingBlock struct {
	AmmoType              int      `json:"ammoType" yaml:"ammoType,omitempty"`
	Attributes            int      `json:"attributes" yaml:"attributes,omitempty"`
	DisplayStrings        []string `json:"displayStrings" yaml:"displayStrings,omitempty"`
	EquipmentSlotTypeHash int      `json:"equipmentSlotTypeHash" yaml:"equipmentSlotTypeHash,omitempty"`
	EquippingSoundHash    int      `json:"equippingSoundHash" yaml:"equippingSoundHash,omitempty"`
	HornSoundHash         int      `json:"hornSoundHash" yaml:"hornSoundHash,omitempty"`
	UniqueLabelHash       int      `json:"uniqueLabelHash" yaml:"uniqueLabelHash,omitempty"`
}
type Inventory struct {
	BucketTypeHash                           int    `json:"bucketTypeHash" yaml:"bucketTypeHash,omitempty"`
	ExpirationTooltip                        string `json:"expirationTooltip" yaml:"expirationTooltip,omitempty"`
	ExpiredInActivityMessage                 string `json:"expiredInActivityMessage" yaml:"expiredInActivityMessage,omitempty"`
	ExpiredInOrbitMessage                    string `json:"expiredInOrbitMessage" yaml:"expiredInOrbitMessage,omitempty"`
	IsInstanceItem                           bool   `json:"isInstanceItem" yaml:"isInstanceItem,omitempty"`
	MaxStackSize                             int    `json:"maxStackSize" yaml:"maxStackSize,omitempty"`
	NonTransferrableOriginal                 bool   `json:"nonTransferrableOriginal" yaml:"nonTransferrableOriginal,omitempty"`
	RecoveryBucketTypeHash                   int    `json:"recoveryBucketTypeHash" yaml:"recoveryBucketTypeHash,omitempty"`
	SuppressExpirationWhenObjectivesComplete bool   `json:"suppressExpirationWhenObjectivesComplete" yaml:"suppressExpirationWhenObjectivesComplete,omitempty"`
	TierType                                 int    `json:"tierType" yaml:"tierType,omitempty"`
	TierTypeHash                             int64  `json:"tierTypeHash" yaml:"tierTypeHash,omitempty"`
	TierTypeName                             string `json:"tierTypeName" yaml:"tierTypeName,omitempty"`
}
type InvestmentStats struct {
	IsConditionallyActive bool  `json:"isConditionallyActive" yaml:"isConditionallyActive,omitempty"`
	StatTypeHash          int64 `json:"statTypeHash" yaml:"statTypeHash,omitempty"`
	Value                 int   `json:"value" yaml:"value,omitempty"`
}
type Preview struct {
	PreviewActionString string `json:"previewActionString" yaml:"previewActionString,omitempty"`
	PreviewVendorHash   int    `json:"previewVendorHash" yaml:"previewVendorHash,omitempty"`
	ScreenStyle         string `json:"screenStyle" yaml:"screenStyle,omitempty"`
}
type Quality struct {
	InfusionCategoryHash            int64         `json:"infusionCategoryHash" yaml:"infusionCategoryHash,omitempty"`
	InfusionCategoryHashes          []int64       `json:"infusionCategoryHashes" yaml:"infusionCategoryHashes,omitempty"`
	InfusionCategoryName            string        `json:"infusionCategoryName" yaml:"infusionCategoryName,omitempty"`
	ItemLevels                      []interface{} `json:"itemLevels" yaml:"itemLevels,omitempty"`
	ProgressionLevelRequirementHash int64         `json:"progressionLevelRequirementHash" yaml:"progressionLevelRequirementHash,omitempty"`
	QualityLevel                    int           `json:"qualityLevel" yaml:"qualityLevel,omitempty"`
}
type IntrinsicSockets struct {
	DefaultVisible bool  `json:"defaultVisible" yaml:"defaultVisible,omitempty"`
	PlugItemHash   int64 `json:"plugItemHash" yaml:"plugItemHash,omitempty"`
	SocketTypeHash int   `json:"socketTypeHash" yaml:"socketTypeHash,omitempty"`
}
type SocketCategories struct {
	SocketCategoryHash int   `json:"socketCategoryHash" yaml:"socketCategoryHash,omitempty"`
	SocketIndexes      []int `json:"socketIndexes" yaml:"socketIndexes,omitempty"`
}
type SocketEntries struct {
	DefaultVisible                        bool          `json:"defaultVisible" yaml:"defaultVisible,omitempty"`
	HidePerksInItemTooltip                bool          `json:"hidePerksInItemTooltip" yaml:"hidePerksInItemTooltip,omitempty"`
	OverridesUiAppearance                 bool          `json:"overridesUiAppearance" yaml:"overridesUiAppearance,omitempty"`
	PlugSources                           int           `json:"plugSources" yaml:"plugSources,omitempty"`
	PreventInitializationOnVendorPurchase bool          `json:"preventInitializationOnVendorPurchase" yaml:"preventInitializationOnVendorPurchase,omitempty"`
	PreventInitializationWhenVersioning   bool          `json:"preventInitializationWhenVersioning" yaml:"preventInitializationWhenVersioning,omitempty"`
	ReusablePlugItems                     []interface{} `json:"reusablePlugItems" yaml:"reusablePlugItems,omitempty"`
	ReusablePlugSetHash                   int64         `json:"reusablePlugSetHash" yaml:"reusablePlugSetHash,omitempty"`
	SingleInitialItemHash                 int           `json:"singleInitialItemHash" yaml:"singleInitialItemHash,omitempty"`
	SocketTypeHash                        int           `json:"socketTypeHash" yaml:"socketTypeHash,omitempty"`
}
type Sockets struct {
	Detail           string             `json:"detail" yaml:"detail,omitempty"`
	IntrinsicSockets []IntrinsicSockets `json:"intrinsicSockets" yaml:"intrinsicSockets,omitempty"`
	SocketCategories []SocketCategories `json:"socketCategories" yaml:"socketCategories,omitempty"`
	SocketEntries    []SocketEntries    `json:"socketEntries" yaml:"socketEntries,omitempty"`
}
type SourceData struct {
	Exclusive     int           `json:"exclusive" yaml:"exclusive,omitempty"`
	VendorSources []interface{} `json:"vendorSources" yaml:"vendorSources,omitempty"`
}
type Stats struct {
	DisablePrimaryStatDisplay bool                   `json:"disablePrimaryStatDisplay" yaml:"disablePrimaryStatDisplay,omitempty"`
	HasDisplayableStats       bool                   `json:"hasDisplayableStats" yaml:"hasDisplayableStats,omitempty"`
	PrimaryBaseStatHash       int64                  `json:"primaryBaseStatHash" yaml:"primaryBaseStatHash,omitempty"`
	StatGroupHash             int64                  `json:"statGroupHash" yaml:"statGroupHash,omitempty"`
	Stats                     map[string]interface{} `json:"stats" yaml:"stats,omitempty"`
}
type TalentGrid struct {
	HudDamageType    int    `json:"hudDamageType" yaml:"hudDamageType,omitempty"`
	ItemDetailString string `json:"itemDetailString" yaml:"itemDetailString,omitempty"`
	TalentGridHash   int    `json:"talentGridHash" yaml:"talentGridHash,omitempty"`
}

func (d InventoryItemDefinition) Name() string {
	return InventoryItemDefinitionName
}

func (d InventoryItemDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d InventoryItemDefinition) PrettyJson() ([]byte, error) {
	jout, err := d.Json()
	if err != nil {
		return nil, err
	}
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, jout, "", "  "); err != nil {
		return nil, err
	}
	return pretty.Bytes(), nil
}

func (d InventoryItemDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
