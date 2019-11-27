// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityModeDefinitionName = "activity-mode"

type ActivityModeDefinitions map[string]ActivityModeDefinition

func (d ActivityModeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityModeDefinitions) Values() (out []ActivityModeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityModeDefinitions) Find(id int) (out ActivityModeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityModeDefinition{}
}

func (d ActivityModeDefinitions) Name() string {
	return ActivityModeDefinitionName
}

type ActivityModeDefinition struct {
	ActivityModeCategory  int               `json:"activityModeCategory" yaml:"activityModeCategory,omitempty"`
	Blacklisted           bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Display               bool              `json:"display" yaml:"display,omitempty"`
	DisplayProperties     DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	FriendlyName          string            `json:"friendlyName" yaml:"friendlyName,omitempty"`
	Hash                  int64             `json:"hash" yaml:"hash,omitempty"`
	Index                 int               `json:"index" yaml:"index,omitempty"`
	IsAggregateMode       bool              `json:"isAggregateMode" yaml:"isAggregateMode,omitempty"`
	IsTeamBased           bool              `json:"isTeamBased" yaml:"isTeamBased,omitempty"`
	ModeType              int               `json:"modeType" yaml:"modeType,omitempty"`
	Order                 int               `json:"order" yaml:"order,omitempty"`
	ParentHashes          []int             `json:"parentHashes" yaml:"parentHashes,omitempty"`
	PgcrImage             string            `json:"pgcrImage" yaml:"pgcrImage,omitempty"`
	Redacted              bool              `json:"redacted" yaml:"redacted,omitempty"`
	SupportsFeedFiltering bool              `json:"supportsFeedFiltering" yaml:"supportsFeedFiltering,omitempty"`
	Tier                  int               `json:"tier" yaml:"tier,omitempty"`
}

func (d ActivityModeDefinition) Name() string {
	return ActivityModeDefinitionName
}

func (d ActivityModeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityModeDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityModeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
