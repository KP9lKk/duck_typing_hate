package repo

type Repository interface {
	GetById(id any) (*any, error)
	GetAll() []*any
	Create(v *any) error
	Update(v *any) error
	Delete(id any) error
}
