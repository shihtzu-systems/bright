// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const StatDefinitionName = "stat"

type StatDefinitions map[string]StatDefinition

func (d StatDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d StatDefinitions) Values() (out []StatDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d StatDefinitions) Find(id int) (out StatDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return StatDefinition{}
}

func (d StatDefinitions) Name() string {
	return StatDefinitionName
}

type StatDefinition struct {
	AggregationType   int               `json:"aggregationType" yaml:"aggregationType,omitempty"`
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	HasComputedBlock  bool              `json:"hasComputedBlock" yaml:"hasComputedBlock,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Interpolate       bool              `json:"interpolate" yaml:"interpolate,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	StatCategory      int               `json:"statCategory" yaml:"statCategory,omitempty"`
}

func (d StatDefinition) Name() string {
	return StatDefinitionName
}

func (d StatDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d StatDefinition) PrettyJson() ([]byte, error) {
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

func (d StatDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
