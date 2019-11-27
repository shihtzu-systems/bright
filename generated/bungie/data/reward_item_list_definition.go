// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardItemListDefinitionName = "reward-item-list"

type RewardItemListDefinitions map[string]RewardItemListDefinition

func (d RewardItemListDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardItemListDefinitions) Values() (out []RewardItemListDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardItemListDefinitions) Find(id int) (out RewardItemListDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardItemListDefinition{}
}

func (d RewardItemListDefinitions) Name() string {
	return RewardItemListDefinitionName
}

type RewardItemListDefinition struct {
	Blacklisted bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int  `json:"hash" yaml:"hash,omitempty"`
	Index       int  `json:"index" yaml:"index,omitempty"`
	Redacted    bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d RewardItemListDefinition) Name() string {
	return RewardItemListDefinitionName
}

func (d RewardItemListDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardItemListDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardItemListDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
