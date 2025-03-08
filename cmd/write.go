package cmd

import (
	"fmt"
	"os"

	"github.com/gitkashish/golog/internal/core"
	"github.com/gitkashish/golog/internal/helpers"
	"github.com/spf13/cobra"
)

var outputFilePath string
var showOutput bool

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		sourceLines := helpers.ReadFileToArray(inputFilePath)
		template, err := core.GetTemplateFromFile()
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		logEntries := []string{}
		for _, line := range sourceLines {
			formattedLog := template.Parse(line)
			logEntries = append(logEntries, formattedLog)
			if showOutput {
				fmt.Print(formattedLog)
			}
		}

		helpers.WriteArrayToFile(logEntries, outputFilePath)
		fmt.Printf("Logs written to %v", outputFilePath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	writeCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Path to input file (required)")
	writeCmd.MarkFlagFilename("input")
	writeCmd.MarkFlagRequired("input")

	writeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Path to output file (required)")
	writeCmd.MarkFlagRequired("output")
	writeCmd.MarkFlagFilename("output")

	writeCmd.Flags().BoolVarP(&showOutput, "show", "s", false, "Show output on console")
}
