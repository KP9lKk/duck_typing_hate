package common

import "math/rand"

const (
	_letters     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_nonceLength = 12
)

func Generate() string {
	nonce := make([]byte, _nonceLength)
	for i := range _nonceLength {
		nonce[i] = _letters[rand.Intn(len(_letters))]
	}
	return string(nonce)
}
