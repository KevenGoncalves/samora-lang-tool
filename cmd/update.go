package cmd

import (
	"github.com/KevenGoncalves/samora-lang-tool/internal/handler"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the SML interpreter",
	Long:  `Update the SML interpreter`,
	Run:   handler.RunUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
