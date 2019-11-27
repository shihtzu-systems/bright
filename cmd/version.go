package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCommand = &cobra.Command{
	Use:  "version",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("server.v4.sessionKey"))
		fmt.Printf("%s+on.%s.at.%s", version, datestamp, timestamp)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
