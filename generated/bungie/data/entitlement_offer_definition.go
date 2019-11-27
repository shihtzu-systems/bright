// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const EntitlementOfferDefinitionName = "entitlement-offer"

type EntitlementOfferDefinitions map[string]EntitlementOfferDefinition

func (d EntitlementOfferDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d EntitlementOfferDefinitions) Values() (out []EntitlementOfferDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d EntitlementOfferDefinitions) Find(id int) (out EntitlementOfferDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return EntitlementOfferDefinition{}
}

func (d EntitlementOfferDefinitions) Name() string {
	return EntitlementOfferDefinitionName
}

type EntitlementOfferDefinition struct {
	Blacklisted bool   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	Hash        int    `json:"hash" yaml:"hash,omitempty"`
	Index       int    `json:"index" yaml:"index,omitempty"`
	OfferKey    string `json:"offerKey" yaml:"offerKey,omitempty"`
	Redacted    bool   `json:"redacted" yaml:"redacted,omitempty"`
}

func (d EntitlementOfferDefinition) Name() string {
	return EntitlementOfferDefinitionName
}

func (d EntitlementOfferDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d EntitlementOfferDefinition) PrettyJson() ([]byte, error) {
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

func (d EntitlementOfferDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
