package repository_test

import (
	"database/sql"
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/internal/infra/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db *sql.DB

func SetUp() {
	db, _ = sql.Open("sqlite3", ":memory:")
	CreateBookTable(db)
}

func CreateBookTable(db *sql.DB) {
	db.Exec(`CREATE TABLE books(id string, title string, author string, description string, available boolean, expired_borrow_date string);`)
}

func TestBookRepository(t *testing.T) {
	SetUp()
	bookRepository := repository.NewBookRepository(db)

	book := entity.NewBook("Harry Potter", "J.K. Rowling", "Fantasy")
	err := bookRepository.Create(book)
	assert.Nil(t, err)

	id := book.GetID().String()
	bookFounded, err := bookRepository.GetById(id)
	assert.Nil(t, err)

	assert.Equal(t, book.GetID(), bookFounded.GetID())
	assert.Equal(t, book.GetTitle(), bookFounded.GetTitle())
	assert.Equal(t, book.GetAuthor(), bookFounded.GetAuthor())
	assert.Equal(t, book.GetDescription(), bookFounded.GetDescription())
}
