package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "view recent log entries",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		logPath := filepath.Join(homeDir, ".sympal", "sympal.log")

		data, err := os.ReadFile(logPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read log", err)
			os.Exit(1)
		}

		lines := strings.Split(string(data), "\n")

		start := len(lines) - 20
		if start < 0 {
			start = 0
		}

		for _, line := range lines[start:] {
			if line != "" {
				fmt.Println(line)
			}
		}
	},
}
