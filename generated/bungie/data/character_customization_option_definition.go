// Code generated; DANGER ZONE FOR EDITS

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

const CharacterCustomizationOptionDefinitionName = "character-customization-option"

type CharacterCustomizationOptionDefinitions map[string]CharacterCustomizationOptionDefinition

func (d CharacterCustomizationOptionDefinitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d CharacterCustomizationOptionDefinitions) Values() (out []CharacterCustomizationOptionDefinition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d CharacterCustomizationOptionDefinitions) Find(id int) (out CharacterCustomizationOptionDefinition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return CharacterCustomizationOptionDefinition{}
}

func (d CharacterCustomizationOptionDefinitions) Name() string {
	return CharacterCustomizationOptionDefinitionName
}

type CharacterCustomizationOptionDefinition struct {
	Blacklisted                       bool                `json:"blacklisted" yaml:"blacklisted,omitempty"`
	DecalColorOptions                 DecalColorOptions   `json:"decalColorOptions" yaml:"decalColorOptions,omitempty"`
	DecalOptions                      DecalOptions        `json:"decalOptions" yaml:"decalOptions,omitempty"`
	DisplayProperties                 DisplayProperties   `json:"displayProperties" yaml:"displayProperties,omitempty"`
	EyeColorOptions                   EyeColorOptions     `json:"eyeColorOptions" yaml:"eyeColorOptions,omitempty"`
	FaceOptionCinematicHostPatternIds []interface{}       `json:"faceOptionCinematicHostPatternIds" yaml:"faceOptionCinematicHostPatternIds,omitempty"`
	FaceOptions                       FaceOptions         `json:"faceOptions" yaml:"faceOptions,omitempty"`
	FeatureColorOptions               FeatureColorOptions `json:"featureColorOptions" yaml:"featureColorOptions,omitempty"`
	FeatureOptions                    FeatureOptions      `json:"featureOptions" yaml:"featureOptions,omitempty"`
	GenderHash                        int64               `json:"genderHash" yaml:"genderHash,omitempty"`
	HairColorOptions                  HairColorOptions    `json:"hairColorOptions" yaml:"hairColorOptions,omitempty"`
	HairOptions                       HairOptions         `json:"hairOptions" yaml:"hairOptions,omitempty"`
	Hash                              int                 `json:"hash" yaml:"hash,omitempty"`
	HelmetPreferences                 HelmetPreferences   `json:"helmetPreferences" yaml:"helmetPreferences,omitempty"`
	Index                             int                 `json:"index" yaml:"index,omitempty"`
	LipColorOptions                   LipColorOptions     `json:"lipColorOptions" yaml:"lipColorOptions,omitempty"`
	PersonalityOptions                PersonalityOptions  `json:"personalityOptions" yaml:"personalityOptions,omitempty"`
	RaceHash                          int                 `json:"raceHash" yaml:"raceHash,omitempty"`
	Redacted                          bool                `json:"redacted" yaml:"redacted,omitempty"`
	SkinColorOptions                  SkinColorOptions    `json:"skinColorOptions" yaml:"skinColorOptions,omitempty"`
}
type FloatSliceOptions struct {
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Value             []float64         `json:"value" yaml:"value,omitempty"`
}
type IntOptions struct {
	DisplayProperties DisplayProperties `json:"displayProperties" yaml:"displayProperties,omitempty"`
	Value             int64             `json:"value" yaml:"value,omitempty"`
}
type DecalColorOptions struct {
	CustomizationCategoryHash int          `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type DecalOptions struct {
	CustomizationCategoryHash int          `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type EyeColorOptions struct {
	CustomizationCategoryHash int64        `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type FaceOptions struct {
	CustomizationCategoryHash int64        `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type FeatureColorOptions struct {
	CustomizationCategoryHash int           `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string        `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []interface{} `json:"options" yaml:"options,omitempty"`
}
type FeatureOptions struct {
	CustomizationCategoryHash int          `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type HairColorOptions struct {
	CustomizationCategoryHash float64             `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string              `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []FloatSliceOptions `json:"options" yaml:"options,omitempty"`
}
type HairOptions struct {
	CustomizationCategoryHash int          `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type HelmetPreferences struct {
	CustomizationCategoryHash int           `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string        `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []interface{} `json:"options" yaml:"options,omitempty"`
}
type LipColorOptions struct {
	CustomizationCategoryHash int64        `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}
type PersonalityOptions struct {
	CustomizationCategoryHash int           `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string        `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []interface{} `json:"options" yaml:"options,omitempty"`
}
type SkinColorOptions struct {
	CustomizationCategoryHash int64        `json:"customizationCategoryHash" yaml:"customizationCategoryHash,omitempty"`
	DisplayName               string       `json:"displayName" yaml:"displayName,omitempty"`
	Options                   []IntOptions `json:"options" yaml:"options,omitempty"`
}

func (d CharacterCustomizationOptionDefinition) Name() string {
	return CharacterCustomizationOptionDefinitionName
}

func (d CharacterCustomizationOptionDefinition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d CharacterCustomizationOptionDefinition) PrettyJson() ([]byte, error) {
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

func (d CharacterCustomizationOptionDefinition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}
