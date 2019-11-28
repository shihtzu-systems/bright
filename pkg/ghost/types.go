package ghost

import (
	"github.com/shihtzu-systems/bright/pkg/bungo"
)

type Ghost struct {
	Id        string `yaml:"id" json:"id"`
	SessionId string `yaml:"sessionId" json:"session_id"`

	Bright Soul `yaml:"bright" json:"bright"`
	Shadow Soul `yaml:"shadow" json:"shadow"`
	Try    Soul `yaml:"try" json:"try"`

	Token BungieToken `yaml:"token" json:"token"`
}

type Soul struct {
	Gamer     bungo.Gamer    `yaml:"gamer" json:"gamer"`
	Possessed string         `yaml:"possessed" json:"possessed"`
	One       bungo.Guardian `yaml:"one" json:"one"`
	Two       bungo.Guardian `yaml:"two" json:"two"`
	Three     bungo.Guardian `yaml:"three" json:"three"`
}

type BungieToken struct {
	AccessToken      string `yaml:"accessToken" json:"access_token"`
	TokenType        string `yaml:"tokenType" json:"token_type"`
	ExpiresIn        int    `yaml:"expiresIn" json:"expires_in"`
	RefreshToken     string `yaml:"refreshToken" json:"refresh_token"`
	RefreshExpiresIn int    `yaml:"refreshExpiresIn" json:"refresh_expires_in"`
	MembershipId     string `yaml:"membershipId" json:"membership_id"`
}
