package app

import (
	"duck_typing_hate/auth-service/config"
	"duck_typing_hate/auth-service/internal/controller/grpc"
	"duck_typing_hate/auth-service/internal/repo/persistent"
	"duck_typing_hate/auth-service/internal/usecase/nonce"
	"duck_typing_hate/auth-service/pkg/grpcserver"
	"duck_typing_hate/auth-service/pkg/reddis"
	"fmt"
	"os"
)

func Run(cfg *config.Config) {

	rdb := reddis.New(cfg.RDB.Url, cfg.RDB.Password, cfg.RDB.Db)
	defer rdb.Close()
	nonceUseCase := nonce.New(
		persistent.New(rdb),
	)

	grpcServer := grpcserver.New(cfg.GRPC.Port)
	grpc.NewRouter(grpcServer.App, nonceUseCase)

	grpcServer.Start()

	interrupt := make(chan os.Signal, 1)

	select {
	case s := <-interrupt:
		fmt.Println("app - Run - signal: %w", s.String())
	case err := <-grpcServer.Notify():
		fmt.Println("app - Run - grpcServer.Notify: %w", err)
	}
}
