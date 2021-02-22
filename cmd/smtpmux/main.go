package main

import (
	"os"
	"runtime"
	"sync"

	"github.com/burungbangkai/fakesmtp/internal/adapter/cache"
	"github.com/burungbangkai/fakesmtp/internal/adapter/config"
	"github.com/burungbangkai/fakesmtp/internal/adapter/event"
	"github.com/burungbangkai/fakesmtp/internal/adapter/grpc"
	"github.com/burungbangkai/fakesmtp/internal/adapter/smtpmux"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
)

func main() {
	var (
		cfg      model.Config
		cfgLoadr port.LoadConfig = config.EnvConfigLoader
		hostname string
		wg       *sync.WaitGroup
	)
	if err := cfgLoadr(&cfg); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	sc, scStop, err := event.NewNATSConnection(cfg.NatsURI, cfg.NatsClusterID, hostname)
	if err != nil {
		panic(err)
	}
	var (
		uInboxCfgGetter port.GetUserInboxConfigByKey = cache.NewUserInboxConfigSychedMapCache()
		rawEmailSender  port.SendRawEmail            = event.NewNATSRawEmailReceivedEventSender(sc, grpc.Serializer)
		loginer                                      = usecase.NewEmailSessionProcessor(uInboxCfgGetter)
		emailReceiver                                = usecase.NewEmailReceiver(cfg.MaxEmailSize, rawEmailSender)
	)
	smtp, smtpStop := smtpmux.New(cfg.SMTPPort, hostname, emailReceiver, loginer)
	wg = usecase.GracefulStop(smtpStop, scStop)
	smtp()
	wg.Wait()
}
