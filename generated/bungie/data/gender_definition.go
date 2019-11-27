// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const GenderDefinitionName = "gender"

type GenderDefinitions map[string]GenderDefinition

func (d GenderDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d GenderDefinitions) Values() (out []GenderDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d GenderDefinitions) Find(id int) (out GenderDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return GenderDefinition{}
}

func (d GenderDefinitions) Name() string {
	return GenderDefinitionName
}

type GenderDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	GenderType        int               `json:"genderType" yaml:"genderType,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d GenderDefinition) Name() string {
	return GenderDefinitionName
}

func (d GenderDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d GenderDefinition) PrettyJson() ([]byte, error) {
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

func (d GenderDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
