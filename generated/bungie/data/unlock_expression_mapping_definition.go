// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const UnlockExpressionMappingDefinitionName = "unlock-expression-mapping"

type UnlockExpressionMappingDefinitions map[string]UnlockExpressionMappingDefinition

func (d UnlockExpressionMappingDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d UnlockExpressionMappingDefinitions) Values() (out []UnlockExpressionMappingDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d UnlockExpressionMappingDefinitions) Find(id int) (out UnlockExpressionMappingDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return UnlockExpressionMappingDefinition{}
}

func (d UnlockExpressionMappingDefinitions) Name() string {
	return UnlockExpressionMappingDefinitionName
}

type UnlockExpressionMappingDefinition struct {
	Blacklisted bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int  `json:"hash" yaml:"hash,omitempty"`
	Index       int  `json:"index" yaml:"index,omitempty"`
	Redacted    bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d UnlockExpressionMappingDefinition) Name() string {
	return UnlockExpressionMappingDefinitionName
}

func (d UnlockExpressionMappingDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d UnlockExpressionMappingDefinition) PrettyJson() ([]byte, error) {
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

func (d UnlockExpressionMappingDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
