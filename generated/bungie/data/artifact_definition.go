// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ArtifactDefinitionName = "artifact"

type ArtifactDefinitions map[string]ArtifactDefinition

func (d ArtifactDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ArtifactDefinitions) Values() (out []ArtifactDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ArtifactDefinitions) Find(id int) (out ArtifactDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ArtifactDefinition{}
}

func (d ArtifactDefinitions) Name() string {
	return ArtifactDefinitionName
}

type ArtifactDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	Tiers             []Tiers           `json:"tiers" yaml:"tiers,omitempty"`
	TranslationBlock  TranslationBlock  `json:"translationBlock" yaml:"translationBlock,omitempty"`
}
type ArtifactItems struct {
	ActiveUnlockHash int   `json:"activeUnlockHash" yaml:"activeUnlockHash,omitempty"`
	ItemHash         int64 `json:"itemHash" yaml:"itemHash,omitempty"`
}
type Tiers struct {
	DisplayTitle                       string          `json:"displayTitle" yaml:"displayTitle,omitempty"`
	Items                              []ArtifactItems `json:"items" yaml:"items,omitempty"`
	MinimumUnlockPointsUsedRequirement int             `json:"minimumUnlockPointsUsedRequirement" yaml:"minimumUnlockPointsUsedRequirement,omitempty"`
	ProgressRequirementMessage         string          `json:"progressRequirementMessage" yaml:"progressRequirementMessage,omitempty"`
	TierHash                           int64           `json:"tierHash" yaml:"tierHash,omitempty"`
}

func (d ArtifactDefinition) Name() string {
	return ArtifactDefinitionName
}

func (d ArtifactDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ArtifactDefinition) PrettyJson() ([]byte, error) {
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

func (d ArtifactDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
