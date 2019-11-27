// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const NodeStepSummaryDefinitionName = "node-step-summary"

type NodeStepSummaryDefinitions map[string]NodeStepSummaryDefinition

func (d NodeStepSummaryDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d NodeStepSummaryDefinitions) Values() (out []NodeStepSummaryDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d NodeStepSummaryDefinitions) Find(id int) (out NodeStepSummaryDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return NodeStepSummaryDefinition{}
}

func (d NodeStepSummaryDefinitions) Name() string {
	return NodeStepSummaryDefinitionName
}

type NodeStepSummaryDefinition struct {
}

func (d NodeStepSummaryDefinition) Name() string {
	return NodeStepSummaryDefinitionName
}

func (d NodeStepSummaryDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d NodeStepSummaryDefinition) PrettyJson() ([]byte, error) {
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

func (d NodeStepSummaryDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
