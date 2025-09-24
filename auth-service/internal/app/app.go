package app

import (
	"duck_typing_hate/auth-service/config"
	"duck_typing_hate/auth-service/internal/controller/grpc"
	"duck_typing_hate/auth-service/internal/repo/persistent"
	"duck_typing_hate/auth-service/internal/usecase/nonce"
	"duck_typing_hate/shared/pkg/grpcserver"
	"duck_typing_hate/shared/pkg/logger"
	"duck_typing_hate/shared/pkg/reddis"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	rdb := reddis.New(cfg.RDB.Url, cfg.RDB.Password, cfg.RDB.Db)
	defer rdb.Close()
	nonceUseCase := nonce.New(
		persistent.New(rdb),
	)

	logger := logger.New("error")
	defer logger.Logger.Sync()

	grpcServer := grpcserver.New(cfg.GRPC.Port)
	grpc.NewRouter(grpcServer.App, nonceUseCase, *logger)

	grpcServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		logger.Logger.Info(fmt.Sprintf("app - Run - signal: %s", s.String()))
	case err := <-grpcServer.Notify():
		logger.Logger.Error(fmt.Sprintf("app - Run - grpcServer.Notify: %s", err))
	}

	err := grpcServer.ShutDown()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("app - Run - grpcServer.ShutDown: %s", err))
	}

}
