/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/saiduzzaman44/cli-todo/db"
	"github.com/spf13/cobra"
)

func listTodos(cmd *cobra.Command, args []string) {
	query := "SELECT * FROM todos"

	conditions := []string{"completed = 0"}

	dateStr, _ := cmd.Flags().GetString("date")
	if dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatal("Error parsing date. Make sure date is in YYYY-MM-DD format")
		}

		conditions = append(conditions, fmt.Sprintf("date = %v", date))
	}

	importance, _ := cmd.Flags().GetString("importance")
	if importance != "" {
		if importance != "low" && importance != "mid" && importance != "high" {
			log.Fatalf(`Invalid value for importance. Make sure it's one of ("low", "mid", "high")`)
		}
		conditions = append(conditions, fmt.Sprintf("importance = %v", importance))
	}

	query += " WHERE " + strings.Join(conditions, " AND ")

	limit, err := cmd.Flags().GetInt("number")
	if err != nil {
		log.Fatalf("Invalid value for number flag. Make sure it's an integer")
	}

	query += fmt.Sprintf(" LIMIT %d", limit)

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatalf("Failed to fetch todos: %v", err)
	}
	defer rows.Close()

	lowColor := color.New(color.FgGreen).SprintFunc()
	midColor := color.New(color.FgYellow).SprintFunc()
	highColor := color.New(color.FgRed).SprintFunc()

	tbl := table.New("ID", "Name", "Importance", "Completed", "Date")

	for rows.Next() {
		var id int
		var name, importance, date string
		var completed bool
		err := rows.Scan(&id, &name, &importance, &completed, &date)
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}

		var coloredImportance string
		switch importance {
		case "low":
			coloredImportance = lowColor(importance)
		case "mid":
			coloredImportance = midColor(importance)
		case "high":
			coloredImportance = highColor(importance)
		default:
			coloredImportance = importance
		}

		tbl.AddRow(id, name, coloredImportance, completed, date)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	tbl.Print()
}

var listCmd = &cobra.Command{
	Use:   "list [-d date] [-i importance] [-n number]",
	Short: "List all your to-dos with optional filters",
	Long: `The "list" command displays all the tasks in your to-do list.

	You can filter the tasks based on the following optional flags:
	- -d [date]   Show tasks due on a specific date in YYYY-MM-DD format.
	- -i [importance]   Filter tasks based on importance (low, mid, high).
	- -n [number]   Limit the number of tasks shown (default is to show all tasks).

	If no flags are provided, the command will display all tasks in the to-do list.

	Examples:
	cli-todo list
	cli-todo list -i high
	cli-todo list -d 2024-12-01
	cli-todo list -n 5 -i low -d 2024-12-01`,
	Args: cobra.NoArgs,
	Run:  listTodos,
}

func init() {

	listCmd.Flags().StringP("importance", "i", "", "Filter todos by the importance level (low, mid, high)")
	listCmd.Flags().StringP("date", "d", "", "Filter todos by date (YYYY-MM-DD format)")
	listCmd.Flags().IntP("number", "n", 5, "number of todos to list")

	rootCmd.AddCommand(listCmd)
}
