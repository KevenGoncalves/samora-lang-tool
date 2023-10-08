package cmd

import (
	"github.com/KevenGoncalves/samora-lang-tool/internal/handler"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Run:   handler.RunHandler,
	Short: "Run an samora lang file",
	Long: `
Run and samora lang file
Example: sml run <file.sml>
  `,
}

func init() {
	rootCmd.AddCommand(runCmd)
}
