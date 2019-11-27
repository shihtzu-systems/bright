// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ArtDyeReferenceDefinitionName = "art-dye-reference"

type ArtDyeReferenceDefinitions map[string]ArtDyeReferenceDefinition

func (d ArtDyeReferenceDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ArtDyeReferenceDefinitions) Values() (out []ArtDyeReferenceDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ArtDyeReferenceDefinitions) Find(id int) (out ArtDyeReferenceDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ArtDyeReferenceDefinition{}
}

func (d ArtDyeReferenceDefinitions) Name() string {
	return ArtDyeReferenceDefinitionName
}

type ArtDyeReferenceDefinition struct {
	ArtDyeHash      int   `json:"artDyeHash" yaml:"artDyeHash,omitempty"`
	Blacklisted     bool  `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DyeManifestHash int64 `json:"dyeManifestHash" yaml:"dyeManifestHash,omitempty"`
	Hash            int   `json:"hash" yaml:"hash,omitempty"`
	Index           int   `json:"index" yaml:"index,omitempty"`
	Redacted        bool  `json:"redacted" yaml:"redacted,omitempty"`
}

func (d ArtDyeReferenceDefinition) Name() string {
	return ArtDyeReferenceDefinitionName
}

func (d ArtDyeReferenceDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ArtDyeReferenceDefinition) PrettyJson() ([]byte, error) {
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

func (d ArtDyeReferenceDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
