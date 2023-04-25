package cmd

import (
	"github.com/karchx/xml-go/pkg/filters"
	"github.com/spf13/cobra"
)

var filterXmlCmd = &cobra.Command{
	Use:     "filterxml",
	Aliases: []string{"fxml"},
	Short:   "Filter xml file",
	Run: func(cmd *cobra.Command, args []string) {
		filters.GetFiles(args[0])
		if len(args) > 1 {
			filters.GetFiltersCsv(args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(filterXmlCmd)
}
