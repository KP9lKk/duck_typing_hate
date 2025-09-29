package shortlink

import (
	"context"
	"duck_typing_hate/link-service/internal/entity"
)

type (
	ShortlinkRepo interface {
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, sl entity.ShortLink) error
		Create(ctx context.Context, sl *entity.ShortLink) error
		GetById(ctx context.Context, id int) (*entity.ShortLink, error)
		GetAll(ctx context.Context) (*[]entity.ShortLink, error)
		GetByCode(ctx context.Context, code string) (*entity.ShortLink, error)
	}
)
