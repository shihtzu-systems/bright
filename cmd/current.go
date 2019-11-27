package cmd

import (
	. "github.com/shihtzu-systems/bright/pkg/bright"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var currentCommand = &cobra.Command{
	Use: "current",
}

var currentUserCommand = &cobra.Command{
	Use:  "user",
	Args: cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := CurrentUser(CurrentUserArgs{

			OathClientId:     viper.GetString("oauth.v2.clientId"),
			OathClientSecret: viper.GetString("oauth.v2.clientSecret"),

			Auth: BungieToken{
				AccessToken:      viper.GetString("hackToken.v2.accessToken"),
				TokenType:        viper.GetString("hackToken.v2.tokenType"),
				ExpiresIn:        viper.GetInt("hackToken.v2.expiresIn"),
				RefreshToken:     viper.GetString("hackToken.v2.refreshToken"),
				RefreshExpiresIn: viper.GetInt("hackToken.v2.refreshExpiresIn"),
				MembershipId:     viper.GetString("hackToken.v2.membershipId"),
			},

			BungieClient: bungieClient,
			Destiny:      *destiny,
		})
		return err
	},
}

func init() {
	rootCmd.AddCommand(currentCommand)

	currentCommand.AddCommand(currentUserCommand)
}
