package kubectl

import (
	"log"
	"os"
)

func logf(msg string, args ...interface{}) {
	ppid := os.Getppid()
	pid := os.Getpid()
	log.Printf("[DEBUG] kubectl-provider(pid=%d,ppid=%d): "+msg, append([]interface{}{pid, ppid}, args...)...)
}
