package service

import (
	"context"
	"duck_typing_hate/link-service/entity"
	"duck_typing_hate/link-service/internal/service/request"
)

type ShortLinkService interface {
	GetByCode(ctx context.Context, code string) (*entity.ShortLink, error)
	RedirectByCode(ctx context.Context, code string) (*entity.ShortLink, error)
	Create(ctx context.Context, rq request.ShortLinkCreateRequest) (*entity.ShortLink, error)
}
