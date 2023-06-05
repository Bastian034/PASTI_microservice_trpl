package cmd

import (
	"gereja-services/config"
	"gereja-services/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mode",
	Run: func(cmd *cobra.Command, args []string) {
		config.Env.Server.Mode = cmd.Flag("appMode").Value.String()
	},
}

func Run() {
	rootCmd.Flags().String("appMode", "dev", "SELECT YOUR APP MODE")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
	app.New(config.Env.Server.Mode)

}
