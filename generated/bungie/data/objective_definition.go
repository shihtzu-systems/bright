// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ObjectiveDefinitionName = "objective"

type ObjectiveDefinitions map[string]ObjectiveDefinition

func (d ObjectiveDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ObjectiveDefinitions) Values() (out []ObjectiveDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ObjectiveDefinitions) Find(id int) (out ObjectiveDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ObjectiveDefinition{}
}

func (d ObjectiveDefinitions) Name() string {
	return ObjectiveDefinitionName
}

type ObjectiveDefinition struct {
	AllowNegativeValue            bool              `json:"allowNegativeValue" yaml:"allowNegativeValue,omitempty"`
	AllowOvercompletion           bool              `json:"allowOvercompletion" yaml:"allowOvercompletion,omitempty"`
	AllowValueChangeWhenCompleted bool              `json:"allowValueChangeWhenCompleted" yaml:"allowValueChangeWhenCompleted,omitempty"`
	Blacklisted                   bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CompletedValueStyle           int               `json:"completedValueStyle" yaml:"completedValueStyle,omitempty"`
	CompletionValue               int               `json:"completionValue" yaml:"completionValue,omitempty"`
	DisplayProperties             DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                          int64             `json:"hash" yaml:"hash,omitempty"`
	InProgressValueStyle          int               `json:"inProgressValueStyle" yaml:"inProgressValueStyle,omitempty"`
	Index                         int               `json:"index" yaml:"index,omitempty"`
	IsCountingDownward            bool              `json:"isCountingDownward" yaml:"isCountingDownward,omitempty"`
	IsDisplayOnlyObjective        bool              `json:"isDisplayOnlyObjective" yaml:"isDisplayOnlyObjective,omitempty"`
	LocationHash                  int               `json:"locationHash" yaml:"locationHash,omitempty"`
	MinimumVisibilityThreshold    int               `json:"minimumVisibilityThreshold" yaml:"minimumVisibilityThreshold,omitempty"`
	ProgressDescription           string            `json:"progressDescription" yaml:"progressDescription,omitempty"`
	Redacted                      bool              `json:"redacted" yaml:"redacted,omitempty"`
	Scope                         int               `json:"scope" yaml:"scope,omitempty"`
	ShowValueOnComplete           bool              `json:"showValueOnComplete" yaml:"showValueOnComplete,omitempty"`
	UnlockValueHash               int               `json:"unlockValueHash" yaml:"unlockValueHash,omitempty"`
	ValueStyle                    int               `json:"valueStyle" yaml:"valueStyle,omitempty"`
}

func (d ObjectiveDefinition) Name() string {
	return ObjectiveDefinitionName
}

func (d ObjectiveDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ObjectiveDefinition) PrettyJson() ([]byte, error) {
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

func (d ObjectiveDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
