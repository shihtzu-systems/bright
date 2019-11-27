// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardMappingDefinitionName = "reward-mapping"

type RewardMappingDefinitions map[string]RewardMappingDefinition

func (d RewardMappingDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardMappingDefinitions) Values() (out []RewardMappingDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardMappingDefinitions) Find(id int) (out RewardMappingDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardMappingDefinition{}
}

func (d RewardMappingDefinitions) Name() string {
	return RewardMappingDefinitionName
}

type RewardMappingDefinition struct {
	Blacklisted  bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash         int  `json:"hash" yaml:"hash,omitempty"`
	Index        int  `json:"index" yaml:"index,omitempty"`
	IsGlobal     bool `json:"isGlobal" yaml:"isGlobal,omitempty"`
	MappingHash  int  `json:"mappingHash" yaml:"mappingHash,omitempty"`
	MappingIndex int  `json:"mappingIndex" yaml:"mappingIndex,omitempty"`
	Redacted     bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d RewardMappingDefinition) Name() string {
	return RewardMappingDefinitionName
}

func (d RewardMappingDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardMappingDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardMappingDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
