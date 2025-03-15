/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gitKashish/golog/internal/server"
	"github.com/gitKashish/golog/pkg/logger"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving the golog web app on specified port",
	Long: `Start an HTTP server on the specified port to serve the golog webapp.
Default port is 2600 but you can specify a different port using the -p or --port flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Override port from command line if specified
		if cmd.Flags().Changed("port") {
			cfg.Server.Port = port
		}

		// Create and configure server
		srv := server.NewServer(cfg)
		srv.SetupRoutes()

		// Start server in a goroutine
		go func() {
			logger.Info("Starting server on port %d", cfg.Server.Port)
			if err := srv.Start(); err != nil {
				logger.Fatal("Server error: %v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shut down the server
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		logger.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			logger.Error("Server forced to shutdown: %v", err)
			return err
		}

		logger.Info("Server gracefully stopped")
		return nil
	},
}

// Port number to serve the web app on
var port int

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 2600, "Port to start the server on")
}
