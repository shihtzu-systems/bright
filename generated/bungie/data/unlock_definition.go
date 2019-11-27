// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const UnlockDefinitionName = "unlock"

type UnlockDefinitions map[string]UnlockDefinition

func (d UnlockDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d UnlockDefinitions) Values() (out []UnlockDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d UnlockDefinitions) Find(id int) (out UnlockDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return UnlockDefinition{}
}

func (d UnlockDefinitions) Name() string {
	return UnlockDefinitionName
}

type UnlockDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	Scope             int               `json:"scope" yaml:"scope,omitempty"`
	UnlockType        int               `json:"unlockType" yaml:"unlockType,omitempty"`
}

func (d UnlockDefinition) Name() string {
	return UnlockDefinitionName
}

func (d UnlockDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d UnlockDefinition) PrettyJson() ([]byte, error) {
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

func (d UnlockDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
