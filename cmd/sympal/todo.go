package main

import (
	"fmt"
	"os"

	"github.com/david-fitzgerald/sympal/internal/db"
	"github.com/david-fitzgerald/sympal/internal/log"
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
		log.Info("todo added", "id", id, "content", content)
		fmt.Printf("Added todo #%d: %s\n", id, content)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.DB.Query("SELECT id, content, status FROM todos ORDER BY id")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to list todos:", err)
			os.Exit(1)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var content, status string
			rows.Scan(&id, &content, &status)

			marker := "[ ]"
			if status == "done" {
				marker = "[x]"
			}
			fmt.Printf("%s #%d: %s\n", marker, id, content)
		}
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a todo as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		result, err := db.DB.Exec("UPDATE todos SET status = 'done' WHERE id = ?", id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to update todo", err)
			os.Exit(1)
		}

		count, _ := result.RowsAffected()
		if count == 0 {
			fmt.Fprintln(os.Stderr, "No todo found with id", id)
			os.Exit(1)
		}

		fmt.Printf("Completed todo #%s\n", id)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		result, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to delete todo", err)
			os.Exit(1)
		}

		count, _ := result.RowsAffected()
		if count == 0 {
			fmt.Fprintln(os.Stderr, "No todo found with id", id)
			os.Exit(1)
		}

		fmt.Printf("Deleted todo #%s\n", id)
	},
}

func init() {
	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(listCmd)
	todoCmd.AddCommand(doneCmd)
	todoCmd.AddCommand(deleteCmd)
}
