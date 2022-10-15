package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const defaultShell = "bash"
const defaultLines = 20

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

func detectLines() int {
	lines := os.Getenv("LINES")

	if len(lines) == 0 {
		return defaultLines
	}

	lnum, err := strconv.Atoi(lines)

	if err != nil {
		return defaultLines
	}

	return lnum
}

func isMocked() bool {
	shell := detectShell()

	return shell == "/bin/true" || shell == "/bin/false"
}

// Name returns a called binary name
func Name() string {
	return os.Args[0]
}

// IsRunningInside tells us if we try to run inside a shell spawned by ttg
func IsRunningInside() bool {
	return os.Getenv("TTG") == "true"
}

// Getppid returns numerical process ID of the parent ttg process
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

// PrintAndExit prints a string passed and exits after
func PrintAndExit(s string) {
	fmt.Printf("%s\n", s)

	if !isMocked() {
		os.Exit(0)
	}
}
