package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type InterruptChan chan os.Signal

func WaitInterrupt() InterruptChan {
	interruptChannel := make(InterruptChan, 1)
	signal.Notify(interruptChannel, []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
	}...)

	return interruptChannel
}
