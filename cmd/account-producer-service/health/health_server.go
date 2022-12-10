package health

import (
	"account-producer-service/internal/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

func NewHealthServer() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killsignal := <-interrupt
	switch killsignal {
	case os.Interrupt:
		utils.Logger.Info("got sigint signal... interrupt")
	case syscall.SIGTERM:
		utils.Logger.Info("got sigterm signal... interrupt")
	}

}
