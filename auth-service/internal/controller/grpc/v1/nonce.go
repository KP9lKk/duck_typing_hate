package v1

import (
	"context"
	v1 "duck_typing_hate/auth-service/docs/proto/v1"
	"duck_typing_hate/auth-service/internal/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *V1) Get(ctx context.Context, request *v1.GetNonceRequest) (*v1.GetNonceResponse, error) {
	nonce, err := r.n.Add(ctx, request.PublicAddres)
	if err != nil {
		return nil, err
	}
	return &v1.GetNonceResponse{Nonce: nonce}, nil
}

func (r *V1) Verify(ctx context.Context, request *v1.VerifyNonceRequest) (*v1.VerifyNonceResponse, error) {
	signedNonce := &entity.SignedNonce{
		SignedNonce:  request.SignedNonce,
		PublicAddres: request.PublicAddres,
	}
	err := r.n.Verify(ctx, *signedNonce)
	if err != nil {
		if err == entity.ErrNonceNotFound {
			return &v1.VerifyNonceResponse{Succsess: false}, status.New(codes.NotFound, err.Error()).Err()
		}
		return &v1.VerifyNonceResponse{Succsess: false}, status.New(codes.Internal, err.Error()).Err()
	}
	return &v1.VerifyNonceResponse{Succsess: true, Address: request.PublicAddres}, nil
}
