// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardAdjusterProgressionMapDefinitionName = "reward-adjuster-progression-map"

type RewardAdjusterProgressionMapDefinitions map[string]RewardAdjusterProgressionMapDefinition

func (d RewardAdjusterProgressionMapDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardAdjusterProgressionMapDefinitions) Values() (out []RewardAdjusterProgressionMapDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardAdjusterProgressionMapDefinitions) Find(id int) (out RewardAdjusterProgressionMapDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardAdjusterProgressionMapDefinition{}
}

func (d RewardAdjusterProgressionMapDefinitions) Name() string {
	return RewardAdjusterProgressionMapDefinitionName
}

type RewardAdjusterProgressionMapDefinition struct {
	Blacklisted bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int  `json:"hash" yaml:"hash,omitempty"`
	Index       int  `json:"index" yaml:"index,omitempty"`
	IsAdditive  bool `json:"isAdditive" yaml:"isAdditive,omitempty"`
	Redacted    bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d RewardAdjusterProgressionMapDefinition) Name() string {
	return RewardAdjusterProgressionMapDefinitionName
}

func (d RewardAdjusterProgressionMapDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardAdjusterProgressionMapDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardAdjusterProgressionMapDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
