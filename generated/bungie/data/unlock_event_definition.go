// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const UnlockEventDefinitionName = "unlock-event"

type UnlockEventDefinitions map[string]UnlockEventDefinition

func (d UnlockEventDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d UnlockEventDefinitions) Values() (out []UnlockEventDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d UnlockEventDefinitions) Find(id int) (out UnlockEventDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return UnlockEventDefinition{}
}

func (d UnlockEventDefinitions) Name() string {
	return UnlockEventDefinitionName
}

type UnlockEventDefinition struct {
	Blacklisted                        bool            `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash                               int64           `json:"hash" yaml:"hash,omitempty"`
	Index                              int             `json:"index" yaml:"index,omitempty"`
	NewSequenceRewardSiteHash          int             `json:"newSequenceRewardSiteHash" yaml:"newSequenceRewardSiteHash,omitempty"`
	Redacted                           bool            `json:"redacted" yaml:"redacted,omitempty"`
	SequenceLastUpdatedUnlockValueHash int             `json:"sequenceLastUpdatedUnlockValueHash" yaml:"sequenceLastUpdatedUnlockValueHash,omitempty"`
	SequenceUnlockValueHash            int             `json:"sequenceUnlockValueHash" yaml:"sequenceUnlockValueHash,omitempty"`
	UnlockEntries                      []UnlockEntries `json:"unlockEntries" yaml:"unlockEntries,omitempty"`
}
type UnlockEntries struct {
	ClearedValue    float64 `json:"clearedValue" yaml:"clearedValue,omitempty"`
	SelectedValue   float64 `json:"selectedValue" yaml:"selectedValue,omitempty"`
	UnlockHash      float64 `json:"unlockHash" yaml:"unlockHash,omitempty"`
	UnlockValueHash float64 `json:"unlockValueHash" yaml:"unlockValueHash,omitempty"`
}

func (d UnlockEventDefinition) Name() string {
	return UnlockEventDefinitionName
}

func (d UnlockEventDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d UnlockEventDefinition) PrettyJson() ([]byte, error) {
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

func (d UnlockEventDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
