package cmd

import (
	"fmt"

	"github.com/gitKashish/golog/internal/core/formatter"
	"github.com/gitKashish/golog/internal/core/parser"
	"github.com/gitKashish/golog/pkg/fileutil"
	"github.com/gitKashish/golog/pkg/logger"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show formatted logs from a file",
	Long:  `Read logs from a file and display them formatted according to the template.`,
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

		// Format and print each line
		for _, line := range sourceLines {
			fmt.Println(f.FormatLog(line))
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
