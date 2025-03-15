package cmd

import (
	"fmt"

	"github.com/gitKashish/golog/internal/core/formatter"
	"github.com/gitKashish/golog/internal/core/parser"
	"github.com/gitKashish/golog/pkg/fileutil"
	"github.com/gitKashish/golog/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	outputFilePath string
	showOutput     bool
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write formatted logs to a file",
	Long:  `Read logs from a file, format them according to the template, and write them to another file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create file utility
		fileUtil := fileutil.NewFileUtil()

		// Read source lines from file
		sourceLines, err := fileUtil.ReadLines(inputFilePath)
		if err != nil {
			logger.Error("Error reading file: %v", err)
			return err
		}

		// Create parser and formatter
		p := parser.NewTemplateParser()
		err = p.LoadTemplate(cfg.Template.TemplatePath)
		if err != nil {
			logger.Error("Error loading template: %v", err)
			return err
		}

		f := formatter.NewTemplateFormatter(p)

		// Format logs
		formattedLogs := f.FormatLogs(sourceLines)

		// Show output if requested
		if showOutput {
			for _, log := range formattedLogs {
				fmt.Println(log)
			}
		}

		// Write to output file
		err = fileUtil.WriteLines(formattedLogs, outputFilePath)
		if err != nil {
			logger.Error("Error writing to file: %v", err)
			return err
		}

		logger.Info("Logs written to %s", outputFilePath)
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
