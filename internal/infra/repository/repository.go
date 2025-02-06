package repository

import "github.com/SchunckLeonardo/library-management-system/internal/domain/entity"

type BookRepositoryInterface interface {
	Create(book entity.Book) error
	GetById(id string) (entity.Book, error)
	ListAll() ([]entity.Book, error)
	Update(book entity.Book) error
	Delete(id string) error
}

type PatronRepositoryInterface interface {
	Create(patron entity.Patron) error
	Update(patron entity.Patron) error
	Delete(id string) error
}
