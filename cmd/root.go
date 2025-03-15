package cmd

import (
	"os"

	"github.com/gitKashish/golog/internal/config"
	"github.com/gitKashish/golog/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	// Configuration
	cfg *config.Config

	// Command flags
	inputFilePath string
	verbose       bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golog",
	Short: "A simple tool to format your logs.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			if verbose {
				logger.SetLevel(logger.DEBUG)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize configuration
	cfg = config.NewConfig()

	// Add persistent flags that are valid for all commands
	rootCmd.PersistentFlags().BoolVarP(&verbose, "debug", "d", false, "Enable verbose output for debugging")
}
