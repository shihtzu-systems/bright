// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityInteractableDefinitionName = "activity-interactable"

type ActivityInteractableDefinitions map[string]ActivityInteractableDefinition

func (d ActivityInteractableDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityInteractableDefinitions) Values() (out []ActivityInteractableDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityInteractableDefinitions) Find(id int) (out ActivityInteractableDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityInteractableDefinition{}
}

func (d ActivityInteractableDefinitions) Name() string {
	return ActivityInteractableDefinitionName
}

type ActivityInteractableDefinition struct {
	Blacklisted bool                          `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Entries     []ActivityInteractableEntries `json:"entries" yaml:"entries,omitempty"`
	Hash        int                           `json:"hash" yaml:"hash,omitempty"`
	Index       int                           `json:"index" yaml:"index,omitempty"`
	Redacted    bool                          `json:"redacted" yaml:"redacted,omitempty"`
}
type ActivityInteractableEntries struct {
	ActivityHash int `json:"activityHash" yaml:"activityHash,omitempty"`
}

func (d ActivityInteractableDefinition) Name() string {
	return ActivityInteractableDefinitionName
}

func (d ActivityInteractableDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityInteractableDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityInteractableDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
