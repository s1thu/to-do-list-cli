package cmd

import (
	"fmt"
	"strconv"
	"to_do_cli/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task from the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Printf("Invalid task ID: %v\n", err)
			return
		}

		manager := todo.NewManager()
		err = manager.LoadFromFile("tasks.json")
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		err = manager.DeleteTask(taskID)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}

		if err := todo.SaveTasks("tasks.json", manager.GetTasks()); err != nil {
			cmd.Printf("Error saving tasks: %v\n", err)
			return
		}

		fmt.Println("✅ Task deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
