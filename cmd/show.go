package cmd

import (
	"fmt"
	"os"

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
		template, err := core.GetTemplateFromFile()
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		for _, line := range sourceLines {
			fmt.Print(template.Parse(line))
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
