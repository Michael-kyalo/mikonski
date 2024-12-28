package cmd

import (
	"fmt"
	"github.com/Michael-kyalo/mikonski/pkg/reminders"
	"github.com/spf13/cobra"
	"time"
)

var (
	reminderDescription string
	reminderTime        string
)

// Scheduler manages reminders during the session.
var scheduler = reminders.NewScheduler()

// reminderCmd represents the parent command for reminder-related operations.
var reminderCmd = &cobra.Command{
	Use:   "reminder",
	Short: "Manage reminders in Mikonski",
	Long: `The "reminder" command lets you manage reminders. 
Subcommands include "set", "list", and "clear".`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand: set, list, clear")
	},
}

// Subcommand: `set` - Adds a new reminder.
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a new reminder",
	Long: `The "set" subcommand allows you to create a new reminder.
You need to provide both a description and a time.`,
	Run: func(cmd *cobra.Command, args []string) {
		if reminderDescription == "" || reminderTime == "" {
			fmt.Println("Please provide both --description and --time flags.")
			return
		}

		// Parse the reminder time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", reminderTime)
		if err != nil {
			fmt.Printf("Invalid time format. Use 'YYYY-MM-DD HH:MM:SS'. Error: %v\n", err)
			return
		}

		// Add the reminder to the scheduler
		err = scheduler.Set(reminderDescription, parsedTime)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Reminder set successfully!")
	},
}

// Subcommand: `list` - Displays all active reminders.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all reminders",
	Run: func(cmd *cobra.Command, args []string) {
		reminders := scheduler.List()
		if len(reminders) == 0 {
			fmt.Println("No reminders set.")
			return
		}

		// Print all reminders
		for i, reminder := range reminders {
			fmt.Printf("%d. %s at %s\n", i+1, reminder.Description, reminder.Time.Format("2006-01-02 15:04:05"))
		}
	},
}

// Subcommand: `clear` - Removes all reminders.
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all reminders",
	Run: func(cmd *cobra.Command, args []string) {
		scheduler.Clear()
		fmt.Println("All reminders cleared!")
	},
}

func init() {
	// Set subcommand flags
	setCmd.Flags().StringVarP(&reminderDescription, "description", "d", "", "Description of the reminder")
	setCmd.Flags().StringVarP(&reminderTime, "time", "t", "", "Time of the reminder in 'YYYY-MM-DD HH:MM:SS' format")
	// Register subcommands under `reminder`
	reminderCmd.AddCommand(setCmd)
	reminderCmd.AddCommand(listCmd)
	reminderCmd.AddCommand(clearCmd)
	// Register the `reminder` command as a subcommand of the root
	rootCmd.AddCommand(reminderCmd)
}
