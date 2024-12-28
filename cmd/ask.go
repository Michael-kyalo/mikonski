package cmd

import (
	"fmt"

	"github.com/Michael-kyalo/mikonski/pkg/ai"
	"github.com/Michael-kyalo/mikonski/pkg/config"
	"github.com/Michael-kyalo/mikonski/pkg/session"
	"github.com/spf13/cobra"
)

var (
	question string
	apiKey   string
	model    string
)

// Intialize session management
var sess = session.NewSession()

// askCmd represents the `ask` command for querying Mikonski.
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question and get an answer from Mikonski",
	Run: func(cmd *cobra.Command, args []string) {
		// Load base configuration
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		// Apply CLI overrides
		overrides := map[string]string{
			"APIKey": apiKey,
			"Model":  model,
		}
		cfg = config.ApplyOverrides(cfg, overrides)

		if question == "" {
			fmt.Println("Please provide a question using the --question flag.")
			return
		}

		client := ai.NewOpenAIClient(cfg.APIKey, cfg.Model)
		context := sess.GetContext()

		// Append context if available
		if context != "" {
			question = context + " " + question
		}

		response, err := client.Ask(question)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Update session context
		sess.AddContext(question, response)

		fmt.Printf("Mikonski: %s\n", response)
	},
}

func init() {
	askCmd.Flags().StringVarP(&question, "question", "q", "", "The question to ask")
	askCmd.Flags().StringVar(&apiKey, "apikey", "", "OpenAI API Key (overrides configuration)")
	askCmd.Flags().StringVar(&model, "model", "", "OpenAI model to use (overrides configuration)")

	rootCmd.AddCommand(askCmd)
}
