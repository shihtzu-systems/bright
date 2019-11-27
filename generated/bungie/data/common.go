// Code generated manually; DANGER ZONE FOR EDITS

package data

type DisplayProperties struct {
	Description         string        `json:"description,omitempty"`
	HasIcon             bool          `json:"hasIcon,omitempty"`
	Icon                string        `json:"icon,omitempty"`
	Name                string        `json:"name,omitempty"`
	OriginalIcon        string        `json:"originalIcon,omitempty"`
	RequirementsDisplay []interface{} `json:"requirementsDisplay,omitempty"`
	Subtitle            string        `json:"subtitle,omitempty"`
}

type Arrangements struct {
	ArtArrangementHash int `json:"artArrangementHash,omitempty"`
	ClassHash          int `json:"classHash,omitempty"`
}
type DefaultDyes struct {
	ChannelHash int `json:"channelHash,omitempty"`
	DyeHash     int `json:"dyeHash,omitempty"`
}
type TranslationBlock struct {
	Arrangements      []Arrangements `json:"arrangements,omitempty"`
	CustomDyes        []interface{}  `json:"customDyes,omitempty"`
	DefaultDyes       []DefaultDyes  `json:"defaultDyes,omitempty"`
	HasGeometry       bool           `json:"hasGeometry,omitempty"`
	LockedDyes        []interface{}  `json:"lockedDyes,omitempty"`
	WeaponPatternHash int64          `json:"weaponPatternHash,omitempty"`
}
