// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardSheetDefinitionName = "reward-sheet"

type RewardSheetDefinitions map[string]RewardSheetDefinition

func (d RewardSheetDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardSheetDefinitions) Values() (out []RewardSheetDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardSheetDefinitions) Find(id int) (out RewardSheetDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardSheetDefinition{}
}

func (d RewardSheetDefinitions) Name() string {
	return RewardSheetDefinitionName
}

type RewardSheetDefinition struct {
	Blacklisted bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int  `json:"hash" yaml:"hash,omitempty"`
	Index       int  `json:"index" yaml:"index,omitempty"`
	Redacted    bool `json:"redacted" yaml:"redacted,omitempty"`
	SheetHash   int  `json:"sheetHash" yaml:"sheetHash,omitempty"`
	SheetIndex  int  `json:"sheetIndex" yaml:"sheetIndex,omitempty"`
}

func (d RewardSheetDefinition) Name() string {
	return RewardSheetDefinitionName
}

func (d RewardSheetDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardSheetDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardSheetDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
