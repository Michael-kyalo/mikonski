package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var logFile = "mikonski.log" // Path to the log file

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View the latest logs from Mikonski",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(logFile)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			return
		}
		defer file.Close()

		fmt.Println("Latest logs from Mikonski:")
		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			fmt.Printf("Error reading log file: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
