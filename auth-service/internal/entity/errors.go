package entity

import "errors"

var (
	ErrInvalidSignature = errors.New("your signature is not valid")
	ErrNonceNotFound    = errors.New("your nonce is not found")
)
