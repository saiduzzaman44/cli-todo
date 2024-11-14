package cmd

import (
	"log"
	"time"

	"github.com/saiduzzaman44/cli-todo/db"
	"github.com/spf13/cobra"
)

func addTodo(cmd *cobra.Command, args []string) {
	name := args[0]

	importance, err := cmd.Flags().GetString("importance")
	if err != nil || (importance != "low" && importance != "mid" && importance != "high") {
		log.Fatalf(`Invalid value for importance. Make sure it's one of ("low", "mid", "high")`)
	}

	dateStr, _ := cmd.Flags().GetString("date")

	var date time.Time

	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Fatalf("Invalid date format. Please use YYYY-MM-DD.")
		}
	}

	_, err = db.DB.Exec(`INSERT INTO todos (name, importance, date) VALUES (?, ?, ?)`, name, importance, date.Format("2006-01-02"))
	if err != nil {
		log.Fatalf("Failed to insert todo: %v", err)
	}

	log.Println("Todo added successfully!")
}

var addCmd = &cobra.Command{
	Use:   "add [name] [-i importance] [-d date]",
	Short: "Add a new task to your to-do list",
	Long: `Add a new task to your to-do list with an optional importance level and due date.

	The "add" command allows you to create a new task with the name as a required argument.
	You can also specify the importance level (-i) as "low", "mid", or "high" and a due date (-d) 
	in YYYY-MM-DD format, both of which are optional.

	Examples:
	cli-todo add "Finish project" -i high -d 2024-12-01
	cli-todo add "Buy groceries"
	`,
	Args: cobra.MinimumNArgs(1),
	Run:  addTodo,
}

func init() {
	addCmd.Flags().StringP("importance", "i", "mid", "Set the importance level (low, mid, high)")
	addCmd.Flags().StringP("date", "d", "", "Set the due date in YYYY-MM-DD format")

	rootCmd.AddCommand(addCmd)
}
