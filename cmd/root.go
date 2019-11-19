package cmd

import (
	"fmt"
	"github.com/nathanburkett/diff_table/algorithm"
	"github.com/nathanburkett/diff_table/input"
	"github.com/nathanburkett/diff_table/output"
	"github.com/nathanburkett/diff_table/run"
	"github.com/nathanburkett/diff_table/transform"
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
	diffCmd.Flags().StringVarP(&app.Args.DownstreamRef, "downstream", "d", "HEAD^1", "Downstream reference - 'from' value")
	diffCmd.Flags().StringVarP(&app.Args.UpstreamRef, "upstream", "u", "HEAD", "Upstream reference - 'to' value")
	diffCmd.Flags().StringVarP(&app.Args.InputType, "input", "i", input.TypeGit, "Input format")
	diffCmd.Flags().StringVarP(&app.Args.OutputType, "output", "o", output.TypeMarkdownCli, "Output format")
	diffCmd.Flags().StringVarP(&app.Args.ReducerType, "reduce", "r", transform.TypeDirectoryReducer, "Map reduce strategy")
	diffCmd.Flags().StringVarP(&app.Args.SortType, "sort", "s", algorithm.TypeTotalDeltaDesc, "Sort strategy")

	rootCmd.AddCommand(diffCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
