package main

import (
	"runtime"
	"sync"

	"github.com/burungbangkai/fakesmtp/internal/adapter/config"
	"github.com/burungbangkai/fakesmtp/internal/adapter/grpc"
	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
)

func main() {
	var (
		cfg      model.Config
		cfgLoadr port.LoadConfig = config.EnvConfigLoader
		wg       *sync.WaitGroup
	)
	if err := cfgLoadr(&cfg); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		gui          usecase.GetUserInboxs
		aui          usecase.AddUserInbox
		dui          usecase.DeleteUserInbox
		uui          usecase.UpdateUserInbox
		rpc, rpcStop = grpc.NewServer(cfg.GRPCPort, gui, aui, dui, uui)
	)
	wg = usecase.GracefulStop(rpcStop)
	rpc()
	wg.Wait()
}
