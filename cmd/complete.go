package cmd

import (
	"fmt"
	"log"

	"github.com/saiduzzaman44/cli-todo/db"
	"github.com/spf13/cobra"
)

func completeTodo(cmd *cobra.Command, args []string) {
	id := args[0]

	result, err := db.DB.Exec("UPDATE todos SET completed = 1 WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to mark todo as complete: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No todo found with the specified ID.")
	} else {
		fmt.Println("Todo marked as complete!")
	}
}

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Marks a todo as complete",
	Args:  cobra.ExactArgs(1),
	Run:   completeTodo,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
