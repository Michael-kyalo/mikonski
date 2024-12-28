package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when no subcommands are invoked.
var rootCmd = &cobra.Command{
	Use:   "mikonski",
	Short: "Mikonski is a text-based personal assistant",
	Long: `Mikonski is a command-line tool that acts as a personal assistant,
capable of answering questions, managing reminders, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Mikonski! Use --help to see available commands.")
	},
}

// Execute initializes the CLI application.
// It should be called from main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
