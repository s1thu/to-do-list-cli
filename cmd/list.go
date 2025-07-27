package cmd

import (
	"fmt"
	"to_do_cli/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todo.LoadTasks("tasks.json")
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			return
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
