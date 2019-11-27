// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ItemTierTypeDefinitionName = "item-tier-type"

type ItemTierTypeDefinitions map[string]ItemTierTypeDefinition

func (d ItemTierTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ItemTierTypeDefinitions) Values() (out []ItemTierTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ItemTierTypeDefinitions) Find(id int) (out ItemTierTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ItemTierTypeDefinition{}
}

func (d ItemTierTypeDefinitions) Name() string {
	return ItemTierTypeDefinitionName
}

type ItemTierTypeDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	InfusionProcess   InfusionProcess   `json:"infusionProcess" yaml:"infusionProcess,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}
type InfusionProcess struct {
	BaseQualityTransferRatio float64 `json:"baseQualityTransferRatio" yaml:"baseQualityTransferRatio,omitempty"`
	MinimumQualityIncrement  int     `json:"minimumQualityIncrement" yaml:"minimumQualityIncrement,omitempty"`
}

func (d ItemTierTypeDefinition) Name() string {
	return ItemTierTypeDefinitionName
}

func (d ItemTierTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ItemTierTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d ItemTierTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
