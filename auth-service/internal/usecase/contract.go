package usecase

import "duck_typing_hate/auth-service/internal/entity"

type (
	NonceUseCase interface {
		Add(pubAddres string) (string, error)
		Verify(sn entity.SignedNonce) error
	}
)
