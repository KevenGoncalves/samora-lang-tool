package handler

import (
	"fmt"
	"github.com/KevenGoncalves/samora-lang-tool/internal/env"
	"github.com/spf13/cobra"
	"os/exec"
)

func RunHandler(cmd *cobra.Command, args []string) {
	var file string

	if len(args) <= 0 {
		fmt.Println("Pass and sml file")
		return
	}
	file = args[0]

	if file == "" {
		fmt.Println("Pass and sml file")
		return
	}

	compiler, err := env.Config()
	if err != nil {
		fmt.Println("Error locating the compiler")
	}

	osCmd := exec.Command(compiler.CompilerPath, file)
	output, err := osCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
