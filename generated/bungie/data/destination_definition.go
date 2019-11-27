// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const DestinationDefinitionName = "destination"

type DestinationDefinitions map[string]DestinationDefinition

func (d DestinationDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d DestinationDefinitions) Values() (out []DestinationDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d DestinationDefinitions) Find(id int) (out DestinationDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return DestinationDefinition{}
}

func (d DestinationDefinitions) Name() string {
	return DestinationDefinitionName
}

type DestinationDefinition struct {
	ActivityGraphEntries        []interface{}     `json:"activityGraphEntries" yaml:"activityGraphEntries,omitempty"`
	Blacklisted                 bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	BubbleSettings              []interface{}     `json:"bubbleSettings" yaml:"bubbleSettings,omitempty"`
	Bubbles                     []interface{}     `json:"bubbles" yaml:"bubbles,omitempty"`
	DefaultFreeroamActivityHash int               `json:"defaultFreeroamActivityHash" yaml:"defaultFreeroamActivityHash,omitempty"`
	DisplayProperties           DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                        int               `json:"hash" yaml:"hash,omitempty"`
	Index                       int               `json:"index" yaml:"index,omitempty"`
	PlaceHash                   int64             `json:"placeHash" yaml:"placeHash,omitempty"`
	Redacted                    bool              `json:"redacted" yaml:"redacted,omitempty"`
}

func (d DestinationDefinition) Name() string {
	return DestinationDefinitionName
}

func (d DestinationDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d DestinationDefinition) PrettyJson() ([]byte, error) {
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

func (d DestinationDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
