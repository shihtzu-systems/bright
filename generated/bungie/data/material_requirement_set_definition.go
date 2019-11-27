// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const MaterialRequirementSetDefinitionName = "material-requirement-set"

type MaterialRequirementSetDefinitions map[string]MaterialRequirementSetDefinition

func (d MaterialRequirementSetDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d MaterialRequirementSetDefinitions) Values() (out []MaterialRequirementSetDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d MaterialRequirementSetDefinitions) Find(id int) (out MaterialRequirementSetDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return MaterialRequirementSetDefinition{}
}

func (d MaterialRequirementSetDefinitions) Name() string {
	return MaterialRequirementSetDefinitionName
}

type MaterialRequirementSetDefinition struct {
	Blacklisted bool        `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int         `json:"hash" yaml:"hash,omitempty"`
	Index       int         `json:"index" yaml:"index,omitempty"`
	Materials   []Materials `json:"materials" yaml:"materials,omitempty"`
	Redacted    bool        `json:"redacted" yaml:"redacted,omitempty"`
}
type Materials struct {
	Count                int   `json:"count" yaml:"count,omitempty"`
	DeleteOnAction       bool  `json:"deleteOnAction" yaml:"deleteOnAction,omitempty"`
	ItemHash             int64 `json:"itemHash" yaml:"itemHash,omitempty"`
	OmitFromRequirements bool  `json:"omitFromRequirements" yaml:"omitFromRequirements,omitempty"`
}

func (d MaterialRequirementSetDefinition) Name() string {
	return MaterialRequirementSetDefinitionName
}

func (d MaterialRequirementSetDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d MaterialRequirementSetDefinition) PrettyJson() ([]byte, error) {
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

func (d MaterialRequirementSetDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
