package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Parse a diff",
	Long:  `Parse a git diff to an output format`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.Attach(); err != nil {
			log.Fatal(err)
		}

		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
