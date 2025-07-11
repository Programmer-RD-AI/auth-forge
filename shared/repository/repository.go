package repository

type Repository interface {
	Connect() (any, error) 
	Create(data any) (any, error)
	Read(id string) (any, error)
	Update(id string, data any) (any, error)
	Delete(id string) error
	Close() error
}

