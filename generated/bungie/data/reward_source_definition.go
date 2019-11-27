// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RewardSourceDefinitionName = "reward-source"

type RewardSourceDefinitions map[string]RewardSourceDefinition

func (d RewardSourceDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RewardSourceDefinitions) Values() (out []RewardSourceDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RewardSourceDefinitions) Find(id int) (out RewardSourceDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RewardSourceDefinition{}
}

func (d RewardSourceDefinitions) Name() string {
	return RewardSourceDefinitionName
}

type RewardSourceDefinition struct {
}

func (d RewardSourceDefinition) Name() string {
	return RewardSourceDefinitionName
}

func (d RewardSourceDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RewardSourceDefinition) PrettyJson() ([]byte, error) {
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

func (d RewardSourceDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
