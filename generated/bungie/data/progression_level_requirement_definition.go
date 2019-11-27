// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ProgressionLevelRequirementDefinitionName = "progression-level-requirement"

type ProgressionLevelRequirementDefinitions map[string]ProgressionLevelRequirementDefinition

func (d ProgressionLevelRequirementDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ProgressionLevelRequirementDefinitions) Values() (out []ProgressionLevelRequirementDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ProgressionLevelRequirementDefinitions) Find(id int) (out ProgressionLevelRequirementDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ProgressionLevelRequirementDefinition{}
}

func (d ProgressionLevelRequirementDefinitions) Name() string {
	return ProgressionLevelRequirementDefinitionName
}

type ProgressionLevelRequirementDefinition struct {
	Blacklisted      bool               `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash             int                `json:"hash" yaml:"hash,omitempty"`
	Index            int                `json:"index" yaml:"index,omitempty"`
	ProgressionHash  int                `json:"progressionHash" yaml:"progressionHash,omitempty"`
	Redacted         bool               `json:"redacted" yaml:"redacted,omitempty"`
	RequirementCurve []RequirementCurve `json:"requirementCurve" yaml:"requirementCurve,omitempty"`
}
type RequirementCurve struct {
	Value  float64 `json:"value" yaml:"value,omitempty"`
	Weight float64 `json:"weight" yaml:"weight,omitempty"`
}

func (d ProgressionLevelRequirementDefinition) Name() string {
	return ProgressionLevelRequirementDefinitionName
}

func (d ProgressionLevelRequirementDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ProgressionLevelRequirementDefinition) PrettyJson() ([]byte, error) {
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

func (d ProgressionLevelRequirementDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
