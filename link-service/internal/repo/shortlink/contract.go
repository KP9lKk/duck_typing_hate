package shortlink

import (
	"context"
	"duck_typing_hate/link-service/entity"
	"duck_typing_hate/shared/repo"
)

type (
	ShortlinkRepo interface {
		repo.Repository
		GetByCode(ctx context.Context, code string) (*entity.ShortLink, error)
	}
)
