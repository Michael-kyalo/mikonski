package cmd

import (
	"fmt"
	"os"

	"github.com/Michael-kyalo/mikonski/pkg/logging"
	"github.com/Michael-kyalo/mikonski/pkg/session"
	"github.com/spf13/cobra"
)

var (
	// sess is the global session instance shared across commands.
	sess *session.Session

	// rootCmd represents the base command when no subcommands are invoked.
	rootCmd = &cobra.Command{
		Use:   "mikonski",
		Short: "Mikonski is a text-based personal assistant",
		Long: `Mikonski is a command-line personal assistant that provides text-based Q&A, 
manages reminders, and supports exporting session data. Use it to simplify your daily tasks!`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Mikonski! Use --help to see available commands.")
		},
	}
)

// Execute initializes the CLI application.
// This should be called from the main function.
func Execute() {
	defer func() {
		// Save session and reminders on exit
		if err := sess.SaveToFile(); err != nil {
			fmt.Printf("Error saving session: %v\n", err)
		}
		if err := scheduler.SaveToFile(); err != nil {
			fmt.Printf("Error saving reminders: %v\n", err)
		}
		logging.Sync() // Flush logs
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Initialize logging
	logging.InitLogger()
	defer logging.Sync() // Ensure logs are flushed on application exit

	// Initialize the shared session instance
	sess = session.NewSession()

}
