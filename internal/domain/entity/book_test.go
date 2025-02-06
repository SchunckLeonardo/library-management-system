package entity_test

import (
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBook(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.Equal(t, book.GetTitle(), "title")
	assert.Equal(t, book.GetAuthor(), "author")
	assert.Equal(t, book.GetDescription(), "description")

	assert.IsType(t, book, &entity.Book{})
}

func TestBook_GetID(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.NotEmpty(t, book.GetID())
}

func TestBook_GetTitle(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.Equal(t, book.GetTitle(), "title")
}

func TestBook_GetAuthor(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.Equal(t, book.GetAuthor(), "author")
}

func TestBook_GetDescription(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.Equal(t, book.GetDescription(), "description")
}

func TestBook_GetExpiredBorrowDate(t *testing.T) {

	book := entity.NewBook("title", "author", "description")
	assert.NotEmpty(t, book.GetExpiredBorrowDate())

	now := book.GetExpiredBorrowDate()

	assert.Equal(t, book.GetExpiredBorrowDate(), now)
}

func TestBook_IsAvailable(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	assert.True(t, book.IsAvailable())

	err := book.Borrow()
	assert.Nil(t, err)

	assert.False(t, book.IsAvailable())
}

func TestBook_ChangeAuthor(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	err := book.ChangeAuthor("new author")
	assert.Nil(t, err)

	assert.Equal(t, book.GetAuthor(), "new author")
}

func TestBook_ChangeDescription(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	err := book.ChangeDescription("new description")
	assert.Nil(t, err)

	assert.Equal(t, book.GetDescription(), "new description")
}

func TestBook_ChangeTitle(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	err := book.ChangeTitle("new title")
	assert.Nil(t, err)

	assert.Equal(t, book.GetTitle(), "new title")
}

func TestBook_Borrow(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	err := book.Borrow()
	assert.Nil(t, err)

	assert.False(t, book.IsAvailable())
	assert.Equal(t, time.UnixMilli(book.GetExpiredBorrowDate()).Day(), time.Now().AddDate(0, 0, 7).Day())

	err = book.Borrow()
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrBookIsNotAvailable())
}

func TestBook_Return(t *testing.T) {
	book := entity.NewBook("title", "author", "description")
	err := book.Borrow()
	assert.Nil(t, err)

	assert.False(t, book.IsAvailable())

	err = book.Return()
	assert.Nil(t, err)

	assert.True(t, book.IsAvailable())

	err = book.Return()
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrBookAlreadyInLibrary())
}
