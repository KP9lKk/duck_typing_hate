package v1

import (
	v1 "duck_typing_hate/auth-service/docs/proto/v1"
	"duck_typing_hate/auth-service/internal/usecase"
	"duck_typing_hate/shared/pkg/logger"

	pbgrpc "google.golang.org/grpc"
)

func NewNonceRoutes(app *pbgrpc.Server, n usecase.NonceUseCase, l logger.Logger) {
	r := &V1{n: n, l: l}
	{
		v1.RegisterNonceServer(app, r)
	}
}
