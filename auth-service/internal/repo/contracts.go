package repo

import (
	"context"
	"duck_typing_hate/auth-service/internal/entity"
)

type (
	NonceRepo interface {
		Add(ctx context.Context, nonce entity.Nonce) error
		Get(ctx context.Context, pubAddres string) (entity.Nonce, error)
	}
)
