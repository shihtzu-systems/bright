// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const PlugSetDefinitionName = "plug-set"

type PlugSetDefinitions map[string]PlugSetDefinition

func (d PlugSetDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d PlugSetDefinitions) Values() (out []PlugSetDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d PlugSetDefinitions) Find(id int) (out PlugSetDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return PlugSetDefinition{}
}

func (d PlugSetDefinitions) Name() string {
	return PlugSetDefinitionName
}

type PlugSetDefinition struct {
	Blacklisted       bool                `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties   `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int                 `json:"hash" yaml:"hash,omitempty"`
	Index             int                 `json:"index" yaml:"index,omitempty"`
	Redacted          bool                `json:"redacted" yaml:"redacted,omitempty"`
	ReusablePlugItems []ReusablePlugItems `json:"reusablePlugItems" yaml:"reusablePlugItems,omitempty"`
}
type ReusablePlugItems struct {
	AlternateWeight float64 `json:"alternateWeight" yaml:"alternateWeight,omitempty"`
	PlugItemHash    float64 `json:"plugItemHash" yaml:"plugItemHash,omitempty"`
	Weight          float64 `json:"weight" yaml:"weight,omitempty"`
}

func (d PlugSetDefinition) Name() string {
	return PlugSetDefinitionName
}

func (d PlugSetDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d PlugSetDefinition) PrettyJson() ([]byte, error) {
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

func (d PlugSetDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
