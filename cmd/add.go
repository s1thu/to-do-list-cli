package cmd

import (
	"fmt"
	"to_do_cli/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]

		manager := todo.NewManager()
		err := manager.LoadFromFile("tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		task := manager.AddTask(taskName)

		err = todo.SaveTasks("tasks.json", manager.GetTasks())
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Println("✅ Added task:", task.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
