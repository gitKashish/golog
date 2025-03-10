/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strconv"

	server "github.com/gitKashish/golog/internal"
	"github.com/spf13/cobra"
)

// Port number to serve the web app on
var port int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving the golog web app on specified port (default: 2600, overriden by PORT environment variable)",
	Long: `Start an HTTP server on the specified port to serve the golog webapp.
	Default port is 2600 but you can specify a different port using the -p or --port flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			port = 2600
		}
		server := server.NewServer(port)
		return server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", 2600, "Port to start the server on")
}
