package cmd

import (
	"github.com/karchx/xml-go/pkg/filters"
	"github.com/karchx/xml-go/pkg/utils"
	"github.com/spf13/cobra"
)

var filterXmlCmd = &cobra.Command{
	Use:     "filterxml [directory] [file csv]",
	Aliases: []string{"fxml"},
	Short:   "Filter xml file",
	Run: func(cmd *cobra.Command, args []string) {
		var filtersCsv []utils.CSV
		if len(args) > 1 {
			filtersCsv = filters.GetFiltersCsv(args[1])
		} else {
			filtersCsv = []utils.CSV{}
		}
		filters.GetFiles(args[0], filtersCsv)
	},
}

func init() {
	rootCmd.AddCommand(filterXmlCmd)
}
