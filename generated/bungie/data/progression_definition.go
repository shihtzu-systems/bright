// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ProgressionDefinitionName = "progression"

type ProgressionDefinitions map[string]ProgressionDefinition

func (d ProgressionDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ProgressionDefinitions) Values() (out []ProgressionDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ProgressionDefinitions) Find(id int) (out ProgressionDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ProgressionDefinition{}
}

func (d ProgressionDefinitions) Name() string {
	return ProgressionDefinitionName
}

type ProgressionDefinition struct {
	Blacklisted                      bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CurrentResetCountUnlockValueHash int               `json:"currentResetCountUnlockValueHash" yaml:"currentResetCountUnlockValueHash,omitempty"`
	DisplayProperties                DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                             int64             `json:"hash" yaml:"hash,omitempty"`
	Index                            int               `json:"index" yaml:"index,omitempty"`
	ProgressToNextStepScaling        int               `json:"progressToNextStepScaling" yaml:"progressToNextStepScaling,omitempty"`
	Redacted                         bool              `json:"redacted" yaml:"redacted,omitempty"`
	RepeatLastStep                   bool              `json:"repeatLastStep" yaml:"repeatLastStep,omitempty"`
	RewardItems                      []interface{}     `json:"rewardItems" yaml:"rewardItems,omitempty"`
	Scope                            int               `json:"scope" yaml:"scope,omitempty"`
	Steps                            []Steps           `json:"steps" yaml:"steps,omitempty"`
	StorageMappingIndex              int               `json:"storageMappingIndex" yaml:"storageMappingIndex,omitempty"`
	Visible                          bool              `json:"visible" yaml:"visible,omitempty"`
}
type Steps struct {
	DisplayEffectType int    `json:"displayEffectType" yaml:"displayEffectType,omitempty"`
	ProgressTotal     int    `json:"progressTotal" yaml:"progressTotal,omitempty"`
	StepName          string `json:"stepName" yaml:"stepName,omitempty"`
}

func (d ProgressionDefinition) Name() string {
	return ProgressionDefinitionName
}

func (d ProgressionDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ProgressionDefinition) PrettyJson() ([]byte, error) {
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

func (d ProgressionDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
