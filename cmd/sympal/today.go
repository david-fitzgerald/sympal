package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/calendar"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Run today command",
	Run: func(cmd *cobra.Command, args []string) {
		todayCalResults, err := calendar.GetTodayEvents()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to import todays calendar", err)
			os.Exit(1)
		}
		fmt.Println("Imported today's calendar events:")

		for _, item := range todayCalResults {
			fmt.Printf("%s - %s\n", item.StartTime, item.Title)
		}
	},
}
