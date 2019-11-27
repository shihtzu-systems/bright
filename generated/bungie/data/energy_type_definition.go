// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const EnergyTypeDefinitionName = "energy-type"

type EnergyTypeDefinitions map[string]EnergyTypeDefinition

func (d EnergyTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d EnergyTypeDefinitions) Values() (out []EnergyTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d EnergyTypeDefinitions) Find(id int) (out EnergyTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return EnergyTypeDefinition{}
}

func (d EnergyTypeDefinitions) Name() string {
	return EnergyTypeDefinitionName
}

type EnergyTypeDefinition struct {
	Blacklisted         bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CapacityStatHash    int               `json:"capacityStatHash" yaml:"capacityStatHash,omitempty"`
	CostStatHash        int64             `json:"costStatHash" yaml:"costStatHash,omitempty"`
	DisplayProperties   DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	EnumValue           int               `json:"enumValue" yaml:"enumValue,omitempty"`
	Hash                int               `json:"hash" yaml:"hash,omitempty"`
	Index               int               `json:"index" yaml:"index,omitempty"`
	Redacted            bool              `json:"redacted" yaml:"redacted,omitempty"`
	ShowIcon            bool              `json:"showIcon" yaml:"showIcon,omitempty"`
	TransparentIconPath string            `json:"transparentIconPath" yaml:"transparentIconPath,omitempty"`
}

func (d EnergyTypeDefinition) Name() string {
	return EnergyTypeDefinitionName
}

func (d EnergyTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d EnergyTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d EnergyTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
