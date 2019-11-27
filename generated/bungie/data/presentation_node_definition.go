// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const PresentationNodeDefinitionName = "presentation-node"

type PresentationNodeDefinitions map[string]PresentationNodeDefinition

func (d PresentationNodeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d PresentationNodeDefinitions) Values() (out []PresentationNodeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d PresentationNodeDefinitions) Find(id int) (out PresentationNodeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return PresentationNodeDefinition{}
}

func (d PresentationNodeDefinitions) Name() string {
	return PresentationNodeDefinitionName
}

type PresentationNodeDefinition struct {
	Blacklisted                     bool                     `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Children                        Children                 `json:"children" yaml:"children,omitempty"`
	DisableChildSubscreenNavigation bool                     `json:"disableChildSubscreenNavigation" yaml:"disableChildSubscreenNavigation,omitempty"`
	DisplayProperties               DisplayProperties        `json:"displayProperties" yaml:"displayProperties,omitempty"`
	DisplayStyle                    int                      `json:"displayStyle" yaml:"displayStyle,omitempty"`
	Hash                            int                      `json:"hash" yaml:"hash,omitempty"`
	Index                           int                      `json:"index" yaml:"index,omitempty"`
	NodeType                        int                      `json:"nodeType" yaml:"nodeType,omitempty"`
	ObjectiveHash                   int                      `json:"objectiveHash" yaml:"objectiveHash,omitempty"`
	OriginalIcon                    string                   `json:"originalIcon" yaml:"originalIcon,omitempty"`
	ParentNodeHashes                []int64                  `json:"parentNodeHashes" yaml:"parentNodeHashes,omitempty"`
	Redacted                        bool                     `json:"redacted" yaml:"redacted,omitempty"`
	Requirements                    PresentationRequirements `json:"requirements" yaml:"requirements,omitempty"`
	RootViewIcon                    string                   `json:"rootViewIcon" yaml:"rootViewIcon,omitempty"`
	Scope                           int                      `json:"scope" yaml:"scope,omitempty"`
	ScreenStyle                     int                      `json:"screenStyle" yaml:"screenStyle,omitempty"`
}
type PresentationNodes struct {
	PresentationNodeHash int64 `json:"presentationNodeHash" yaml:"presentationNodeHash,omitempty"`
}
type Children struct {
	Collectibles      []interface{}       `json:"collectibles" yaml:"collectibles,omitempty"`
	PresentationNodes []PresentationNodes `json:"presentationNodes" yaml:"presentationNodes,omitempty"`
	Records           []interface{}       `json:"records" yaml:"records,omitempty"`
}
type PresentationRequirements struct {
	EntitlementUnavailableMessage string `json:"entitlementUnavailableMessage" yaml:"entitlementUnavailableMessage,omitempty"`
}

func (d PresentationNodeDefinition) Name() string {
	return PresentationNodeDefinitionName
}

func (d PresentationNodeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d PresentationNodeDefinition) PrettyJson() ([]byte, error) {
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

func (d PresentationNodeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
