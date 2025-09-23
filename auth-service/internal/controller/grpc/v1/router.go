package v1

import (
	v1 "duck_typing_hate/auth-service/docs/proto/v1"
	"duck_typing_hate/auth-service/internal/usecase"

	pbgrpc "google.golang.org/grpc"
)

func NewNonceRoutes(app *pbgrpc.Server, n usecase.NonceUseCase) {
	r := &V1{n: n}
	{
		v1.RegisterNonceServer(app, r)
	}
}
