// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const CollectibleDefinitionName = "collectible"

type CollectibleDefinitions map[string]CollectibleDefinition

func (d CollectibleDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d CollectibleDefinitions) Values() (out []CollectibleDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d CollectibleDefinitions) Find(id int) (out CollectibleDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return CollectibleDefinition{}
}

func (d CollectibleDefinitions) Name() string {
	return CollectibleDefinitionName
}

type CollectibleDefinition struct {
	AcquisitionInfo   AcquisitionInfo             `json:"acquisitionInfo" yaml:"acquisitionInfo,omitempty"`
	Blacklisted       bool                        `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties           `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int64                       `json:"hash" yaml:"hash,omitempty"`
	Index             int                         `json:"index" yaml:"index,omitempty"`
	ItemHash          int                         `json:"itemHash" yaml:"itemHash,omitempty"`
	PresentationInfo  CollectiblePresentationInfo `json:"presentationInfo" yaml:"presentationInfo,omitempty"`
	Redacted          bool                        `json:"redacted" yaml:"redacted,omitempty"`
	Scope             int                         `json:"scope" yaml:"scope,omitempty"`
	SourceHash        int                         `json:"sourceHash" yaml:"sourceHash,omitempty"`
	SourceString      string                      `json:"sourceString" yaml:"sourceString,omitempty"`
	StateInfo         CollectibleStateInfo        `json:"stateInfo" yaml:"stateInfo,omitempty"`
}
type AcquisitionInfo struct {
	AcquireMaterialRequirementHash int  `json:"acquireMaterialRequirementHash" yaml:"acquireMaterialRequirementHash,omitempty"`
	RunOnlyAcquisitionRewardSite   bool `json:"runOnlyAcquisitionRewardSite" yaml:"runOnlyAcquisitionRewardSite,omitempty"`
}
type CollectiblePresentationInfo struct {
	DisplayStyle                 int   `json:"displayStyle" yaml:"displayStyle,omitempty"`
	ParentPresentationNodeHashes []int `json:"parentPresentationNodeHashes" yaml:"parentPresentationNodeHashes,omitempty"`
	PresentationNodeType         int   `json:"presentationNodeType" yaml:"presentationNodeType,omitempty"`
}
type CollectibleRequirements struct {
	EntitlementUnavailableMessage string `json:"entitlementUnavailableMessage" yaml:"entitlementUnavailableMessage,omitempty"`
}
type CollectibleStateInfo struct {
	Requirements CollectibleRequirements `json:"requirements" yaml:"requirements,omitempty"`
}

func (d CollectibleDefinition) Name() string {
	return CollectibleDefinitionName
}

func (d CollectibleDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d CollectibleDefinition) PrettyJson() ([]byte, error) {
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

func (d CollectibleDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
