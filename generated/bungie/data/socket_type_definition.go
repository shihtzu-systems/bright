// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SocketTypeDefinitionName = "socket-type"

type SocketTypeDefinitions map[string]SocketTypeDefinition

func (d SocketTypeDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SocketTypeDefinitions) Values() (out []SocketTypeDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SocketTypeDefinitions) Find(id int) (out SocketTypeDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SocketTypeDefinition{}
}

func (d SocketTypeDefinitions) Name() string {
	return SocketTypeDefinitionName
}

type SocketTypeDefinition struct {
	AlwaysRandomizeSockets          bool              `json:"alwaysRandomizeSockets" yaml:"alwaysRandomizeSockets,omitempty"`
	AvoidDuplicatesOnInitialization bool              `json:"avoidDuplicatesOnInitialization" yaml:"avoidDuplicatesOnInitialization,omitempty"`
	Blacklisted                     bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CurrencyScalars                 []interface{}     `json:"currencyScalars" yaml:"currencyScalars,omitempty"`
	DisplayProperties               DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash                            int64             `json:"hash" yaml:"hash,omitempty"`
	HideDuplicateReusablePlugs      bool              `json:"hideDuplicateReusablePlugs" yaml:"hideDuplicateReusablePlugs,omitempty"`
	Index                           int               `json:"index" yaml:"index,omitempty"`
	InsertAction                    InsertAction      `json:"insertAction" yaml:"insertAction,omitempty"`
	IsPreviewEnabled                bool              `json:"isPreviewEnabled" yaml:"isPreviewEnabled,omitempty"`
	OverridesUiAppearance           bool              `json:"overridesUiAppearance" yaml:"overridesUiAppearance,omitempty"`
	PlugWhitelist                   []PlugWhitelist   `json:"plugWhitelist" yaml:"plugWhitelist,omitempty"`
	Redacted                        bool              `json:"redacted" yaml:"redacted,omitempty"`
	SocketCategoryHash              int               `json:"socketCategoryHash" yaml:"socketCategoryHash,omitempty"`
	Visibility                      int               `json:"visibility" yaml:"visibility,omitempty"`
}
type InsertAction struct {
	ActionExecuteSeconds int  `json:"actionExecuteSeconds" yaml:"actionExecuteSeconds,omitempty"`
	ActionSoundHash      int  `json:"actionSoundHash" yaml:"actionSoundHash,omitempty"`
	ActionType           int  `json:"actionType" yaml:"actionType,omitempty"`
	IsPositiveAction     bool `json:"isPositiveAction" yaml:"isPositiveAction,omitempty"`
}
type PlugWhitelist struct {
	CategoryHash       int    `json:"categoryHash" yaml:"categoryHash,omitempty"`
	CategoryIdentifier string `json:"categoryIdentifier" yaml:"categoryIdentifier,omitempty"`
}

func (d SocketTypeDefinition) Name() string {
	return SocketTypeDefinitionName
}

func (d SocketTypeDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SocketTypeDefinition) PrettyJson() ([]byte, error) {
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

func (d SocketTypeDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
