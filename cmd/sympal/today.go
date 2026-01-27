package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/calendar"
	"github.com/david-fitzgerald/sympal/internal/db"

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
		fmt.Println("\nToday's Calendar")

		for _, item := range todayCalResults {
			fmt.Printf("%s - %s\n", item.StartTime, item.Title)
		}
		// ToDo fetching
		fmt.Println("\nToDos:")
		rows, err := db.DB.Query("SELECT id, content FROM todos WHERE status != 'done' ORDER BY id")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to retrieve todos", err)
			os.Exit(1)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var content string
			rows.Scan(&id, &content)

			fmt.Printf("[ ] #%d: %s\n", id, content)
		}

	},
}
