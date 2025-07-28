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
		m := todo.NewManager()
		err := m.LoadFromFile("tasks.json")
		if err != nil {
			cmd.Printf("Error loading tasks: %v\n", err)
			return
		}

		for _, t := range m.GetTasks() {
			fmt.Println(t.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
