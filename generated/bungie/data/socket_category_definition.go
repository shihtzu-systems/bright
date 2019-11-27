// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SocketCategoryDefinitionName = "socket-category"

type SocketCategoryDefinitions map[string]SocketCategoryDefinition

func (d SocketCategoryDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SocketCategoryDefinitions) Values() (out []SocketCategoryDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SocketCategoryDefinitions) Find(id int) (out SocketCategoryDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SocketCategoryDefinition{}
}

func (d SocketCategoryDefinitions) Name() string {
	return SocketCategoryDefinitionName
}

type SocketCategoryDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CategoryStyle     int               `json:"categoryStyle" yaml:"categoryStyle,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
	UiCategoryStyle   int               `json:"uiCategoryStyle" yaml:"uiCategoryStyle,omitempty"`
}

func (d SocketCategoryDefinition) Name() string {
	return SocketCategoryDefinitionName
}

func (d SocketCategoryDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SocketCategoryDefinition) PrettyJson() ([]byte, error) {
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

func (d SocketCategoryDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
