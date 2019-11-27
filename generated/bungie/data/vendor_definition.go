// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const VendorDefinitionName = "vendor"

type VendorDefinitions map[string]VendorDefinition

func (d VendorDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d VendorDefinitions) Values() (out []VendorDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d VendorDefinitions) Find(id int) (out VendorDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return VendorDefinition{}
}

func (d VendorDefinitions) Name() string {
	return VendorDefinitionName
}

type VendorDefinition struct {
	AcceptedItems           []interface{}        `json:"acceptedItems" yaml:"acceptedItems,omitempty"`
	Actions                 []interface{}        `json:"actions" yaml:"actions,omitempty"`
	Blacklisted             bool                 `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Categories              []Categories         `json:"categories" yaml:"categories,omitempty"`
	ConsolidateCategories   bool                 `json:"consolidateCategories" yaml:"consolidateCategories,omitempty"`
	DisplayCategories       []DisplayCategories  `json:"displayCategories" yaml:"displayCategories,omitempty"`
	DisplayItemHash         int                  `json:"displayItemHash" yaml:"displayItemHash,omitempty"`
	DisplayProperties       DisplayProperties    `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Enabled                 bool                 `json:"enabled" yaml:"enabled,omitempty"`
	FactionHash             int                  `json:"factionHash" yaml:"factionHash,omitempty"`
	FailureStrings          []interface{}        `json:"failureStrings" yaml:"failureStrings,omitempty"`
	Hash                    int                  `json:"hash" yaml:"hash,omitempty"`
	IgnoreSaleItemHashes    []interface{}        `json:"ignoreSaleItemHashes" yaml:"ignoreSaleItemHashes,omitempty"`
	Index                   int                  `json:"index" yaml:"index,omitempty"`
	InhibitBuying           bool                 `json:"inhibitBuying" yaml:"inhibitBuying,omitempty"`
	InhibitSelling          bool                 `json:"inhibitSelling" yaml:"inhibitSelling,omitempty"`
	Interactions            []interface{}        `json:"interactions" yaml:"interactions,omitempty"`
	InventoryFlyouts        []interface{}        `json:"inventoryFlyouts" yaml:"inventoryFlyouts,omitempty"`
	ItemList                []ItemList           `json:"itemList" yaml:"itemList,omitempty"`
	OriginalCategories      []OriginalCategories `json:"originalCategories" yaml:"originalCategories,omitempty"`
	Redacted                bool                 `json:"redacted" yaml:"redacted,omitempty"`
	ResetIntervalMinutes    int                  `json:"resetIntervalMinutes" yaml:"resetIntervalMinutes,omitempty"`
	ResetOffsetMinutes      int                  `json:"resetOffsetMinutes" yaml:"resetOffsetMinutes,omitempty"`
	ReturnWithVendorRequest bool                 `json:"returnWithVendorRequest" yaml:"returnWithVendorRequest,omitempty"`
	Services                []interface{}        `json:"services" yaml:"services,omitempty"`
	UnlockRanges            []interface{}        `json:"unlockRanges" yaml:"unlockRanges,omitempty"`
	UnlockValueHash         int                  `json:"unlockValueHash" yaml:"unlockValueHash,omitempty"`
	Visible                 bool                 `json:"visible" yaml:"visible,omitempty"`
}
type Categories struct {
	BuyStringOverride            string `json:"buyStringOverride" yaml:"buyStringOverride,omitempty"`
	CategoryHash                 int    `json:"categoryHash" yaml:"categoryHash,omitempty"`
	CategoryIndex                int    `json:"categoryIndex" yaml:"categoryIndex,omitempty"`
	DisabledDescription          string `json:"disabledDescription" yaml:"disabledDescription,omitempty"`
	HideFromRegularPurchase      bool   `json:"hideFromRegularPurchase" yaml:"hideFromRegularPurchase,omitempty"`
	HideIfNoCurrency             bool   `json:"hideIfNoCurrency" yaml:"hideIfNoCurrency,omitempty"`
	IsDisplayOnly                bool   `json:"isDisplayOnly" yaml:"isDisplayOnly,omitempty"`
	IsPreview                    bool   `json:"isPreview" yaml:"isPreview,omitempty"`
	QuantityAvailable            int    `json:"quantityAvailable" yaml:"quantityAvailable,omitempty"`
	ResetIntervalMinutesOverride int    `json:"resetIntervalMinutesOverride" yaml:"resetIntervalMinutesOverride,omitempty"`
	ResetOffsetMinutesOverride   int    `json:"resetOffsetMinutesOverride" yaml:"resetOffsetMinutesOverride,omitempty"`
	ShowUnavailableItems         bool   `json:"showUnavailableItems" yaml:"showUnavailableItems,omitempty"`
	SortValue                    int    `json:"sortValue" yaml:"sortValue,omitempty"`
	VendorItemIndexes            []int  `json:"vendorItemIndexes" yaml:"vendorItemIndexes,omitempty"`
}
type DisplayCategories struct {
	DisplayCategoryHash int               `json:"displayCategoryHash" yaml:"displayCategoryHash,omitempty"`
	DisplayInBanner     bool              `json:"displayInBanner" yaml:"displayInBanner,omitempty"`
	DisplayProperties   DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Identifier          string            `json:"identifier" yaml:"identifier,omitempty"`
	Index               int               `json:"index" yaml:"index,omitempty"`
	SortOrder           int               `json:"sortOrder" yaml:"sortOrder,omitempty"`
}
type VendorAction struct {
	ExecuteSeconds float64 `json:"executeSeconds" yaml:"executeSeconds,omitempty"`
	IsPositive     bool    `json:"isPositive" yaml:"isPositive,omitempty"`
}
type CreationLevels struct {
	Level int `json:"level" yaml:"level,omitempty"`
}
type ItemList struct {
	Action                    VendorAction     `json:"action" yaml:"action,omitempty"`
	CategoryIndex             int              `json:"categoryIndex" yaml:"categoryIndex,omitempty"`
	CreationLevels            []CreationLevels `json:"creationLevels" yaml:"creationLevels,omitempty"`
	Currencies                []interface{}    `json:"currencies" yaml:"currencies,omitempty"`
	DisplayCategory           string           `json:"displayCategory" yaml:"displayCategory,omitempty"`
	DisplayCategoryIndex      int              `json:"displayCategoryIndex" yaml:"displayCategoryIndex,omitempty"`
	Exclusivity               int              `json:"exclusivity" yaml:"exclusivity,omitempty"`
	ExpirationTooltip         string           `json:"expirationTooltip" yaml:"expirationTooltip,omitempty"`
	FailureIndexes            []interface{}    `json:"failureIndexes" yaml:"failureIndexes,omitempty"`
	InventoryBucketHash       int64            `json:"inventoryBucketHash" yaml:"inventoryBucketHash,omitempty"`
	ItemHash                  int              `json:"itemHash" yaml:"itemHash,omitempty"`
	LicenseUnlockHash         int              `json:"licenseUnlockHash" yaml:"licenseUnlockHash,omitempty"`
	MaximumLevel              int              `json:"maximumLevel" yaml:"maximumLevel,omitempty"`
	MinimumLevel              int              `json:"minimumLevel" yaml:"minimumLevel,omitempty"`
	OriginalCategoryIndex     int              `json:"originalCategoryIndex" yaml:"originalCategoryIndex,omitempty"`
	PriceOverrideEnabled      bool             `json:"priceOverrideEnabled" yaml:"priceOverrideEnabled,omitempty"`
	PurchasableScope          int              `json:"purchasableScope" yaml:"purchasableScope,omitempty"`
	Quantity                  int              `json:"quantity" yaml:"quantity,omitempty"`
	RedirectToSaleIndexes     []interface{}    `json:"redirectToSaleIndexes" yaml:"redirectToSaleIndexes,omitempty"`
	RefundPolicy              int              `json:"refundPolicy" yaml:"refundPolicy,omitempty"`
	RefundTimeLimit           int              `json:"refundTimeLimit" yaml:"refundTimeLimit,omitempty"`
	RewardAdjustorPointerHash int              `json:"rewardAdjustorPointerHash" yaml:"rewardAdjustorPointerHash,omitempty"`
	SeedOverride              int              `json:"seedOverride" yaml:"seedOverride,omitempty"`
	SocketOverrides           []interface{}    `json:"socketOverrides" yaml:"socketOverrides,omitempty"`
	SortValue                 int              `json:"sortValue" yaml:"sortValue,omitempty"`
	VendorItemIndex           int              `json:"vendorItemIndex" yaml:"vendorItemIndex,omitempty"`
	VisibilityScope           int              `json:"visibilityScope" yaml:"visibilityScope,omitempty"`
	Weight                    float64          `json:"weight" yaml:"weight,omitempty"`
}
type OriginalCategories struct {
	BuyStringOverride            string `json:"buyStringOverride" yaml:"buyStringOverride,omitempty"`
	CategoryHash                 int    `json:"categoryHash" yaml:"categoryHash,omitempty"`
	CategoryIndex                int    `json:"categoryIndex" yaml:"categoryIndex,omitempty"`
	DisabledDescription          string `json:"disabledDescription" yaml:"disabledDescription,omitempty"`
	HideFromRegularPurchase      bool   `json:"hideFromRegularPurchase" yaml:"hideFromRegularPurchase,omitempty"`
	HideIfNoCurrency             bool   `json:"hideIfNoCurrency" yaml:"hideIfNoCurrency,omitempty"`
	IsDisplayOnly                bool   `json:"isDisplayOnly" yaml:"isDisplayOnly,omitempty"`
	IsPreview                    bool   `json:"isPreview" yaml:"isPreview,omitempty"`
	QuantityAvailable            int    `json:"quantityAvailable" yaml:"quantityAvailable,omitempty"`
	ResetIntervalMinutesOverride int    `json:"resetIntervalMinutesOverride" yaml:"resetIntervalMinutesOverride,omitempty"`
	ResetOffsetMinutesOverride   int    `json:"resetOffsetMinutesOverride" yaml:"resetOffsetMinutesOverride,omitempty"`
	ShowUnavailableItems         bool   `json:"showUnavailableItems" yaml:"showUnavailableItems,omitempty"`
	SortValue                    int    `json:"sortValue" yaml:"sortValue,omitempty"`
	VendorItemIndexes            []int  `json:"vendorItemIndexes" yaml:"vendorItemIndexes,omitempty"`
}

func (d VendorDefinition) Name() string {
	return VendorDefinitionName
}

func (d VendorDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d VendorDefinition) PrettyJson() ([]byte, error) {
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

func (d VendorDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
