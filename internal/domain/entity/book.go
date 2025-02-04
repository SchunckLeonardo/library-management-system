package entity

import (
	"github.com/SchunckLeonardo/library-management-system/pkg/errors"
	value_objects "github.com/SchunckLeonardo/library-management-system/pkg/value-objects"
	"time"
)

type Book struct {
	id                value_objects.ID
	title             string
	author            string
	description       string
	available         bool
	expiredBorrowDate time.Time
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
		id:                *id,
		title:             title,
		author:            author,
		description:       description,
		available:         true,
		expiredBorrowDate: time.Now(),
	}
}

func (b *Book) GetID() value_objects.ID {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetDescription() string {
	return b.description
}

func (b *Book) GetExpiredBorrowDate() time.Time {
	return b.expiredBorrowDate
}

func (b *Book) IsAvailable() bool {
	return b.available
}

func (b *Book) ChangeTitle(newTitle string) error {
	if newTitle == "" {
		return errors.ErrStringIsEmpty("title")
	}

	b.title = newTitle
	return nil
}

func (b *Book) ChangeAuthor(newAuthor string) error {
	if newAuthor == "" {
		return errors.ErrStringIsEmpty("author")
	}

	b.author = newAuthor
	return nil
}

func (b *Book) ChangeDescription(newDescription string) error {
	if newDescription == "" {
		return errors.ErrStringIsEmpty("description")
	}

	b.description = newDescription
	return nil
}

func (b *Book) Borrow() error {
	if !b.available && !b.expiredBorrowDate.Before(time.Now()) {
		return errors.ErrBookIsNotAvailable()
	}

	b.available = false
	b.expiredBorrowDate = time.Now().AddDate(0, 0, 7)
	return nil
}

func (b *Book) Return() error {
	if b.available {
		return errors.ErrBookAlreadyInLibrary()
	}

	b.available = true
	b.expiredBorrowDate = time.Now()
	return nil
}
