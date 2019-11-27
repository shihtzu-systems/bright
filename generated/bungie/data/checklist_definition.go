// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ChecklistDefinitionName = "checklist"

type ChecklistDefinitions map[string]ChecklistDefinition

func (d ChecklistDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ChecklistDefinitions) Values() (out []ChecklistDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ChecklistDefinitions) Find(id int) (out ChecklistDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ChecklistDefinition{}
}

func (d ChecklistDefinitions) Name() string {
	return ChecklistDefinitionName
}

type ChecklistDefinition struct {
	Blacklisted       bool               `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties  `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Entries           []ChecklistEntries `json:"entries" yaml:"entries,omitempty"`
	Hash              int64              `json:"hash" yaml:"hash,omitempty"`
	Index             int                `json:"index" yaml:"index,omitempty"`
	Redacted          bool               `json:"redacted" yaml:"redacted,omitempty"`
	Scope             int                `json:"scope" yaml:"scope,omitempty"`
	ViewActionString  string             `json:"viewActionString" yaml:"viewActionString,omitempty"`
}
type ChecklistEntries struct {
	BubbleHash        int64             `json:"bubbleHash" yaml:"bubbleHash,omitempty"`
	DestinationHash   int64             `json:"destinationHash" yaml:"destinationHash,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64             `json:"hash" yaml:"hash,omitempty"`
	LocationHash      int               `json:"locationHash" yaml:"locationHash,omitempty"`
	Scope             int               `json:"scope" yaml:"scope,omitempty"`
}

func (d ChecklistDefinition) Name() string {
	return ChecklistDefinitionName
}

func (d ChecklistDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ChecklistDefinition) PrettyJson() ([]byte, error) {
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

func (d ChecklistDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
