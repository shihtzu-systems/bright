// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ActivityGraphDefinitionName = "activity-graph"

type ActivityGraphDefinitions map[string]ActivityGraphDefinition

func (d ActivityGraphDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ActivityGraphDefinitions) Values() (out []ActivityGraphDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ActivityGraphDefinitions) Find(id int) (out ActivityGraphDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ActivityGraphDefinition{}
}

func (d ActivityGraphDefinitions) Name() string {
	return ActivityGraphDefinitionName
}

type ActivityGraphDefinition struct {
	ArtElements         []interface{} `json:"artElements" yaml:"artElements,omitempty"`
	Blacklisted         bool          `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Connections         []interface{} `json:"connections" yaml:"connections,omitempty"`
	DisplayObjectives   []interface{} `json:"displayObjectives" yaml:"displayObjectives,omitempty"`
	DisplayProgressions []interface{} `json:"displayProgressions" yaml:"displayProgressions,omitempty"`
	Hash                int           `json:"hash" yaml:"hash,omitempty"`
	IgnoreForMilestones bool          `json:"ignoreForMilestones" yaml:"ignoreForMilestones,omitempty"`
	Index               int           `json:"index" yaml:"index,omitempty"`
	LinkedGraphs        []interface{} `json:"linkedGraphs" yaml:"linkedGraphs,omitempty"`
	Nodes               []interface{} `json:"nodes" yaml:"nodes,omitempty"`
	Redacted            bool          `json:"redacted" yaml:"redacted,omitempty"`
	UiScreen            int           `json:"uiScreen" yaml:"uiScreen,omitempty"`
}

func (d ActivityGraphDefinition) Name() string {
	return ActivityGraphDefinitionName
}

func (d ActivityGraphDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ActivityGraphDefinition) PrettyJson() ([]byte, error) {
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

func (d ActivityGraphDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
