package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/auth"
	"github.com/david-fitzgerald/sympal/internal/config"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Begin auth flow.",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := auth.Authenticate(config.Current.Google.ClientID, config.Current.Google.ClientSecret)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed auth flow", err)
			os.Exit(1)
		}
		fmt.Println("Completed auth flow")

	},
}
