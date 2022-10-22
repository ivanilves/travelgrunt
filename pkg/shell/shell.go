package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const defaultShell = "bash"

func detectShell() string {
	shell := os.Getenv("SHELL")

	if len(shell) == 0 {
		return defaultShell
	}

	if strings.HasSuffix(shell, "/bash") || strings.HasSuffix(shell, "/zsh") {
		return shell
	}

	if shell == "/bin/true" || shell == "/bin/false" {
		return shell
	}

	return defaultShell
}

// Name returns a called binary name
func Name() string {
	return os.Args[0]
}

// IsRunningInside tells us if we try to run inside a shell spawned by travelgrunt
func IsRunningInside() bool {
	return os.Getenv("TTG") == "true"
}

// Getppid returns numerical process ID of the parent travelgrunt process
func Getppid() int {
	ppid, _ := strconv.Atoi(os.Getenv("TTG_PID"))

	return ppid
}

// Spawn creates a shell in the working directory of selected Terragrunt project
func Spawn(path string) error {
	cmd := exec.Command(detectShell())

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "TTG=true", fmt.Sprintf("TTG_PID=%d", os.Getpid()))

	cmd.Dir = path

	return cmd.Run()
}
