package log

import (
	"fmt"

	"github.com/pterm/pterm"
)

var logger pterm.Logger = pterm.DefaultLogger

func Infof(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	logger.Info(s)
}
