// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityTypeDefinitionName = "activity-type"

type ActivityTypeDefinitions map[string]ActivityTypeDefinition

func (d ActivityTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityTypeDefinitions) Values() (out []ActivityTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityTypeDefinitions) Find(id int) (out ActivityTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityTypeDefinition{}
}

func (d ActivityTypeDefinitions) Name() string {
	return ActivityTypeDefinitionName
}

type ActivityTypeDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d ActivityTypeDefinition) Name() string {
	return ActivityTypeDefinitionName
}

func (d ActivityTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
