// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SeasonDefinitionName = "season"

type SeasonDefinitions map[string]SeasonDefinition

func (d SeasonDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SeasonDefinitions) Values() (out []SeasonDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SeasonDefinitions) Find(id int) (out SeasonDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SeasonDefinition{}
}

func (d SeasonDefinitions) Name() string {
	return SeasonDefinitionName
}

type SeasonDefinition struct {
	Blacklisted               bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties         DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                      int64             `json:"hash" yaml:"hash,omitempty"`
	Index                     int               `json:"index" yaml:"index,omitempty"`
	Redacted                  bool              `json:"redacted" yaml:"redacted,omitempty"`
	SeasonNumber              int               `json:"seasonNumber" yaml:"seasonNumber,omitempty"`
	SeasonPassProgressionHash int               `json:"seasonPassProgressionHash" yaml:"seasonPassProgressionHash,omitempty"`
	SeasonPassUnlockHash      int               `json:"seasonPassUnlockHash" yaml:"seasonPassUnlockHash,omitempty"`
	StartTimeInSeconds        string            `json:"startTimeInSeconds" yaml:"startTimeInSeconds,omitempty"`
}

func (d SeasonDefinition) Name() string {
	return SeasonDefinitionName
}

func (d SeasonDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SeasonDefinition) PrettyJson() ([]byte, error) {
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

func (d SeasonDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
