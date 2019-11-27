// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const UnlockCountMappingDefinitionName = "unlock-count-mapping"

type UnlockCountMappingDefinitions map[string]UnlockCountMappingDefinition

func (d UnlockCountMappingDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d UnlockCountMappingDefinitions) Values() (out []UnlockCountMappingDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d UnlockCountMappingDefinitions) Find(id int) (out UnlockCountMappingDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return UnlockCountMappingDefinition{}
}

func (d UnlockCountMappingDefinitions) Name() string {
	return UnlockCountMappingDefinitionName
}

type UnlockCountMappingDefinition struct {
	Blacklisted     bool  `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash            int64 `json:"hash" yaml:"hash,omitempty"`
	Index           int   `json:"index" yaml:"index,omitempty"`
	Redacted        bool  `json:"redacted" yaml:"redacted,omitempty"`
	UnlockValueHash int   `json:"unlockValueHash" yaml:"unlockValueHash,omitempty"`
}

func (d UnlockCountMappingDefinition) Name() string {
	return UnlockCountMappingDefinitionName
}

func (d UnlockCountMappingDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d UnlockCountMappingDefinition) PrettyJson() ([]byte, error) {
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

func (d UnlockCountMappingDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
