package v1

import (
	"duck_typing_hate/auth-service/internal/usecase"
	"duck_typing_hate/shared/pkg/logger"
	v1 "duck_typing_hate/shared/proto/v1/nonce"

	pbgrpc "google.golang.org/grpc"
)

func NewNonceRoutes(app *pbgrpc.Server, n usecase.NonceUseCase, l logger.Logger) {
	r := &V1{n: n, l: l}
	{
		v1.RegisterNonceServer(app, r)
	}
}
