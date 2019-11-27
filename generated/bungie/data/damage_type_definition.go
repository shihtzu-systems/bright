// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const DamageTypeDefinitionName = "damage-type"

type DamageTypeDefinitions map[string]DamageTypeDefinition

func (d DamageTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d DamageTypeDefinitions) Values() (out []DamageTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d DamageTypeDefinitions) Find(id int) (out DamageTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return DamageTypeDefinition{}
}

func (d DamageTypeDefinitions) Name() string {
	return DamageTypeDefinitionName
}

type DamageTypeDefinition struct {
	Blacklisted         bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties   DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	EnumValue           int               `json:"enumValue" yaml:"enumValue,omitempty"`
	Hash                int               `json:"hash" yaml:"hash,omitempty"`
	Index               int               `json:"index" yaml:"index,omitempty"`
	Redacted            bool              `json:"redacted" yaml:"redacted,omitempty"`
	ShowIcon            bool              `json:"showIcon" yaml:"showIcon,omitempty"`
	TransparentIconPath string            `json:"transparentIconPath" yaml:"transparentIconPath,omitempty"`
}

func (d DamageTypeDefinition) Name() string {
	return DamageTypeDefinitionName
}

func (d DamageTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d DamageTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d DamageTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
