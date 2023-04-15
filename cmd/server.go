package cmd

import (
	"github.com/Redish101/fiber-starter/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}