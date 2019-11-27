// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const InventoryBucketDefinitionName = "inventory-bucket"

type InventoryBucketDefinitions map[string]InventoryBucketDefinition

func (d InventoryBucketDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d InventoryBucketDefinitions) Values() (out []InventoryBucketDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d InventoryBucketDefinitions) Find(id int) (out InventoryBucketDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return InventoryBucketDefinition{}
}

func (d InventoryBucketDefinitions) Name() string {
	return InventoryBucketDefinitionName
}

type InventoryBucketDefinition struct {
	Blacklisted            bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	BucketOrder            int               `json:"bucketOrder" yaml:"bucketOrder,omitempty"`
	Category               int               `json:"category" yaml:"category,omitempty"`
	DisplayProperties      DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Enabled                bool              `json:"enabled" yaml:"enabled,omitempty"`
	Fifo                   bool              `json:"fifo" yaml:"fifo,omitempty"`
	HasTransferDestination bool              `json:"hasTransferDestination" yaml:"hasTransferDestination,omitempty"`
	Hash                   int               `json:"hash" yaml:"hash,omitempty"`
	Index                  int               `json:"index" yaml:"index,omitempty"`
	ItemCount              int               `json:"itemCount" yaml:"itemCount,omitempty"`
	Location               int               `json:"location" yaml:"location,omitempty"`
	Redacted               bool              `json:"redacted" yaml:"redacted,omitempty"`
	Scope                  int               `json:"scope" yaml:"scope,omitempty"`
}

func (d InventoryBucketDefinition) Name() string {
	return InventoryBucketDefinitionName
}

func (d InventoryBucketDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d InventoryBucketDefinition) PrettyJson() ([]byte, error) {
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

func (d InventoryBucketDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
