// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const FactionDefinitionName = "faction"

type FactionDefinitions map[string]FactionDefinition

func (d FactionDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d FactionDefinitions) Values() (out []FactionDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d FactionDefinitions) Find(id int) (out FactionDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return FactionDefinition{}
}

func (d FactionDefinitions) Name() string {
	return FactionDefinitionName
}

type FactionDefinition struct {
	Blacklisted       bool                   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties      `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int                    `json:"hash" yaml:"hash,omitempty"`
	Index             int                    `json:"index" yaml:"index,omitempty"`
	ProgressionHash   int                    `json:"progressionHash" yaml:"progressionHash,omitempty"`
	Redacted          bool                   `json:"redacted" yaml:"redacted,omitempty"`
	RewardItemHash    int64                  `json:"rewardItemHash" yaml:"rewardItemHash,omitempty"`
	RewardVendorHash  int64                  `json:"rewardVendorHash" yaml:"rewardVendorHash,omitempty"`
	TokenValues       map[string]interface{} `json:"tokenValues" yaml:"tokenValues,omitempty"`
	Vendors           []Vendors              `json:"vendors" yaml:"vendors,omitempty"`
}
type Vendors struct {
	BackgroundImagePath string `json:"backgroundImagePath" yaml:"backgroundImagePath,omitempty"`
	DestinationHash     int64  `json:"destinationHash" yaml:"destinationHash,omitempty"`
	VendorHash          int64  `json:"vendorHash" yaml:"vendorHash,omitempty"`
}

func (d FactionDefinition) Name() string {
	return FactionDefinitionName
}

func (d FactionDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d FactionDefinition) PrettyJson() ([]byte, error) {
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

func (d FactionDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
