package repo

import "context"

type Repository interface {
	GetById(ctx context.Context, id any) (*any, error)
	GetAll(ctx context.Context) ([]*any, error)
	Create(ctx context.Context, v *any) error
	Update(ctx context.Context, v *any) error
	Delete(ctx context.Context, id any) error
}
