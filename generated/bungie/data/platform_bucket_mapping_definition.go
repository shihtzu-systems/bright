// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const PlatformBucketMappingDefinitionName = "platform-bucket-mapping"

type PlatformBucketMappingDefinitions map[string]PlatformBucketMappingDefinition

func (d PlatformBucketMappingDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d PlatformBucketMappingDefinitions) Values() (out []PlatformBucketMappingDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d PlatformBucketMappingDefinitions) Find(id int) (out PlatformBucketMappingDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return PlatformBucketMappingDefinition{}
}

func (d PlatformBucketMappingDefinitions) Name() string {
	return PlatformBucketMappingDefinitionName
}

type PlatformBucketMappingDefinition struct {
	Blacklisted    bool `json:"blacklisted" yaml:"blacklisted,omitempty"`
	BucketHash     int  `json:"bucketHash" yaml:"bucketHash,omitempty"`
	Hash           int  `json:"hash" yaml:"hash,omitempty"`
	Index          int  `json:"index" yaml:"index,omitempty"`
	MembershipType int  `json:"membershipType" yaml:"membershipType,omitempty"`
	Redacted       bool `json:"redacted" yaml:"redacted,omitempty"`
}

func (d PlatformBucketMappingDefinition) Name() string {
	return PlatformBucketMappingDefinitionName
}

func (d PlatformBucketMappingDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d PlatformBucketMappingDefinition) PrettyJson() ([]byte, error) {
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

func (d PlatformBucketMappingDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
