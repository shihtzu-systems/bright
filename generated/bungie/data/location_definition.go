// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const LocationDefinitionName = "location"

type LocationDefinitions map[string]LocationDefinition

func (d LocationDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d LocationDefinitions) Values() (out []LocationDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d LocationDefinitions) Find(id int) (out LocationDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return LocationDefinition{}
}

func (d LocationDefinitions) Name() string {
	return LocationDefinitionName
}

type LocationDefinition struct {
	Blacklisted      bool               `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash             int                `json:"hash" yaml:"hash,omitempty"`
	Index            int                `json:"index" yaml:"index,omitempty"`
	LocationReleases []LocationReleases `json:"locationReleases" yaml:"locationReleases,omitempty"`
	Redacted         bool               `json:"redacted" yaml:"redacted,omitempty"`
	VendorHash       int                `json:"vendorHash" yaml:"vendorHash,omitempty"`
}
type LocationReleases struct {
	ActivityBubbleName      int64             `json:"activityBubbleName" yaml:"activityBubbleName,omitempty"`
	ActivityGraphHash       int64             `json:"activityGraphHash" yaml:"activityGraphHash,omitempty"`
	ActivityGraphNodeHash   int               `json:"activityGraphNodeHash" yaml:"activityGraphNodeHash,omitempty"`
	ActivityHash            int               `json:"activityHash" yaml:"activityHash,omitempty"`
	ActivityPathBundle      int               `json:"activityPathBundle" yaml:"activityPathBundle,omitempty"`
	ActivityPathDestination int               `json:"activityPathDestination" yaml:"activityPathDestination,omitempty"`
	DestinationHash         int               `json:"destinationHash" yaml:"destinationHash,omitempty"`
	DisplayProperties       DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	NavPointType            int               `json:"navPointType" yaml:"navPointType,omitempty"`
	SpawnPoint              int               `json:"spawnPoint" yaml:"spawnPoint,omitempty"`
	WorldPosition           []int             `json:"worldPosition" yaml:"worldPosition,omitempty"`
}

func (d LocationDefinition) Name() string {
	return LocationDefinitionName
}

func (d LocationDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d LocationDefinition) PrettyJson() ([]byte, error) {
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

func (d LocationDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
