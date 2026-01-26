package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/config"
	"github.com/david-fitzgerald/sympal/internal/db"
	"github.com/david-fitzgerald/sympal/internal/log"
	"github.com/spf13/cobra"
)

func main() {
	if err := db.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialise database", err)
		os.Exit(1)
	}

	if err := config.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load config:", err)
		os.Exit(1)
	}

	if err := log.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialise logging:", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "sympal",
		Short: "Privacy layer for AI-assisted workflows",
	}

	rootCmd.AddCommand(todoCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(todayCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
