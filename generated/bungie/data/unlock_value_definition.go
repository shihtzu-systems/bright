// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const UnlockValueDefinitionName = "unlock-value"

type UnlockValueDefinitions map[string]UnlockValueDefinition

func (d UnlockValueDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d UnlockValueDefinitions) Values() (out []UnlockValueDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d UnlockValueDefinitions) Find(id int) (out UnlockValueDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return UnlockValueDefinition{}
}

func (d UnlockValueDefinitions) Name() string {
	return UnlockValueDefinitionName
}

type UnlockValueDefinition struct {
	AggregationType int  `json:"aggregationType" yaml:"aggregationType,omitempty"`
	Blacklisted     bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash            int  `json:"hash" yaml:"hash,omitempty"`
	Index           int  `json:"index" yaml:"index,omitempty"`
	MappingIndex    int  `json:"mappingIndex" yaml:"mappingIndex,omitempty"`
	Redacted        bool `json:"redacted" yaml:"redacted,omitempty"`
	Scope           int  `json:"scope" yaml:"scope,omitempty"`
}

func (d UnlockValueDefinition) Name() string {
	return UnlockValueDefinitionName
}

func (d UnlockValueDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d UnlockValueDefinition) PrettyJson() ([]byte, error) {
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

func (d UnlockValueDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
