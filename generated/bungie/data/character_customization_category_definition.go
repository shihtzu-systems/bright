// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const CharacterCustomizationCategoryDefinitionName = "character-customization-category"

type CharacterCustomizationCategoryDefinitions map[string]CharacterCustomizationCategoryDefinition

func (d CharacterCustomizationCategoryDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d CharacterCustomizationCategoryDefinitions) Values() (out []CharacterCustomizationCategoryDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d CharacterCustomizationCategoryDefinitions) Find(id int) (out CharacterCustomizationCategoryDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return CharacterCustomizationCategoryDefinition{}
}

func (d CharacterCustomizationCategoryDefinitions) Name() string {
	return CharacterCustomizationCategoryDefinitionName
}

type CharacterCustomizationCategoryDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d CharacterCustomizationCategoryDefinition) Name() string {
	return CharacterCustomizationCategoryDefinitionName
}

func (d CharacterCustomizationCategoryDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d CharacterCustomizationCategoryDefinition) PrettyJson() ([]byte, error) {
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

func (d CharacterCustomizationCategoryDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
