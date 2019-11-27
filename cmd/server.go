package cmd

import (
	"fmt"
	"github.com/shihtzu-systems/bright/pkg/bright"
	"github.com/shihtzu-systems/bright/pkg/brightx"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCommand = &cobra.Command{
	Use:  "server",
	Args: cobra.ExactArgs(0),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {

		bungieClient = bright.NewBungieClient(bright.NewBungieClientArgs{
			Host:   viper.GetString("bungie.v2.host"),
			Base:   viper.GetString("bungie.v2.base"),
			ApiKey: viper.GetString("bungie.v2.apiKey"),

			AppVersion: viper.GetString("app.v5.version"),
			AppId:      viper.GetString("app.v5.id"),
			AppUrl:     viper.GetString("app.v5.url"),
			AppEmail:   viper.GetString("app.v5.email"),
		})

		destiny, _, err = bright.NewDestinyContent(bright.NewDestinyContentArgs{
			Language:      viper.GetString("destinyContent.v5.language"),
			BungieClient:  bungieClient,
			BungieNetHost: viper.GetString("bungie.v2.host"),
		})
		if err != nil {
			return err
		}
		return err

	},
	RunE: func(cmd *cobra.Command, args []string) error {

		return bright.Server(bright.ServerArgs{
			Serial: fmt.Sprintf("%s+on.%s.at.%s", version, datestamp, timestamp),

			SessionKey:    viper.GetString("server.v4.sessionKey"),
			SessionSecret: []byte(viper.GetString("server.v4.sessionSecret")),

			OathClientId:     viper.GetString("oauth.v2.clientId"),
			OathClientSecret: viper.GetString("oauth.v2.clientSecret"),
			OathRedirectUrl:  viper.GetString("oauth.v2.redirectUrl"),

			BungieClient: bungieClient,
			Destiny:      *destiny,

			RedisAddress: viper.GetString("redis.v4.address"),
			RedisPort:    viper.GetString("redis.v4.port"),

			HackMode: viper.GetBool("modes.v2.hack"),
			BnetMode: viper.GetBool("modes.v2.bnet"),
			TryMode:  viper.GetBool("modes.v2.try"),

			DadModifier:    viper.GetBool("modifiers.v3.dad"),
			MayhemModifier: viper.GetBool("modifiers.v3.mayhem"),

			SystemName: viper.GetString("system.v2.name"),
			Trace:      viper.GetBool("system.v2.trace"),
			Debug:      viper.GetBool("system.v2.debug"),
			Terminator: viper.GetBool("system.v2.terminator"),

			HackToken: bright.BungieToken{
				AccessToken:      viper.GetString("hackToken.v2.accessToken"),
				TokenType:        viper.GetString("hackToken.v2.tokenType"),
				ExpiresIn:        viper.GetInt("hackToken.v2.expiresIn"),
				RefreshToken:     viper.GetString("hackToken.v2.refreshToken"),
				RefreshExpiresIn: viper.GetInt("hackToken.v2.refreshExpiresIn"),
				MembershipId:     viper.GetString("hackToken.v2.membershipId"),
			},
			GoogleAnalyticsId: viper.GetString("tools.v2.googleAnalyticsId"),
		})
	},
}

var servexCommand = &cobra.Command{
	Use:     "servex",
	Aliases: []string{"x"},
	Run: func(cmd *cobra.Command, args []string) {
		brightx.Serve(brightx.ServeArgs{

			BungieClient: bright.NewBungieClient(bright.NewBungieClientArgs{
				Host:   viper.GetString("bungie.v2.host"),
				Base:   viper.GetString("bungie.v2.base"),
				ApiKey: viper.GetString("bungie.v2.apiKey"),

				AppVersion: viper.GetString("app.v5.version"),
				AppId:      viper.GetString("app.v5.id"),
				AppUrl:     viper.GetString("app.v5.url"),
				AppEmail:   viper.GetString("app.v5.email"),
			}),

			Tower: tower.Tower{
				App: tower.App{
					Version:   version,
					Datestamp: datestamp,
					Timestamp: timestamp,
				},
				System: tower.System{
					Id: viper.GetString("system.v2.name"),
				},
				Serial: fmt.Sprintf("%s+on.%s.at.%s", version, datestamp, timestamp),

				SessionKey:    viper.GetString("server.v4.sessionKey"),
				SessionSecret: []byte(viper.GetString("server.v4.sessionSecret")),

				Oauth: tower.Oauth{
					ClientId:     viper.GetString("oauth.v2.clientId"),
					ClientSecret: viper.GetString("oauth.v2.clientSecret"),
					RedirectUrl:  viper.GetString("oauth.v2.redirectUrl"),
				},

				Hack:   viper.GetBool("modes.v2.hack"),
				Bnet:   viper.GetBool("modes.v2.bnet"),
				Try:    viper.GetBool("modes.v2.try"),
				Trace:  viper.GetBool("system.v2.trace"),
				Debug:  viper.GetBool("system.v2.debug"),
				Dad:    viper.GetBool("modifiers.v3.dad"),
				Mayhem: viper.GetBool("modifiers.v3.mayhem"),

				JaegerAgent:       viper.GetString("tools.v2.jaegerAgent"),
				GoogleAnalyticsId: viper.GetString("tools.v2.googleAnalyticsId"),

				BungieNetUrl: viper.GetString("bungie.v2.url"),

				Language:      viper.GetString("destinyContent.v5.language"),
				BungieNetHost: viper.GetString("bungie.v2.host"),

				TemplatesPath: "tpl",

				Redis: tower.Redis{
					Address:  viper.GetString("redis.v4.address"),
					Port:     viper.GetInt("redis.v4.port"),
					Database: viper.GetInt("redis.v4.database"),
				},
			},

			HackToken: tower.HackToken{
				AccessToken:      viper.GetString("hackToken.v2.accessToken"),
				TokenType:        viper.GetString("hackToken.v2.tokenType"),
				ExpiresIn:        viper.GetInt("hackToken.v2.expiresIn"),
				RefreshToken:     viper.GetString("hackToken.v2.refreshToken"),
				RefreshExpiresIn: viper.GetInt("hackToken.v2.refreshExpiresIn"),
				MembershipId:     viper.GetString("hackToken.v2.membershipId"),
			},
		})
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)

	rootCmd.AddCommand(servexCommand)
}