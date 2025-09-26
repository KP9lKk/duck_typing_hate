package nonce

import (
	"context"
	"duck_typing_hate/auth-service/internal/common"
	"duck_typing_hate/auth-service/internal/entity"
	"duck_typing_hate/auth-service/internal/repo"
	sh "duck_typing_hate/shared/common"

	"github.com/redis/go-redis/v9"
)

type NonceUseCase struct {
	repo repo.NonceRepo
}

func New(r repo.NonceRepo) *NonceUseCase {
	return &NonceUseCase{r}
}

func (nuc *NonceUseCase) Add(ctx context.Context, pubAddres string) (string, error) {
	nonce := &entity.Nonce{}
	nonce.Nonce = sh.Generate()[:]
	nonce.PublicAddres = pubAddres
	err := nuc.repo.Add(ctx, *nonce)
	if err != nil {
		return "", err
	}
	return nonce.Nonce, nil
}

func (nuc *NonceUseCase) Verify(ctx context.Context, sn entity.SignedNonce) error {
	nonce, err := nuc.repo.Get(ctx, sn.PublicAddres)
	if err != nil {
		if err == redis.Nil {
			return entity.ErrNonceNotFound
		}
		return err
	}
	err = common.VerifySignature(sn.PublicAddres, nonce.Nonce, sn.SignedNonce)
	return err
}
