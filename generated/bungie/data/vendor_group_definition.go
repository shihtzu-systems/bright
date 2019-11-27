// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const VendorGroupDefinitionName = "vendor-group"

type VendorGroupDefinitions map[string]VendorGroupDefinition

func (d VendorGroupDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d VendorGroupDefinitions) Values() (out []VendorGroupDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d VendorGroupDefinitions) Find(id int) (out VendorGroupDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return VendorGroupDefinition{}
}

func (d VendorGroupDefinitions) Name() string {
	return VendorGroupDefinitionName
}

type VendorGroupDefinition struct {
	Blacklisted  bool   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CategoryName string `json:"categoryName" yaml:"categoryName,omitempty"`
	Hash         int    `json:"hash" yaml:"hash,omitempty"`
	Index        int    `json:"index" yaml:"index,omitempty"`
	Order        int    `json:"order" yaml:"order,omitempty"`
	Redacted     bool   `json:"redacted" yaml:"redacted,omitempty"`
}

func (d VendorGroupDefinition) Name() string {
	return VendorGroupDefinitionName
}

func (d VendorGroupDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d VendorGroupDefinition) PrettyJson() ([]byte, error) {
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

func (d VendorGroupDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
