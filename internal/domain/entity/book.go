package entity

import (
	"github.com/SchunckLeonardo/library-management-system/pkg/errors"
	value_objects "github.com/SchunckLeonardo/library-management-system/pkg/value-objects"
	"time"
)

type Book struct {
	ID                value_objects.ID
	Title             string
	Author            string
	Description       string
	Available         bool
	ExpiredBorrowDate time.Time
}

type BookInterface interface {
	Borrow() error
	Return() error
	GetID() value_objects.ID
	GetTitle() string
	GetAuthor() string
	GetDescription() string
	GetExpiredBorrowDate() time.Time
	IsAvailable() bool
	ChangeTitle(newTitle string) error
	ChangeAuthor(newAuthor string) error
	ChangeDescription(newDescription string) error
}

func NewBook(title, author, description string) *Book {
	id := value_objects.NewID()
	return &Book{
		ID:                *id,
		Title:             title,
		Author:            author,
		Description:       description,
		Available:         true,
		ExpiredBorrowDate: time.Now(),
	}
}

func (b *Book) GetID() value_objects.ID {
	return b.ID
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetDescription() string {
	return b.Description
}

func (b *Book) GetExpiredBorrowDate() time.Time {
	return b.ExpiredBorrowDate
}

func (b *Book) IsAvailable() bool {
	return b.Available
}

func (b *Book) ChangeTitle(newTitle string) error {
	if newTitle == "" {
		return errors.ErrStringIsEmpty("title")
	}

	b.Title = newTitle
	return nil
}

func (b *Book) ChangeAuthor(newAuthor string) error {
	if newAuthor == "" {
		return errors.ErrStringIsEmpty("author")
	}

	b.Author = newAuthor
	return nil
}

func (b *Book) ChangeDescription(newDescription string) error {
	if newDescription == "" {
		return errors.ErrStringIsEmpty("description")
	}

	b.Description = newDescription
	return nil
}

func (b *Book) Borrow() error {
	if !b.Available && !b.ExpiredBorrowDate.Before(time.Now()) {
		return errors.ErrBookIsNotAvailable()
	}

	b.Available = false
	b.ExpiredBorrowDate = time.Now().AddDate(0, 0, 7)
	return nil
}

func (b *Book) Return() error {
	if b.Available {
		return errors.ErrBookAlreadyInLibrary()
	}

	b.Available = true
	b.ExpiredBorrowDate = time.Now()
	return nil
}
