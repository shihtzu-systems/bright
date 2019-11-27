// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const ReportReasonCategoryDefinitionName = "report-reason-category"

type ReportReasonCategoryDefinitions map[string]ReportReasonCategoryDefinition

func (d ReportReasonCategoryDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ReportReasonCategoryDefinitions) Values() (out []ReportReasonCategoryDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ReportReasonCategoryDefinitions) Find(id int) (out ReportReasonCategoryDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ReportReasonCategoryDefinition{}
}

func (d ReportReasonCategoryDefinitions) Name() string {
	return ReportReasonCategoryDefinitionName
}

type ReportReasonCategoryDefinition struct {
	Blacklisted       bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Hash              int               `json:"hash" yaml:"hash,omitempty"`
	Index             int               `json:"index" yaml:"index,omitempty"`
	Reasons           map[string]Reason `json:"reasons" yaml:"reasons,omitempty"`
	Redacted          bool              `json:"redacted" yaml:"redacted,omitempty"`
}
type Reason struct {
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	ReasonHash        int               `json:"reasonHash" yaml:"reasonHash,omitempty"`
}

func (d ReportReasonCategoryDefinition) Name() string {
	return ReportReasonCategoryDefinitionName
}

func (d ReportReasonCategoryDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ReportReasonCategoryDefinition) PrettyJson() ([]byte, error) {
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

func (d ReportReasonCategoryDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
