// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SackRewardItemListDefinitionName = "sack-reward-item-list"

type SackRewardItemListDefinitions map[string]SackRewardItemListDefinition

func (d SackRewardItemListDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SackRewardItemListDefinitions) Values() (out []SackRewardItemListDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SackRewardItemListDefinitions) Find(id int) (out SackRewardItemListDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SackRewardItemListDefinition{}
}

func (d SackRewardItemListDefinitions) Name() string {
	return SackRewardItemListDefinitionName
}

type SackRewardItemListDefinition struct {
	Blacklisted bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int  `json:"hash" yaml:"hash,omitempty"`
	Index       int  `json:"index" yaml:"index,omitempty"`
	Redacted    bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d SackRewardItemListDefinition) Name() string {
	return SackRewardItemListDefinitionName
}

func (d SackRewardItemListDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SackRewardItemListDefinition) PrettyJson() ([]byte, error) {
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

func (d SackRewardItemListDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
