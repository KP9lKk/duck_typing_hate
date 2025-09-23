package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(pubAddres string, nonce string, signedNonce string) error {

	publicAddres := common.HexToAddress(pubAddres)
	hash := crypto.Keccak256Hash([]byte(nonce))
	data := hexutil.MustDecode(signedNonce)
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), data)
	if err != nil {
		return err
	}
	sigAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	matches := publicAddres == sigAddress
	if !matches {
		return err
	}
	return nil
}
