package cmd

import (
	"fmt"
	"github.com/nathanburkett/diff-chart/algorithm"
	"github.com/nathanburkett/diff-chart/input"
	"github.com/nathanburkett/diff-chart/output"
	"github.com/nathanburkett/diff-chart/run"
	"github.com/nathanburkett/diff-chart/transform"
	"github.com/spf13/cobra"
	"os"
)

var (
	app          = run.NewApp()
	rootCmd      = &cobra.Command{
		Use:   "diff-chart",
		Short: "DiffChart is an informative and interpretive statistic generator for Git diffs",
		Long: `A fast and flexible diff statistic generator built with
	love by Nathan Burkett in Go.`,
		TraverseChildren: true,
	}
)

func init() {
	diffCmd.Flags().StringVarP(&app.Flags.DownstreamRef, "downstream", "d", "HEAD^1", "Downstream reference - 'from' value")
	diffCmd.Flags().StringVarP(&app.Flags.UpstreamRef, "upstream", "u", "HEAD", "Upstream reference - 'to' value")
	diffCmd.Flags().StringVarP(&app.Flags.InputType, "input", "i", input.TypeGit, "Input format")
	diffCmd.Flags().StringVarP(&app.Flags.OutputType, "output", "o", output.TypeMarkdownCli, "Output format")
	diffCmd.Flags().StringVarP(&app.Flags.ReducerType, "reduce", "r", transform.TypeDirectoryReducer, "Map reduce strategy")
	diffCmd.Flags().StringVarP(&app.Flags.SortType, "sort", "s", algorithm.TypeTotalDeltaDesc, "Sort strategy")

	rootCmd.AddCommand(diffCmd)
}

// Execute root command execution process
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
