package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/db"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage todos",
}

var addCmd = &cobra.Command{
	Use:   "add [content]",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := args[0]

		result, err := db.DB.Exec("INSERT INTO todos (content) VALUES (?)", content)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to add todo:", err)
			os.Exit(1)
		}

		id, _ := result.LastInsertId()
		fmt.Printf("Added todo #%d: %s\n", id, content)
	},
}

func init() {
	todoCmd.AddCommand(addCmd)
}
