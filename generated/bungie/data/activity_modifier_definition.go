// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityModifierDefinitionName = "activity-modifier"

type ActivityModifierDefinitions map[string]ActivityModifierDefinition

func (d ActivityModifierDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityModifierDefinitions) Values() (out []ActivityModifierDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityModifierDefinitions) Find(id int) (out ActivityModifierDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityModifierDefinition{}
}

func (d ActivityModifierDefinitions) Name() string {
	return ActivityModifierDefinitionName
}

type ActivityModifierDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d ActivityModifierDefinition) Name() string {
	return ActivityModifierDefinitionName
}

func (d ActivityModifierDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityModifierDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityModifierDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
