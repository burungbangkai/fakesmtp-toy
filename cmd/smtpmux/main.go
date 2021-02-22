package main

import (
	"os"
	"runtime"
	"sync"

	"github.com/burungbangkai/fakesmtp/internal/adapter/config"
	"github.com/burungbangkai/fakesmtp/internal/adapter/smtpmux"
	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
)

func main() {
	var (
		cfg      model.Config
		cfgLoadr port.LoadConfig
		hostname string
		wg       *sync.WaitGroup
	)
	cfgLoadr = config.EnvConfigLoader
	if err := cfgLoadr(&cfg); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	var (
		uInboxCfgGetter port.GetUserInboxConfigByKey
		rawEmailSender  port.SendRawEmail
		loginer         = usecase.NewEmailSessionProcessor(uInboxCfgGetter)
		emailReceiver   = usecase.NewEmailReceiver(cfg.MaxEmailSize, rawEmailSender)
	)
	smtp, smtpStop := smtpmux.New(cfg.SMTPPort, hostname, emailReceiver, loginer)
	wg = usecase.GracefulStop(smtpStop)
	smtp()
	wg.Wait()
}
