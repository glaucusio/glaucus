// +build windows

package command

import (
	"os"
)

var defaultSignals = []os.Signal{
	os.Interrupt,
	os.Kill,
}
