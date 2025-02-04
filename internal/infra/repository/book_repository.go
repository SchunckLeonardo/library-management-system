package repository

import (
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/internal/infra/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book entity.BookInterface) error {
	return r.db.Model(model.Book{}).Create(&book).Error
}

func (r *BookRepository) GetById(id string) (entity.BookInterface, error) {
	var book entity.BookInterface
	err := r.db.Model(model.Book{}).First(&book, id).Error
	return book, err
}

func (r *BookRepository) FetchAll() ([]entity.BookInterface, error) {
	var books []entity.BookInterface
	err := r.db.Model(model.Book{}).Find(&books).Error
	return books, err
}

func (r *BookRepository) Update(book entity.BookInterface) error {
	return r.db.Model(model.Book{}).Save(&book).Error
}

func (r *BookRepository) Delete(id string) error {
	return r.db.Model(model.Book{}).Delete(&model.Book{}, id).Error
}
