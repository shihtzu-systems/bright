// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ItemCategoryDefinitionName = "item-category"

type ItemCategoryDefinitions map[string]ItemCategoryDefinition

func (d ItemCategoryDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ItemCategoryDefinitions) Values() (out []ItemCategoryDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ItemCategoryDefinitions) Find(id int) (out ItemCategoryDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ItemCategoryDefinition{}
}

func (d ItemCategoryDefinitions) Name() string {
	return ItemCategoryDefinitionName
}

type ItemCategoryDefinition struct {
	Blacklisted             bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Deprecated              bool              `json:"deprecated" yaml:"deprecated,omitempty"`
	DisplayProperties       DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	GrantDestinyBreakerType int               `json:"grantDestinyBreakerType" yaml:"grantDestinyBreakerType,omitempty"`
	GrantDestinyClass       int               `json:"grantDestinyClass" yaml:"grantDestinyClass,omitempty"`
	GrantDestinyItemType    int               `json:"grantDestinyItemType" yaml:"grantDestinyItemType,omitempty"`
	GrantDestinySubType     int               `json:"grantDestinySubType" yaml:"grantDestinySubType,omitempty"`
	GroupCategoryOnly       bool              `json:"groupCategoryOnly" yaml:"groupCategoryOnly,omitempty"`
	GroupedCategoryHashes   []interface{}     `json:"groupedCategoryHashes" yaml:"groupedCategoryHashes,omitempty"`
	Hash                    int               `json:"hash" yaml:"hash,omitempty"`
	Index                   int               `json:"index" yaml:"index,omitempty"`
	IsPlug                  bool              `json:"isPlug" yaml:"isPlug,omitempty"`
	ParentCategoryHashes    []int             `json:"parentCategoryHashes" yaml:"parentCategoryHashes,omitempty"`
	Redacted                bool              `json:"redacted" yaml:"redacted,omitempty"`
	ShortTitle              string            `json:"shortTitle" yaml:"shortTitle,omitempty"`
	Visible                 bool              `json:"visible" yaml:"visible,omitempty"`
}

func (d ItemCategoryDefinition) Name() string {
	return ItemCategoryDefinitionName
}

func (d ItemCategoryDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ItemCategoryDefinition) PrettyJson() ([]byte, error) {
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

func (d ItemCategoryDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
