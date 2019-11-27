package ghost

import (
	"github.com/shihtzu-systems/bright/pkg/bungo"
)

type Ghost struct {
	Id        string `yaml:"id" json:"id"`
	SessionId string `yaml:"sessionId" json:"session_id"`

	User      User   `yaml:"user" json:"user"`
	CurrentId string `yaml:"id" json:"id"`

	Bright Bright `yaml:"bright" json:"bright"`
	Shadow Shadow `yaml:"shadow" json:"shadow"`

	Try Try `yaml:"try" json:"try"`

	Token BungieToken `yaml:"token" json:"token"`
}

type Bright struct {
	One   BrightCharacter `yaml:"one" json:"one"`
	Two   BrightCharacter `yaml:"two" json:"two"`
	Three BrightCharacter `yaml:"three" json:"three"`
}

type Shadow struct {
	One   ShadowCharacter `yaml:"one" json:"one"`
	Two   ShadowCharacter `yaml:"two" json:"two"`
	Three ShadowCharacter `yaml:"three" json:"three"`
}

type Try struct {
	User      TryUser      `yaml:"user" json:"user"`
	CurrentId string       `yaml:"id" json:"id"`
	One       TryCharacter `yaml:"one" json:"one"`
	Two       TryCharacter `yaml:"two" json:"two"`
	Three     TryCharacter `yaml:"three" json:"three"`
}

type ShadowCharacter bungo.Character
type BrightCharacter bungo.Character
type TryCharacter bungo.Character

type User bungo.CurrentUser
type TryUser bungo.CurrentUser

type BungieToken struct {
	AccessToken      string `yaml:"accessToken" json:"access_token"`
	TokenType        string `yaml:"tokenType" json:"token_type"`
	ExpiresIn        int    `yaml:"expiresIn" json:"expires_in"`
	RefreshToken     string `yaml:"refreshToken" json:"refresh_token"`
	RefreshExpiresIn int    `yaml:"refreshExpiresIn" json:"refresh_expires_in"`
	MembershipId     string `yaml:"membershipId" json:"membership_id"`
}
