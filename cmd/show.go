package cmd

import (
	"fmt"

	"github.com/gitkashish/golog/internal/core"
	"github.com/gitkashish/golog/internal/helpers"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		sourceLines := helpers.ReadFileToArray(inputFilePath)
		for _, line := range sourceLines {
			logEntry := core.ParseLogLine(line)
			fmt.Print(core.FormatLogEntry(logEntry))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Path to input file (required)")
	showCmd.MarkFlagFilename("input")
	showCmd.MarkFlagRequired("input")
}
