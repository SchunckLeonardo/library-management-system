package repository_test

import (
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/internal/infra/model"
	"github.com/SchunckLeonardo/library-management-system/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func SetUp() {
	db, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{})
}

func TestBookRepository(t *testing.T) {
	SetUp()
	bookRepository := repository.NewBookRepository(db)

	book := entity.NewBook("Harry Potter", "J.K. Rowling", "Fantasy")
	err := bookRepository.Create(book)
	assert.Nil(t, err)

	var bookFounded entity.BookInterface
	err = db.Model(model.Book{}).First(&bookFounded, book.GetID()).Error
	assert.Nil(t, err)

	assert.Equal(t, book.GetID(), bookFounded.GetID())
	assert.Equal(t, book.GetTitle(), bookFounded.GetTitle())
	assert.Equal(t, book.GetAuthor(), bookFounded.GetAuthor())
	assert.Equal(t, book.GetDescription(), bookFounded.GetDescription())
}
