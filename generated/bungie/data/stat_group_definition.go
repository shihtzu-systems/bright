// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const StatGroupDefinitionName = "stat-group"

type StatGroupDefinitions map[string]StatGroupDefinition

func (d StatGroupDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d StatGroupDefinitions) Values() (out []StatGroupDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d StatGroupDefinitions) Find(id int) (out StatGroupDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return StatGroupDefinition{}
}

func (d StatGroupDefinitions) Name() string {
	return StatGroupDefinitionName
}

type StatGroupDefinition struct {
	Blacklisted  bool          `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash         int           `json:"hash" yaml:"hash,omitempty"`
	Index        int           `json:"index" yaml:"index,omitempty"`
	MaximumValue int           `json:"maximumValue" yaml:"maximumValue,omitempty"`
	Overrides    interface{}   `json:"overrides" yaml:"overrides,omitempty"`
	Redacted     bool          `json:"redacted" yaml:"redacted,omitempty"`
	ScaledStats  []ScaledStats `json:"scaledStats" yaml:"scaledStats,omitempty"`
	UiPosition   int           `json:"uiPosition" yaml:"uiPosition,omitempty"`
}
type DisplayInterpolation struct {
	Value  int `json:"value" yaml:"value,omitempty"`
	Weight int `json:"weight" yaml:"weight,omitempty"`
}
type ScaledStats struct {
	DisplayAsNumeric     bool                   `json:"displayAsNumeric" yaml:"displayAsNumeric,omitempty"`
	DisplayInterpolation []DisplayInterpolation `json:"displayInterpolation" yaml:"displayInterpolation,omitempty"`
	MaximumValue         int                    `json:"maximumValue" yaml:"maximumValue,omitempty"`
	StatHash             int64                  `json:"statHash" yaml:"statHash,omitempty"`
}

func (d StatGroupDefinition) Name() string {
	return StatGroupDefinitionName
}

func (d StatGroupDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d StatGroupDefinition) PrettyJson() ([]byte, error) {
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

func (d StatGroupDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
