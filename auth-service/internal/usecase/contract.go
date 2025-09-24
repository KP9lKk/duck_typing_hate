package usecase

import (
	"context"
	"duck_typing_hate/auth-service/internal/entity"
)

type (
	NonceUseCase interface {
		Add(ctx context.Context, pubAddres string) (string, error)
		Verify(ctx context.Context, sn entity.SignedNonce) error
	}
)
