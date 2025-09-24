package v1

import (
	v1 "duck_typing_hate/auth-service/docs/proto/v1"
	"duck_typing_hate/auth-service/internal/usecase"
	"duck_typing_hate/shared/pkg/logger"
)

type V1 struct {
	v1.NonceServer
	n usecase.NonceUseCase
	l logger.Logger
}
