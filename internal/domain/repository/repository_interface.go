package repository

type RepositoryInterface[T any] interface {
	Create(T) error
	GetById(string) (T, error)
	FetchAll() ([]T, error)
	Update(T) error
	Delete(string) error
}
