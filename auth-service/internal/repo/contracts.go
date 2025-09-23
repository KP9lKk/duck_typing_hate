package repo

import "duck_typing_hate/auth-service/internal/entity"

type (
	NonceRepo interface {
		Add(nonce entity.Nonce) error
		Get(pubAddres string) (entity.Nonce, error)
		Generate() string
	}
)
