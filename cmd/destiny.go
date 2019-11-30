package cmd

import (
	"fmt"
	"github.com/shihtzu-systems/bright/pkg/brightsvc"
	"github.com/shihtzu-systems/bright/pkg/brightx"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var destinyCommand = &cobra.Command{
	Use: "destiny",
	Run: func(cmd *cobra.Command, args []string) {
		brightx.Destiny(brightx.DestinyArgs{
			Language:        viper.GetString("destinyContent.v5.language"),
			SourceBasePath:  viper.GetString("destinyContent.v5.sourceBasePath"),
			WorkingBasePath: viper.GetString("destinyContent.v5.workingBasePath"),
			BungieNetHost:   viper.GetString("bungie.v2.host"),

			BungieClient: brightsvc.NewBungieClient(brightsvc.NewBungieClientArgs{
				BungieNetUrl: viper.GetString("bungie.v2.url"),
				Host:         viper.GetString("bungie.v2.host"),
				Base:         viper.GetString("bungie.v2.base"),
				ApiKey:       viper.GetString("bungie.v2.apiKey"),

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
					Id: viper.GetString("system.v3.name"),
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
				Trace:  viper.GetBool("system.v3.trace"),
				Debug:  viper.GetBool("system.v3.debug"),
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
		})
	},
}

func init() {
	rootCmd.AddCommand(destinyCommand)
}
