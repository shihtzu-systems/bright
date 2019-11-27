// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const SandboxPatternDefinitionName = "sandbox-pattern"

type SandboxPatternDefinitions map[string]SandboxPatternDefinition

func (d SandboxPatternDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d SandboxPatternDefinitions) Values() (out []SandboxPatternDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d SandboxPatternDefinitions) Find(id int) (out SandboxPatternDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return SandboxPatternDefinition{}
}

func (d SandboxPatternDefinitions) Name() string {
	return SandboxPatternDefinitionName
}

type SandboxPatternDefinition struct {
	Blacklisted                bool     `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Filters                    []Filter `json:"filters" yaml:"filters,omitempty"`
	Hash                       int64    `json:"hash" yaml:"hash,omitempty"`
	Index                      int      `json:"index" yaml:"index,omitempty"`
	PatternGlobalTagIdHash     int64    `json:"patternGlobalTagIdHash" yaml:"patternGlobalTagIdHash,omitempty"`
	PatternHash                int64    `json:"patternHash" yaml:"patternHash,omitempty"`
	Redacted                   bool     `json:"redacted" yaml:"redacted,omitempty"`
	WeaponContentGroupHash     int64    `json:"weaponContentGroupHash" yaml:"weaponContentGroupHash,omitempty"`
	WeaponTranslationGroupHash int64    `json:"weaponTranslationGroupHash" yaml:"weaponTranslationGroupHash,omitempty"`
	WeaponType                 int      `json:"weaponType" yaml:"weaponType,omitempty"`
	WeaponTypeHash             int64    `json:"weaponTypeHash" yaml:"weaponTypeHash,omitempty"`
}
type Filter struct {
	ArrangementIndexByStatValue map[string]int `json:"arrangementIndexByStatValue" yaml:"arrangementIndexByStatValue,omitempty"`
	ArtArrangementRegionHash    int64          `json:"artArrangementRegionHash" yaml:"artArrangementRegionHash,omitempty"`
	ArtArrangementRegionIndex   int            `json:"artArrangementRegionIndex" yaml:"artArrangementRegionIndex,omitempty"`
	StatHash                    int64          `json:"statHash" yaml:"statHash,omitempty"`
}

func (d SandboxPatternDefinition) Name() string {
	return SandboxPatternDefinitionName
}

func (d SandboxPatternDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d SandboxPatternDefinition) PrettyJson() ([]byte, error) {
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

func (d SandboxPatternDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
