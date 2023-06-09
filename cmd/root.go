package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/Redish101/fiber-starter/config"
)

var rootCmd = &cobra.Command{
	Use: "fiber-starter",
	Short: "Redish101's Fiber Starter",
	Long: "A template for Fiber App.",
	Version: config.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.AppName)
		fmt.Println("==========================")
	},
}

func Execute() {
	rootCmd.AddCommand(serverCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}