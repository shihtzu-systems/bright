// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const MedalTierDefinitionName = "medal-tier"

type MedalTierDefinitions map[string]MedalTierDefinition

func (d MedalTierDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d MedalTierDefinitions) Values() (out []MedalTierDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d MedalTierDefinitions) Find(id int) (out MedalTierDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return MedalTierDefinition{}
}

func (d MedalTierDefinitions) Name() string {
	return MedalTierDefinitionName
}

type MedalTierDefinition struct {
	Blacklisted bool   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int    `json:"hash" yaml:"hash,omitempty"`
	Index       int    `json:"index" yaml:"index,omitempty"`
	Order       int    `json:"order" yaml:"order,omitempty"`
	Redacted    bool   `json:"redacted" yaml:"redacted,omitempty"`
	TierName    string `json:"tierName" yaml:"tierName,omitempty"`
}

func (d MedalTierDefinition) Name() string {
	return MedalTierDefinitionName
}

func (d MedalTierDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d MedalTierDefinition) PrettyJson() ([]byte, error) {
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

func (d MedalTierDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
