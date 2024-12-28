package cmd

import (
	"fmt"

	"github.com/Michael-kyalo/mikonski/pkg/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	historyFile   string
	remindersFile string
)

// exportCmd represents the parent command for export operations.
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export session data to files",
	Long:  "Export session history and reminders to JSON files for future reference.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand: history or reminders")
	},
}

// Subcommand: `history` - Export session history.
var exportHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Export session history",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()

		if historyFile == "" {
			historyFile = "session_history.json"
		}

		err := sess.ExportHistory(historyFile)
		if err != nil {
			logger.Error("failed to export session history", zap.Error(err))
			fmt.Printf("Error exporting session history: %v\n", err)
			return
		}

		logger.Info("session history exported", zap.String("file", historyFile))
		fmt.Printf("Session history exported to %s\n", historyFile)
	},
}

// Subcommand: `reminders` - Export reminders.
var exportRemindersCmd = &cobra.Command{
	Use:   "reminders",
	Short: "Export reminders",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()

		if remindersFile == "" {
			remindersFile = "reminders.json"
		}

		err := scheduler.ExportReminders(remindersFile)
		if err != nil {
			logger.Error("failed to export reminders", zap.Error(err))
			fmt.Printf("Error exporting reminders: %v\n", err)
			return
		}

		logger.Info("reminders exported", zap.String("file", remindersFile))
		fmt.Printf("Reminders exported to %s\n", remindersFile)
	},
}

func init() {
	exportHistoryCmd.Flags().StringVar(&historyFile, "file", "", "File to save session history (default: session_history.json)")
	exportRemindersCmd.Flags().StringVar(&remindersFile, "file", "", "File to save reminders (default: reminders.json)")

	exportCmd.AddCommand(exportHistoryCmd)
	exportCmd.AddCommand(exportRemindersCmd)
	rootCmd.AddCommand(exportCmd)
}
