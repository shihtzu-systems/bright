// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const RecordDefinitionName = "record"

type RecordDefinitions map[string]RecordDefinition

func (d RecordDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d RecordDefinitions) Values() (out []RecordDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d RecordDefinitions) Find(id int) (out RecordDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return RecordDefinition{}
}

func (d RecordDefinitions) Name() string {
	return RecordDefinitionName
}

type RecordDefinition struct {
	Blacklisted       bool                   `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CompletionInfo    CompletionInfo         `json:"completionInfo" yaml:"completionInfo,omitempty"`
	DisplayProperties DisplayProperties      `json:"displayProperties" yaml:"displayProperties,omitempty"`
	ExpirationInfo    ExpirationInfo         `json:"expirationInfo" yaml:"expirationInfo,omitempty"`
	Hash              int                    `json:"hash" yaml:"hash,omitempty"`
	Index             int                    `json:"index" yaml:"index,omitempty"`
	IntervalInfo      IntervalInfo           `json:"intervalInfo" yaml:"intervalInfo,omitempty"`
	ObjectiveHashes   []int64                `json:"objectiveHashes" yaml:"objectiveHashes,omitempty"`
	PresentationInfo  RecordPresentationInfo `json:"presentationInfo" yaml:"presentationInfo,omitempty"`
	RecordValueStyle  int                    `json:"recordValueStyle" yaml:"recordValueStyle,omitempty"`
	Redacted          bool                   `json:"redacted" yaml:"redacted,omitempty"`
	Requirements      Requirements           `json:"requirements" yaml:"requirements,omitempty"`
	RewardItems       []interface{}          `json:"rewardItems" yaml:"rewardItems,omitempty"`
	Scope             int                    `json:"scope" yaml:"scope,omitempty"`
	StateInfo         RecordStateInfo        `json:"stateInfo" yaml:"stateInfo,omitempty"`
	TitleInfo         TitleInfo              `json:"titleInfo" yaml:"titleInfo,omitempty"`
}
type CompletionInfo struct {
	ScoreValue                               int  `json:"ScoreValue" yaml:"ScoreValue,omitempty"`
	PartialCompletionObjectiveCountThreshold int  `json:"partialCompletionObjectiveCountThreshold" yaml:"partialCompletionObjectiveCountThreshold,omitempty"`
	ShouldFireToast                          bool `json:"shouldFireToast" yaml:"shouldFireToast,omitempty"`
	ToastStyle                               int  `json:"toastStyle" yaml:"toastStyle,omitempty"`
}
type ExpirationInfo struct {
	Description   string `json:"description" yaml:"description,omitempty"`
	HasExpiration bool   `json:"hasExpiration" yaml:"hasExpiration,omitempty"`
}
type IntervalInfo struct {
	IntervalObjectives                   []interface{} `json:"intervalObjectives" yaml:"intervalObjectives,omitempty"`
	IsIntervalVersionedFromNormalRecord  bool          `json:"isIntervalVersionedFromNormalRecord" yaml:"isIntervalVersionedFromNormalRecord,omitempty"`
	OriginalObjectiveArrayInsertionIndex int           `json:"originalObjectiveArrayInsertionIndex" yaml:"originalObjectiveArrayInsertionIndex,omitempty"`
}
type RecordPresentationInfo struct {
	DisplayStyle                 int     `json:"displayStyle" yaml:"displayStyle,omitempty"`
	ParentPresentationNodeHashes []int64 `json:"parentPresentationNodeHashes" yaml:"parentPresentationNodeHashes,omitempty"`
	PresentationNodeType         int     `json:"presentationNodeType" yaml:"presentationNodeType,omitempty"`
}
type Requirements struct {
	EntitlementUnavailableMessage string `json:"entitlementUnavailableMessage" yaml:"entitlementUnavailableMessage,omitempty"`
}
type RecordStateInfo struct {
	ClaimedUnlockHash  int   `json:"claimedUnlockHash" yaml:"claimedUnlockHash,omitempty"`
	CompleteUnlockHash int   `json:"completeUnlockHash" yaml:"completeUnlockHash,omitempty"`
	FeaturedPriority   int64 `json:"featuredPriority" yaml:"featuredPriority,omitempty"`
}
type TitleInfo struct {
	HasTitle bool `json:"hasTitle" yaml:"hasTitle,omitempty"`
}

func (d RecordDefinition) Name() string {
	return RecordDefinitionName
}

func (d RecordDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d RecordDefinition) PrettyJson() ([]byte, error) {
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

func (d RecordDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
