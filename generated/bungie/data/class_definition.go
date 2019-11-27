// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ClassDefinitionName = "class"

type ClassDefinitions map[string]ClassDefinition

func (d ClassDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ClassDefinitions) Values() (out []ClassDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ClassDefinitions) Find(id int) (out ClassDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ClassDefinition{}
}

func (d ClassDefinitions) Name() string {
	return ClassDefinitionName
}

type ClassDefinition struct {
	Blacklisted                    bool                   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	ClassType                      int                    `json:"classType" yaml:"classType,omitempty"`
	DisplayProperties              DisplayProperties      `json:"displayProperties" yaml:"displayProperties,omitempty"`
	GenderedClassNames             GenderedClassNames     `json:"genderedClassNames" yaml:"genderedClassNames,omitempty"`
	GenderedClassNamesByGenderHash map[string]interface{} `json:"genderedClassNamesByGenderHash" yaml:"genderedClassNamesByGenderHash,omitempty"`
	Hash                           int                    `json:"hash" yaml:"hash,omitempty"`
	Index                          int                    `json:"index" yaml:"index,omitempty"`
	Redacted                       bool                   `json:"redacted" yaml:"redacted,omitempty"`
}
type GenderedClassNames struct {
	Female string `json:"female" yaml:"female,omitempty"`
	Male   string `json:"male" yaml:"male,omitempty"`
}

func (d ClassDefinition) Name() string {
	return ClassDefinitionName
}

func (d ClassDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ClassDefinition) PrettyJson() ([]byte, error) {
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

func (d ClassDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
