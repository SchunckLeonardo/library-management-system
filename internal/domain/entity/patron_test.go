package entity_test

import (
	"bou.ke/monkey"
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPatron(t *testing.T) {
	patron, err := entity.NewPatron("name", "email", "password")
	assert.Nil(t, err)

	assert.Equal(t, patron.GetName(), "name")
	assert.Equal(t, patron.GetEmail(), "email")
}

func TestPatron_GetID(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.NotEmpty(t, patron.GetID())
}

func TestPatron_GetName(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.Equal(t, patron.GetName(), "name")
}

func TestPatron_GetEmail(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.Equal(t, patron.GetEmail(), "email")
}

func TestPatron_GetHashPassword(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.NotEmpty(t, patron.GetHashPassword())
}

func TestPatron_GetViolations(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.Equal(t, patron.GetViolations(), 0)
}

func TestPatron_GetBooks(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	assert.Empty(t, patron.GetBooks())
}

func TestPatron_ChangeName(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	err := patron.ChangeName("new name")
	assert.Nil(t, err)
	assert.Equal(t, patron.GetName(), "new name")
}

func TestPatron_ChangeEmail(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	err := patron.ChangeEmail("new email")
	assert.Nil(t, err)
	assert.Equal(t, patron.GetEmail(), "new email")
}

func TestPatron_IncreaseViolations(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	err := patron.IncreaseViolations()
	assert.Nil(t, err)
	assert.Equal(t, patron.GetViolations(), 1)

	err = patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.IncreaseViolations()
	assert.NotNil(t, err)

	assert.Equal(t, patron.GetViolations(), 3)
	assert.IsType(t, err, errors.ErrPatronViolationsLimit())
}

func TestPatron_DecreaseViolations(t *testing.T) {
	patron, _ := entity.NewPatron("name", "email", "password")
	err := patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.DecreaseViolations()
	assert.Nil(t, err)
	assert.Equal(t, patron.GetViolations(), 0)

	err = patron.DecreaseViolations()
	assert.NotNil(t, err)
	assert.Equal(t, patron.GetViolations(), 0)
	assert.IsType(t, err, errors.ErrPatronDecreaseViolationLessThan0())
}

func TestPatron_BorrowBook(t *testing.T) {
	book := entity.NewBook("Harry Potter", "J.K. Rowling", "harry potter book")
	book2 := entity.NewBook("Harry Potter 2", "J.K. Rowling", "harry potter book 2")
	book3 := entity.NewBook("Harry Potter 3", "J.K. Rowling", "harry potter book 3")

	patron, err := entity.NewPatron("Leonardo", "leo@gmail.com", "123456")
	assert.Nil(t, err)

	err = patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.IncreaseViolations()
	assert.Nil(t, err)

	err = patron.BorrowBook(book)
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrPatronViolationsLimit())

	err = patron.DecreaseViolations()
	assert.Nil(t, err)

	err = patron.BorrowBook(book)
	assert.Nil(t, err)
	assert.Equal(t, patron.GetBooks()[0].GetTitle(), book.GetTitle())
	assert.Equal(t, len(patron.GetBooks()), 1)

	err = patron.BorrowBook(book)
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrPatronBookAlreadyBorrowed())

	err = patron.BorrowBook(book2)
	assert.Nil(t, err)
	assert.Equal(t, patron.GetBooks()[1].GetTitle(), book2.GetTitle())
	assert.Equal(t, len(patron.GetBooks()), 2)

	err = patron.BorrowBook(book3)
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrPatronBooksLimit())
	assert.Equal(t, len(patron.GetBooks()), 2)
}

func TestPatron_ReturnBook(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2025, 2, 3, 17, 0, 0, 0, time.UTC)
	})
	defer monkey.Unpatch(time.Now)

	book := entity.NewBook("Harry Potter", "J.K. Rowling", "harry potter book")

	patron, err := entity.NewPatron("Leonardo", "leo@gmail.com", "123456")
	assert.Nil(t, err)

	err = patron.BorrowBook(book)
	assert.Nil(t, err)

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2025, 2, 13, 17, 0, 0, 0, time.UTC)
	})
	defer monkey.Unpatch(time.Now)

	err = patron.ReturnBook(book)
	assert.Nil(t, err)
	assert.Empty(t, patron.GetBooks())

	assert.Equal(t, 1, patron.GetViolations())
}
