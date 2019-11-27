// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardAdjusterPointerDefinitionName = "reward-adjuster-pointer"

type RewardAdjusterPointerDefinitions map[string]RewardAdjusterPointerDefinition

func (d RewardAdjusterPointerDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardAdjusterPointerDefinitions) Values() (out []RewardAdjusterPointerDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardAdjusterPointerDefinitions) Find(id int) (out RewardAdjusterPointerDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardAdjusterPointerDefinition{}
}

func (d RewardAdjusterPointerDefinitions) Name() string {
	return RewardAdjusterPointerDefinitionName
}

type RewardAdjusterPointerDefinition struct {
	AdjusterType int  `json:"adjusterType" yaml:"adjusterType,omitempty"`
	Blacklisted  bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash         int  `json:"hash" yaml:"hash,omitempty"`
	Index        int  `json:"index" yaml:"index,omitempty"`
	Redacted     bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d RewardAdjusterPointerDefinition) Name() string {
	return RewardAdjusterPointerDefinitionName
}

func (d RewardAdjusterPointerDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardAdjusterPointerDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardAdjusterPointerDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
