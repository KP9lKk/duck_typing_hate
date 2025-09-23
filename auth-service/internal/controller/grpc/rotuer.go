package grpc

import (
	v1 "duck_typing_hate/auth-service/internal/controller/grpc/v1"
	"duck_typing_hate/auth-service/internal/usecase"

	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter(app *pbgrpc.Server, n usecase.NonceUseCase) {
	{
		v1.NewNonceRoutes(app, n)
	}
	reflection.Register(app)
}
