// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RaceDefinitionName = "race"

type RaceDefinitions map[string]RaceDefinition

func (d RaceDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RaceDefinitions) Values() (out []RaceDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RaceDefinitions) Find(id int) (out RaceDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RaceDefinition{}
}

func (d RaceDefinitions) Name() string {
	return RaceDefinitionName
}

type RaceDefinition struct {
	Blacklisted                   bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties             DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	GenderedRaceNames             GenderedRaceNames `json:"genderedRaceNames" yaml:"genderedRaceNames,omitempty"`
	GenderedRaceNamesByGenderHash map[string]string `json:"genderedRaceNamesByGenderHash" yaml:"genderedRaceNamesByGenderHash,omitempty"`
	Hash                          int64             `json:"hash" yaml:"hash,omitempty"`
	Index                         int               `json:"index" yaml:"index,omitempty"`
	RaceType                      int               `json:"raceType" yaml:"raceType,omitempty"`
	Redacted                      bool              `json:"redacted" yaml:"redacted,omitempty"`
}
type GenderedRaceNames struct {
	Female string `json:"Female" yaml:"Female,omitempty"`
	Male   string `json:"Male" yaml:"Male,omitempty"`
}

func (d RaceDefinition) Name() string {
	return RaceDefinitionName
}

func (d RaceDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RaceDefinition) PrettyJson() ([]byte, error) {
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

func (d RaceDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
