package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple command line todo application",
	Long: `A simple command line todo application to manage your tasks.
	Manage your daily tasks from the terminal!

Examples:
  todo add "Buy milk"
  todo list
  todo complete 1
  todo delete 1
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Display help if no subcommand is provided
	},
}

func Execute() error {
	return rootCmd.Execute()
}
