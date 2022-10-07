package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const defaultShell = "bash"

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
	cmd := exec.Command(defaultShell)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "TTG=true", fmt.Sprintf("TTG_PID=%d", os.Getpid()))

	cmd.Dir = path

	return cmd.Run()
}
