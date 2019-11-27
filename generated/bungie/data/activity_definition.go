// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityDefinitionName = "activity"

type ActivityDefinitions map[string]ActivityDefinition

func (d ActivityDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityDefinitions) Values() (out []ActivityDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityDefinitions) Find(id int) (out ActivityDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityDefinition{}
}

func (d ActivityDefinitions) Name() string {
	return ActivityDefinitionName
}

type ActivityDefinition struct {
	ActivityLevel             int                       `json:"activityLevel" yaml:"activityLevel,omitempty"`
	ActivityLightLevel        int                       `json:"activityLightLevel" yaml:"activityLightLevel,omitempty"`
	ActivityLocationMappings  []interface{}             `json:"activityLocationMappings" yaml:"activityLocationMappings,omitempty"`
	ActivityModeHashes        []int                     `json:"activityModeHashes" yaml:"activityModeHashes,omitempty"`
	ActivityModeTypes         []int                     `json:"activityModeTypes" yaml:"activityModeTypes,omitempty"`
	ActivityTypeHash          int                       `json:"activityTypeHash" yaml:"activityTypeHash,omitempty"`
	Blacklisted               bool                      `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Challenges                []interface{}             `json:"challenges" yaml:"challenges,omitempty"`
	CompletionUnlockHash      int                       `json:"completionUnlockHash" yaml:"completionUnlockHash,omitempty"`
	DestinationHash           int                       `json:"destinationHash" yaml:"destinationHash,omitempty"`
	DirectActivityModeHash    int                       `json:"directActivityModeHash" yaml:"directActivityModeHash,omitempty"`
	DirectActivityModeType    int                       `json:"directActivityModeType" yaml:"directActivityModeType,omitempty"`
	DisplayProperties         DisplayProperties         `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                      int64                     `json:"hash" yaml:"hash,omitempty"`
	Index                     int                       `json:"index" yaml:"index,omitempty"`
	InheritFromFreeRoam       bool                      `json:"inheritFromFreeRoam" yaml:"inheritFromFreeRoam,omitempty"`
	InsertionPoints           []interface{}             `json:"insertionPoints" yaml:"insertionPoints,omitempty"`
	IsPlaylist                bool                      `json:"isPlaylist" yaml:"isPlaylist,omitempty"`
	IsPvP                     bool                      `json:"isPvP" yaml:"isPvP,omitempty"`
	Matchmaking               Matchmaking               `json:"matchmaking" yaml:"matchmaking,omitempty"`
	Modifiers                 []interface{}             `json:"modifiers" yaml:"modifiers,omitempty"`
	OptionalUnlockStrings     []interface{}             `json:"optionalUnlockStrings" yaml:"optionalUnlockStrings,omitempty"`
	OriginalDisplayProperties OriginalDisplayProperties `json:"originalDisplayProperties" yaml:"originalDisplayProperties,omitempty"`
	PgcrImage                 string                    `json:"pgcrImage" yaml:"pgcrImage,omitempty"`
	PlaceHash                 int64                     `json:"placeHash" yaml:"placeHash,omitempty"`
	PlaylistItems             []interface{}             `json:"playlistItems" yaml:"playlistItems,omitempty"`
	Redacted                  bool                      `json:"redacted" yaml:"redacted,omitempty"`
	ReleaseIcon               string                    `json:"releaseIcon" yaml:"releaseIcon,omitempty"`
	ReleaseTime               int                       `json:"releaseTime" yaml:"releaseTime,omitempty"`
	Rewards                   []interface{}             `json:"rewards" yaml:"rewards,omitempty"`
	SuppressOtherRewards      bool                      `json:"suppressOtherRewards" yaml:"suppressOtherRewards,omitempty"`
	Tier                      int                       `json:"tier" yaml:"tier,omitempty"`
}
type Matchmaking struct {
	IsMatchmade          bool `json:"isMatchmade" yaml:"isMatchmade,omitempty"`
	MaxParty             int  `json:"maxParty" yaml:"maxParty,omitempty"`
	MaxPlayers           int  `json:"maxPlayers" yaml:"maxPlayers,omitempty"`
	MinParty             int  `json:"minParty" yaml:"minParty,omitempty"`
	RequiresGuardianOath bool `json:"requiresGuardianOath" yaml:"requiresGuardianOath,omitempty"`
}
type OriginalDisplayProperties struct {
	Description string `json:"description" yaml:"description,omitempty"`
	HasIcon     bool   `json:"hasIcon" yaml:"hasIcon,omitempty"`
	Icon        string `json:"icon" yaml:"icon,omitempty"`
	Name        string `json:"name" yaml:"name,omitempty"`
}

func (d ActivityDefinition) Name() string {
	return ActivityDefinitionName
}

func (d ActivityDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
