// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const LoreDefinitionName = "lore"

type LoreDefinitions map[string]LoreDefinition

func (d LoreDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d LoreDefinitions) Values() (out []LoreDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d LoreDefinitions) Find(id int) (out LoreDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return LoreDefinition{}
}

func (d LoreDefinitions) Name() string {
	return LoreDefinitionName
}

type LoreDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	Subtitle          string            `json:"subtitle" yaml:"subtitle,omitempty"`
}

func (d LoreDefinition) Name() string {
	return LoreDefinitionName
}

func (d LoreDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d LoreDefinition) PrettyJson() ([]byte, error) {
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

func (d LoreDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
