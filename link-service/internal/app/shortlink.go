package app

import (
	"context"
	"duck_typing_hate/link-service/internal/repo/shortlink/persistent"
	"duck_typing_hate/link-service/internal/service/shortlink"
	"duck_typing_hate/shared/pkg/postgres"
)

func Run() {
	pg, err := postgres.New("url", context.Background())
	if err != nil {

	}
	repo := persistent.New(pg)

	service := shortlink.New(repo)
	service.GetByCode(context.Background(), "123")
}
