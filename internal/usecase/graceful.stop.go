package usecase

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/burungbangkai/fakesmtp/internal/port"
)

func GracefulStop(gcs ...port.StopGracefully) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(len(gcs))
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stopChan
		for _, gcfn := range gcs {
			gcfn()
			wg.Done()
		}
	}()
	return wg
}
