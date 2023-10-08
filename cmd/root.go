package cmd

import (
	"os"

	"github.com/KevenGoncalves/samora-lang-tool/internal/handler"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sml",
	Short: "samora lang tool",
	Run:   handler.RunRoot,
	Long: `
A tool to manage the samora lang
To enter in REPL mode only run sml
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
