package shortlink

import (
	"context"
	"duck_typing_hate/link-service/entity"
	"duck_typing_hate/link-service/internal/repo/shortlink"
	"duck_typing_hate/link-service/internal/service/request"
	"duck_typing_hate/shared/common"
)

type ShortLinkService struct {
	r shortlink.ShortlinkRepo
}

func New(r *shortlink.ShortlinkRepo) *ShortLinkService {
	return &ShortLinkService{r: *r}
}

func (s *ShortLinkService) GetByCode(ctx context.Context, code string) (*entity.ShortLink, error) {
	result, err := s.r.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ShortLinkService) RedirectByCode(ctx context.Context, code string) (*entity.ShortLink, error) {
	result, err := s.r.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	result.Clicks += 1
	err = s.r.Update(ctx, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *ShortLinkService) Create(ctx context.Context, rq request.ShortLinkCreateRequest) (*entity.ShortLink, error) {
	sl := &entity.ShortLink{
		Owner:       rq.Owner,
		OriginalUrl: rq.OriginalUrl,
		ShortCode:   common.Generate(),
	}
	err := s.r.Create(ctx, sl)
	if err != nil {
		return nil, err
	}
	return sl, nil
}
