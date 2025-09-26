package persistent

import (
	"context"
	"duck_typing_hate/auth-service/internal/entity"
	"duck_typing_hate/shared/pkg/reddis"
)

type NonceRepo struct {
	rdb *reddis.Reddis
}

func New(rdb *reddis.Reddis) *NonceRepo {
	return &NonceRepo{
		rdb: rdb,
	}
}

func (r *NonceRepo) Add(ctx context.Context, nonce entity.Nonce) error {
	err := r.rdb.Set(ctx, nonce.PublicAddres, nonce.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (r *NonceRepo) Get(ctx context.Context, pubAddres string) (entity.Nonce, error) {
	nonce, err := r.rdb.Get(ctx, pubAddres)
	if err != nil {
		return entity.Nonce{}, err
	}

	return entity.Nonce{
		PublicAddres: pubAddres,
		Nonce:        nonce,
	}, nil
}
