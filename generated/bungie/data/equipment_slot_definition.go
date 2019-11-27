// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const EquipmentSlotDefinitionName = "equipment-slot"

type EquipmentSlotDefinitions map[string]EquipmentSlotDefinition

func (d EquipmentSlotDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d EquipmentSlotDefinitions) Values() (out []EquipmentSlotDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d EquipmentSlotDefinitions) Find(id int) (out EquipmentSlotDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return EquipmentSlotDefinition{}
}

func (d EquipmentSlotDefinitions) Name() string {
	return EquipmentSlotDefinitionName
}

type EquipmentSlotDefinition struct {
	ApplyCustomArtDyes    bool              `json:"applyCustomArtDyes" yaml:"applyCustomArtDyes,omitempty"`
	ArtDyeChannels        []ArtDyeChannels  `json:"artDyeChannels" yaml:"artDyeChannels,omitempty"`
	Blacklisted           bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	BucketTypeHash        int64             `json:"bucketTypeHash" yaml:"bucketTypeHash,omitempty"`
	DisplayProperties     DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	EquipmentCategoryHash int               `json:"equipmentCategoryHash" yaml:"equipmentCategoryHash,omitempty"`
	Hash                  int64             `json:"hash" yaml:"hash,omitempty"`
	Index                 int               `json:"index" yaml:"index,omitempty"`
	Redacted              bool              `json:"redacted" yaml:"redacted,omitempty"`
}
type ArtDyeChannels struct {
	ArtDyeChannelHash int `json:"artDyeChannelHash" yaml:"artDyeChannelHash,omitempty"`
}

func (d EquipmentSlotDefinition) Name() string {
	return EquipmentSlotDefinitionName
}

func (d EquipmentSlotDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d EquipmentSlotDefinition) PrettyJson() ([]byte, error) {
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

func (d EquipmentSlotDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
