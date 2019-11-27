// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const MilestoneDefinitionName = "milestone"

type MilestoneDefinitions map[string]MilestoneDefinition

func (d MilestoneDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d MilestoneDefinitions) Values() (out []MilestoneDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d MilestoneDefinitions) Find(id int) (out MilestoneDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return MilestoneDefinition{}
}

func (d MilestoneDefinitions) Name() string {
	return MilestoneDefinitionName
}

type MilestoneDefinition struct {
	Blacklisted                     bool              `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DefaultOrder                    int               `json:"defaultOrder" yaml:"defaultOrder,omitempty"`
	DisplayProperties               DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	ExplorePrioritizesActivityImage bool              `json:"explorePrioritizesActivityImage" yaml:"explorePrioritizesActivityImage,omitempty"`
	HasPredictableDates             bool              `json:"hasPredictableDates" yaml:"hasPredictableDates,omitempty"`
	Hash                            int               `json:"hash" yaml:"hash,omitempty"`
	Index                           int               `json:"index" yaml:"index,omitempty"`
	IsInGameMilestone               bool              `json:"isInGameMilestone" yaml:"isInGameMilestone,omitempty"`
	MilestoneType                   int               `json:"milestoneType" yaml:"milestoneType,omitempty"`
	Quests                          map[string]Quest  `json:"quests" yaml:"quests,omitempty"`
	Recruitable                     bool              `json:"recruitable" yaml:"recruitable,omitempty"`
	Redacted                        bool              `json:"redacted" yaml:"redacted,omitempty"`
	ShowInExplorer                  bool              `json:"showInExplorer" yaml:"showInExplorer,omitempty"`
	ShowInMilestones                bool              `json:"showInMilestones" yaml:"showInMilestones,omitempty"`
}
type MilestoneItems struct {
	ItemHash int `json:"itemHash" yaml:"itemHash,omitempty"`
	Quantity int `json:"quantity" yaml:"quantity,omitempty"`
}
type QuestRewards struct {
	Items []MilestoneItems `json:"items" yaml:"items,omitempty"`
}
type Quest struct {
	DisplayProperties       DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	QuestItemHash           int               `json:"questItemHash" yaml:"questItemHash,omitempty"`
	QuestRewards            QuestRewards      `json:"questRewards" yaml:"questRewards,omitempty"`
	TrackingUnlockValueHash int               `json:"trackingUnlockValueHash" yaml:"trackingUnlockValueHash,omitempty"`
}

func (d MilestoneDefinition) Name() string {
	return MilestoneDefinitionName
}

func (d MilestoneDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d MilestoneDefinition) PrettyJson() ([]byte, error) {
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

func (d MilestoneDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
