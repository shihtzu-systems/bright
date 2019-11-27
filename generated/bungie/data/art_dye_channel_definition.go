// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ArtDyeChannelDefinitionName = "art-dye-channel"

type ArtDyeChannelDefinitions map[string]ArtDyeChannelDefinition

func (d ArtDyeChannelDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ArtDyeChannelDefinitions) Values() (out []ArtDyeChannelDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ArtDyeChannelDefinitions) Find(id int) (out ArtDyeChannelDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ArtDyeChannelDefinition{}
}

func (d ArtDyeChannelDefinitions) Name() string {
	return ArtDyeChannelDefinitionName
}

type ArtDyeChannelDefinition struct {
	Blacklisted bool  `json:"blacklisted" yaml:"blacklisted,omitempty"`
	ChannelHash int64 `json:"channelHash" yaml:"channelHash,omitempty"`
	Hash        int64 `json:"hash" yaml:"hash,omitempty"`
	Index       int   `json:"index" yaml:"index,omitempty"`
	Redacted    bool  `json:"redacted" yaml:"redacted,omitempty"`
}

func (d ArtDyeChannelDefinition) Name() string {
	return ArtDyeChannelDefinitionName
}

func (d ArtDyeChannelDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ArtDyeChannelDefinition) PrettyJson() ([]byte, error) {
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

func (d ArtDyeChannelDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
