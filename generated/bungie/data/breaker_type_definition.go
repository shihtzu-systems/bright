// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const BreakerTypeDefinitionName = "breaker-type"

type BreakerTypeDefinitions map[string]BreakerTypeDefinition

func (d BreakerTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d BreakerTypeDefinitions) Values() (out []BreakerTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d BreakerTypeDefinitions) Find(id int) (out BreakerTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return BreakerTypeDefinition{}
}

func (d BreakerTypeDefinitions) Name() string {
	return BreakerTypeDefinitionName
}

type BreakerTypeDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	EnumValue         int               `json:"enumValue" yaml:"enumValue,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	UnlockHash        int               `json:"unlockHash" yaml:"unlockHash,omitempty"`
}

func (d BreakerTypeDefinition) Name() string {
	return BreakerTypeDefinitionName
}

func (d BreakerTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d BreakerTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d BreakerTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
