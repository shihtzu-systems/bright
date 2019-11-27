// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const BondDefinitionName = "bond"

type BondDefinitions map[string]BondDefinition

func (d BondDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d BondDefinitions) Values() (out []BondDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d BondDefinitions) Find(id int) (out BondDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return BondDefinition{}
}

func (d BondDefinitions) Name() string {
	return BondDefinitionName
}

type BondDefinition struct {
	Blacklisted             bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties       DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                    int               `json:"hash" yaml:"hash,omitempty"`
	Index                   int               `json:"index" yaml:"index,omitempty"`
	ProvidedUnlockHash      int               `json:"providedUnlockHash" yaml:"providedUnlockHash,omitempty"`
	ProvidedUnlockValueHash int               `json:"providedUnlockValueHash" yaml:"providedUnlockValueHash,omitempty"`
	Redacted                bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d BondDefinition) Name() string {
	return BondDefinitionName
}

func (d BondDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d BondDefinition) PrettyJson() ([]byte, error) {
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

func (d BondDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
