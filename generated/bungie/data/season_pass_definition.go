// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SeasonPassDefinitionName = "season-pass"

type SeasonPassDefinitions map[string]SeasonPassDefinition

func (d SeasonPassDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SeasonPassDefinitions) Values() (out []SeasonPassDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SeasonPassDefinitions) Find(id int) (out SeasonPassDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SeasonPassDefinition{}
}

func (d SeasonPassDefinitions) Name() string {
	return SeasonPassDefinitionName
}

type SeasonPassDefinition struct {
	Blacklisted             bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties       DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                    int               `json:"hash" yaml:"hash,omitempty"`
	Index                   int               `json:"index" yaml:"index,omitempty"`
	PrestigeProgressionHash int64             `json:"prestigeProgressionHash" yaml:"prestigeProgressionHash,omitempty"`
	Redacted                bool              `json:"redacted" yaml:"redacted,omitempty"`
	RewardProgressionHash   int               `json:"rewardProgressionHash" yaml:"rewardProgressionHash,omitempty"`
}

func (d SeasonPassDefinition) Name() string {
	return SeasonPassDefinitionName
}

func (d SeasonPassDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SeasonPassDefinition) PrettyJson() ([]byte, error) {
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

func (d SeasonPassDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
