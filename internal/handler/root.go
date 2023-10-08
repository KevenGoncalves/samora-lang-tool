package handler

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/KevenGoncalves/samora-lang-tool/internal/env"
	"github.com/spf13/cobra"
)

func RunRoot(cmd *cobra.Command, args []string) {
	compiler, err := env.Config()

	if err != nil {
		fmt.Println("Error locating the compiler")
	}

	osCmd := exec.Command(compiler.CompilerPath)
	osCmd.Stdin = os.Stdin
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr

	if err = osCmd.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := osCmd.Wait(); err != nil {
		fmt.Println("Error waiting for SML interpreter:", err)
		return
	}
}
