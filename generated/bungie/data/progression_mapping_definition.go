// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ProgressionMappingDefinitionName = "progression-mapping"

type ProgressionMappingDefinitions map[string]ProgressionMappingDefinition

func (d ProgressionMappingDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ProgressionMappingDefinitions) Values() (out []ProgressionMappingDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ProgressionMappingDefinitions) Find(id int) (out ProgressionMappingDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ProgressionMappingDefinition{}
}

func (d ProgressionMappingDefinitions) Name() string {
	return ProgressionMappingDefinitionName
}

type ProgressionMappingDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d ProgressionMappingDefinition) Name() string {
	return ProgressionMappingDefinitionName
}

func (d ProgressionMappingDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ProgressionMappingDefinition) PrettyJson() ([]byte, error) {
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

func (d ProgressionMappingDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
