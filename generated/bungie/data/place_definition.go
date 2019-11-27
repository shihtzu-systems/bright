// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const PlaceDefinitionName = "place"

type PlaceDefinitions map[string]PlaceDefinition

func (d PlaceDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d PlaceDefinitions) Values() (out []PlaceDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d PlaceDefinitions) Find(id int) (out PlaceDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return PlaceDefinition{}
}

func (d PlaceDefinitions) Name() string {
	return PlaceDefinitionName
}

type PlaceDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d PlaceDefinition) Name() string {
	return PlaceDefinitionName
}

func (d PlaceDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d PlaceDefinition) PrettyJson() ([]byte, error) {
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

func (d PlaceDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
