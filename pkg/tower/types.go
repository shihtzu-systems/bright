package tower

import (
	"github.com/garyburd/redigo/redis"
)

type HackToken struct {
	AccessToken      string `yaml:"accessToken" json:"access_token"`
	TokenType        string `yaml:"tokenType" json:"token_type"`
	ExpiresIn        int    `yaml:"expiresIn" json:"expires_in"`
	RefreshToken     string `yaml:"refreshToken" json:"refresh_token"`
	RefreshExpiresIn int    `yaml:"refreshExpiresIn" json:"refresh_expires_in"`
	MembershipId     string `yaml:"membershipId" json:"membership_id"`
}

type Tower struct {
	App           App    `yaml:"app" json:"app"`
	System        System `yaml:"system" json:"system"`
	Serial        string `yaml:"serial" json:"serial"`
	SessionKey    string `yaml:"sessionKey" json:"session_key"`
	SessionSecret []byte `yaml:"sessionSecret" json:"session_secret"`

	Oauth Oauth `yaml:"oauth" json:"oauth"`

	Hack              bool   `yaml:"hack" json:"hack"`
	Bnet              bool   `yaml:"bnet" json:"bnet"`
	Try               bool   `yaml:"try" json:"try"`
	Trace             bool   `yaml:"trace" json:"trace"`
	Debug             bool   `yaml:"debug" json:"debug"`
	Dad               bool   `yaml:"dad" json:"dad"`
	Mayhem            bool   `yaml:"mayhem" json:"mayhem"`
	GoogleAnalyticsId string `yaml:"googleAnalyticsId" json:"google_analytics_id"`
	JaegerAgent       string `yaml:"jaegerAgent" json:"jaeger_agent"`

	BungieNetUrl string `yaml:"bungieNetUrl" json:"bungie_net_url"`

	Language      string `yaml:"language" json:"language"`
	BungieNetHost string `yaml:"bungieNetHost" json:"bungie_net_host"`

	TemplatesPath string `yaml:"templatesPath" json:"templates_path"`

	Redis Redis `yaml:"redis" json:"redis"`
}

type App struct {
	Version   string `yaml:"version" json:"version"`
	Datestamp string `yaml:"datestamp" json:"datestamp"`
	Timestamp string `yaml:"timestamp" json:"timestamp"`
}

type System struct {
	Id        string `yaml:"id" json:"id"`
	SessionId string `yaml:"sessionId" json:"session_id"`
}

type Oauth struct {
	ClientId     string `yaml:"clientId" json:"client_id"`
	ClientSecret string `yaml:"clientSecret" json:"client_secret"`
	RedirectUrl  string `yaml:"redirectUrl" json:"redirect_url"`
}

type Redis struct {
	Address  string `yaml:"address" json:"address"`
	Port     int    `yaml:"port" json:"port"`
	Database int    `yaml:"database" json:"database"`

	connection redis.Conn
}
