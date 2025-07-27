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

		tasks, err := todo.LoadTasks("tasks.json")
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			return
		}

		m := todo.NewManager()
		for _, t := range tasks {
			m.AddTask(t.Name)
			if t.Completed {
				_ = m.CompletedTask(t.ID)
			}
		}

		task := m.AddTask(taskName)
		err = todo.SaveTasks("tasks.json", m.GetTasks())
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
