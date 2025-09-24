package v1

import (
	"duck_typing_hate/auth-service/internal/usecase"
	"duck_typing_hate/shared/pkg/logger"
	v1 "duck_typing_hate/shared/proto/v1/nonce"
)

type V1 struct {
	v1.NonceServer
	n usecase.NonceUseCase
	l logger.Logger
}
