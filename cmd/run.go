package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an samora lang file",
	Long:  `Run and samora lang file`,
	Run:   handler,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func handler(cmd *cobra.Command, args []string) {
	var file string

	if len(args) <= 0 {
		fmt.Println("Pass and sml file")
		return
	}
	file = args[0]

	if file == "" || !strings.Contains(file, ".sml") {
		fmt.Println("Pass and sml file")
		return
	}

	osCmd := exec.Command("./tmp/compiler", file)
	output, err := osCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
