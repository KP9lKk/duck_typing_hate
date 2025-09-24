package grpc

import (
	v1 "duck_typing_hate/auth-service/internal/controller/grpc/v1"
	"duck_typing_hate/auth-service/internal/usecase"
	"duck_typing_hate/shared/pkg/logger"

	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter(app *pbgrpc.Server, n usecase.NonceUseCase, l logger.Logger) {
	{
		v1.NewNonceRoutes(app, n, l)
	}
	reflection.Register(app)
}
