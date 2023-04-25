package cmd

import (
	"github.com/karchx/xml-go/pkg/filters"
	"github.com/spf13/cobra"
)

var filterXmlCmd = &cobra.Command{
	Use:     "filterxml",
	Aliases: []string{"fxml"},
	Short:   "Filter xml file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filters.GetFiles(args[0])
	},
}

func init() {
	rootCmd.AddCommand(filterXmlCmd)
}
