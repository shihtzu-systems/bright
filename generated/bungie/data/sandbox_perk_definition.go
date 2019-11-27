// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SandboxPerkDefinitionName = "sandbox-perk"

type SandboxPerkDefinitions map[string]SandboxPerkDefinition

func (d SandboxPerkDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SandboxPerkDefinitions) Values() (out []SandboxPerkDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SandboxPerkDefinitions) Find(id int) (out SandboxPerkDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SandboxPerkDefinition{}
}

func (d SandboxPerkDefinitions) Name() string {
	return SandboxPerkDefinitionName
}

type SandboxPerkDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DamageType        int               `json:"damageType" yaml:"damageType,omitempty"`
	DamageTypeHash    int64             `json:"damageTypeHash" yaml:"damageTypeHash,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	IsDisplayable     bool              `json:"isDisplayable" yaml:"isDisplayable,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d SandboxPerkDefinition) Name() string {
	return SandboxPerkDefinitionName
}

func (d SandboxPerkDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SandboxPerkDefinition) PrettyJson() ([]byte, error) {
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

func (d SandboxPerkDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
