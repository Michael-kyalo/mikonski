package cmd

import (
	"fmt"

	"github.com/Michael-kyalo/mikonski/pkg/ai"
	"github.com/Michael-kyalo/mikonski/pkg/session"
	"github.com/spf13/cobra"
)

var question string

// Intialize session management
var sess = session.NewSession()

// askCmd represents the `ask` command for querying Mikonski.
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question and get an answer from Mikonski",
	Long: `The "ask" command allows you to query Mikonski with a question.
For example:
  mikonski ask --question "What is the capital of France?"`,
	Run: func(cmd *cobra.Command, args []string) {
		if question == "" {
			fmt.Println("Please provide a question using the --question flag.")
			return
		}

		// Initialize the AI client
		client := ai.OpenAIClient{}
		context := sess.GetContext()

		// Append context if available
		if context != "" {
			question = context + " " + question
		}

		// Get the response from the AI model
		response, err := client.Ask(question)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Update the session context with the response
		sess.AddContext(question, response)

		// Print the response
		fmt.Printf("Mikonski: %s\n", response)
	},
}

func init() {
	// Add a flag to capture the user's question
	askCmd.Flags().StringVarP(&question, "question", "q", "", "The question to ask")
	// Register the `ask` command as a subcommand of the root
	rootCmd.AddCommand(askCmd)
}
