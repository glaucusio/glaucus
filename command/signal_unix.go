// +build !windows

package command

import (
	"os"
	"syscall"
)

var defaultSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}
