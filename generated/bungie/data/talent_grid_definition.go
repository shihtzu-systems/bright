// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const TalentGridDefinitionName = "talent-grid"

type TalentGridDefinitions map[string]TalentGridDefinition

func (d TalentGridDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d TalentGridDefinitions) Values() (out []TalentGridDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d TalentGridDefinitions) Find(id int) (out TalentGridDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return TalentGridDefinition{}
}

func (d TalentGridDefinitions) Name() string {
	return TalentGridDefinitionName
}

type TalentGridDefinition struct {
	Blacklisted                       bool             `json:"blacklisted" yaml:"blacklisted,omitempty"`
	CalcMaxGridLevel                  int              `json:"calcMaxGridLevel" yaml:"calcMaxGridLevel,omitempty"`
	CalcProgressToMaxLevel            int              `json:"calcProgressToMaxLevel" yaml:"calcProgressToMaxLevel,omitempty"`
	ExclusiveSets                     []ExclusiveSet   `json:"exclusiveSets" yaml:"exclusiveSets,omitempty"`
	GridLevelPerColumn                int              `json:"gridLevelPerColumn" yaml:"gridLevelPerColumn,omitempty"`
	Groups                            map[string]Group `json:"groups" yaml:"groups,omitempty"`
	Hash                              int64            `json:"hash" yaml:"hash,omitempty"`
	IndependentNodeIndexes            []float64        `json:"independentNodeIndexes" yaml:"independentNodeIndexes,omitempty"`
	Index                             int              `json:"index" yaml:"index,omitempty"`
	MaxGridLevel                      int              `json:"maxGridLevel" yaml:"maxGridLevel,omitempty"`
	MaximumRandomMaterialRequirements int              `json:"maximumRandomMaterialRequirements" yaml:"maximumRandomMaterialRequirements,omitempty"`
	NodeCategories                    []NodeCategory   `json:"nodeCategories" yaml:"nodeCategories,omitempty"`
	Nodes                             []Node           `json:"nodes" yaml:"nodes,omitempty"`
	ProgressionHash                   int              `json:"progressionHash" yaml:"progressionHash,omitempty"`
	Redacted                          bool             `json:"redacted" yaml:"redacted,omitempty"`
}
type Group struct {
	GroupHash           int64   `json:"groupHash" yaml:"groupHash,omitempty"`
	NodeHashes          []int64 `json:"nodeHashes" yaml:"nodeHashes,omitempty"`
	OpposingGroupHashes []int64 `json:"opposingGroupHashes" yaml:"opposingGroupHashes,omitempty"`
	OpposingNodeHashes  []int64 `json:"opposingNodeHashes" yaml:"opposingNodeHashes,omitempty"`
}
type ExclusiveSet struct {
}
type NodeCategory struct {
	Identifier        string            `json:"identifier" yaml:"identifier,omitempty"`
	IsLoreDriven      bool              `json:"isLoreDriven" yaml:"isLoreDriven,omitempty"`
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	NodeHashes        []int64           `json:"nodeHashes" yaml:"nodeHashes,omitempty"`
}
type ActivationRequirement struct {
	ExclusiveSetRequiredHash  int     `json:"exclusiveSetRequiredHash" yaml:"exclusiveSetRequiredHash,omitempty"`
	GridLevel                 int     `json:"gridLevel" yaml:"gridLevel,omitempty"`
	MaterialRequirementHashes []int64 `json:"materialRequirementHashes" yaml:"materialRequirementHashes,omitempty"`
}
type Node struct {
	AutoUnlocks                            bool    `json:"autoUnlocks" yaml:"autoUnlocks,omitempty"`
	BinaryPairNodeIndex                    int     `json:"binaryPairNodeIndex" yaml:"binaryPairNodeIndex,omitempty"`
	Column                                 int     `json:"column" yaml:"column,omitempty"`
	ExclusiveSetHash                       int     `json:"exclusiveSetHash" yaml:"exclusiveSetHash,omitempty"`
	ExclusiveWithNodeHashes                []int64 `json:"exclusiveWithNodeHashes" yaml:"exclusiveWithNodeHashes,omitempty"`
	GroupScopeIndex                        int     `json:"groupScopeIndex" yaml:"groupScopeIndex,omitempty"`
	IgnoreForCompletion                    bool    `json:"ignoreForCompletion" yaml:"ignoreForCompletion,omitempty"`
	IsRandom                               bool    `json:"isRandom" yaml:"isRandom,omitempty"`
	IsRandomRepurchasable                  bool    `json:"isRandomRepurchasable" yaml:"isRandomRepurchasable,omitempty"`
	IsRealStepSelectionRandom              bool    `json:"isRealStepSelectionRandom" yaml:"isRealStepSelectionRandom,omitempty"`
	LastStepRepeats                        bool    `json:"lastStepRepeats" yaml:"lastStepRepeats,omitempty"`
	LayoutIdentifier                       string  `json:"layoutIdentifier" yaml:"layoutIdentifier,omitempty"`
	NodeHash                               int     `json:"nodeHash" yaml:"nodeHash,omitempty"`
	NodeIndex                              int     `json:"nodeIndex" yaml:"nodeIndex,omitempty"`
	NodeStyleIdentifier                    string  `json:"nodeStyleIdentifier" yaml:"nodeStyleIdentifier,omitempty"`
	OriginalNodeHash                       int     `json:"originalNodeHash" yaml:"originalNodeHash,omitempty"`
	PrerequisiteNodeIndexes                []int32 `json:"prerequisiteNodeIndexes" yaml:"prerequisiteNodeIndexes,omitempty"`
	RandomStartProgressionBarAtProgression int     `json:"randomStartProgressionBarAtProgression" yaml:"randomStartProgressionBarAtProgression,omitempty"`
	Row                                    int     `json:"row" yaml:"row,omitempty"`
	Steps                                  []Step  `json:"steps" yaml:"steps,omitempty"`
	TalentNodeTypes                        int     `json:"talentNodeTypes" yaml:"talentNodeTypes,omitempty"`
	GroupHash                              int64   `json:"groupHash" yaml:"groupHash,omitempty"`
}
type Step struct {
	ActivationRequirement         ActivationRequirement `json:"activationRequirement" yaml:"activationRequirement,omitempty"`
	AffectsLevel                  bool                  `json:"affectsLevel" yaml:"affectsLevel,omitempty"`
	AffectsQuality                bool                  `json:"affectsQuality" yaml:"affectsQuality,omitempty"`
	CanActivateNextStep           bool                  `json:"canActivateNextStep" yaml:"canActivateNextStep,omitempty"`
	DamageType                    int                   `json:"damageType" yaml:"damageType,omitempty"`
	DisplayProperties             DisplayProperties     `json:"displayProperties" yaml:"displayProperties,omitempty"`
	InteractionDescription        string                `json:"interactionDescription" yaml:"interactionDescription,omitempty"`
	IsNextStepRandom              bool                  `json:"isNextStepRandom" yaml:"isNextStepRandom,omitempty"`
	NextStepIndex                 int                   `json:"nextStepIndex" yaml:"nextStepIndex,omitempty"`
	NodeStepHash                  int                   `json:"nodeStepHash" yaml:"nodeStepHash,omitempty"`
	PerkHashes                    []int64               `json:"perkHashes" yaml:"perkHashes,omitempty"`
	StartProgressionBarAtProgress int                   `json:"startProgressionBarAtProgress" yaml:"startProgressionBarAtProgress,omitempty"`
	StatHashes                    []int64               `json:"statHashes" yaml:"statHashes,omitempty"`
	StepIndex                     int                   `json:"stepIndex" yaml:"stepIndex,omitempty"`
	TruePropertyIndex             int                   `json:"truePropertyIndex" yaml:"truePropertyIndex,omitempty"`
	TrueStepIndex                 int                   `json:"trueStepIndex" yaml:"trueStepIndex,omitempty"`
}

func (d TalentGridDefinition) Name() string {
	return TalentGridDefinitionName
}

func (d TalentGridDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d TalentGridDefinition) PrettyJson() ([]byte, error) {
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

func (d TalentGridDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
