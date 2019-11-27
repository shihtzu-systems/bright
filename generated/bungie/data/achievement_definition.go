// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const AchievementDefinitionName = "achievement"

type AchievementDefinitions map[string]AchievementDefinition

func (d AchievementDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d AchievementDefinitions) Values() (out []AchievementDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d AchievementDefinitions) Find(id int) (out AchievementDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return AchievementDefinition{}
}

func (d AchievementDefinitions) Name() string {
	return AchievementDefinitionName
}

type AchievementDefinition struct {
	AcccumulatorThreshold int               `json:"acccumulatorThreshold" yaml:"acccumulatorThreshold,omitempty"`
	Blacklisted           bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties     DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                  int               `json:"hash" yaml:"hash,omitempty"`
	Index                 int               `json:"index" yaml:"index,omitempty"`
	PlatformIndex         int               `json:"platformIndex" yaml:"platformIndex,omitempty"`
	Redacted              bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d AchievementDefinition) Name() string {
	return AchievementDefinitionName
}

func (d AchievementDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d AchievementDefinition) PrettyJson() ([]byte, error) {
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

func (d AchievementDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
