package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/db"
	"github.com/spf13/cobra"
)

func main() {
	if err := db.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialise database", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "sympal",
		Short: "Privacy layer for AI-assisted workflows",
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
