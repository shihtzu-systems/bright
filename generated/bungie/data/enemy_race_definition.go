// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const EnemyRaceDefinitionName = "enemy-race"

type EnemyRaceDefinitions map[string]EnemyRaceDefinition

func (d EnemyRaceDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d EnemyRaceDefinitions) Values() (out []EnemyRaceDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d EnemyRaceDefinitions) Find(id int) (out EnemyRaceDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return EnemyRaceDefinition{}
}

func (d EnemyRaceDefinitions) Name() string {
	return EnemyRaceDefinitionName
}

type EnemyRaceDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d EnemyRaceDefinition) Name() string {
	return EnemyRaceDefinitionName
}

func (d EnemyRaceDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d EnemyRaceDefinition) PrettyJson() ([]byte, error) {
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

func (d EnemyRaceDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
