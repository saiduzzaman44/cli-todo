package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-todo",
	Short: "A simple command line tool to manage your todos",
	Long: `
	The CLI To-Do application helps you manage your tasks directly from the command line.

	With this tool, you can:
	- Add new tasks using the 'add' command, specifying details like name, importance, and due date. This allows you to create organized and prioritized to-dos with ease.
	- Mark tasks as complete with the 'complete' command. This helps keep track of your progress and ensures you know which tasks are done and which are still pending.
	- List all tasks with the 'list' command to view your to-dos at a glance. The list displays task names, importance levels, completion status, and dates, giving you a quick overview of all your tasks.

	This CLI app stores your tasks in a local SQLite database, so all data is saved locally on your machine. Whether you're a productivity enthusiast or simply want a quick way to manage your to-dos, this application makes task management fast, efficient, and convenient.

	Commands:
	- add       Adds a new task to your to-do list
	- complete  Marks a specified task as completed
	- list      Displays all tasks in your to-do list with details

	Get started by adding your first task, and enjoy seamless to-do management from your terminal!
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
