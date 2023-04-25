package cmd

import (
	"os"

	log "github.com/gothew/l-og"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xml-go",
	Short: "xml-go is a very fast static xml files",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.SetOutput(os.Stderr)
		log.Error(err)
		os.Exit(1)
	}
}
