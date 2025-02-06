package entity

import (
	"github.com/SchunckLeonardo/library-management-system/pkg/errors"
	value_objects "github.com/SchunckLeonardo/library-management-system/pkg/value-objects"
	"slices"
	"time"
)

type Patron struct {
	id           value_objects.ID
	name         string
	email        string
	hashPassword value_objects.HashPassword
	violations   int
	books        []Book
}

type PatronInterface interface {
	GetID() value_objects.ID
	GetName() string
	GetEmail() string
	GetHashPassword() value_objects.HashPassword
	GetViolations() int
	GetBooks() []Book
	BorrowBook(book Book) error
	ReturnBook(book Book) error
	ChangeName(newName string) error
	ChangeEmail(newEmail string) error
	IncreaseViolations() error
	DecreaseViolations() error
}

func NewPatron(name, email, plainTextPassword string) (*Patron, error) {
	id := value_objects.NewID()
	hashPassword, err := value_objects.NewHashPassword(plainTextPassword)
	if err != nil {
		return nil, err
	}
	return &Patron{
		id:           *id,
		name:         name,
		email:        email,
		hashPassword: *hashPassword,
		violations:   0,
		books:        []Book{},
	}, nil
}

func (p *Patron) GetID() value_objects.ID {
	return p.id
}

func (p *Patron) GetName() string {
	return p.name
}

func (p *Patron) GetEmail() string {
	return p.email
}

func (p *Patron) GetHashPassword() value_objects.HashPassword {
	return p.hashPassword
}

func (p *Patron) GetViolations() int {
	return p.violations
}

func (p *Patron) GetBooks() []Book {
	return p.books
}

func (p *Patron) ChangeName(newName string) error {
	if newName == "" {
		return errors.ErrStringIsEmpty("name")
	}

	p.name = newName
	return nil
}

func (p *Patron) ChangeEmail(newEmail string) error {
	if newEmail == "" {
		return errors.ErrStringIsEmpty("email")
	}

	p.email = newEmail
	return nil
}

func (p *Patron) IncreaseViolations() error {
	if p.violations >= 3 {
		return errors.ErrPatronViolationsLimit()
	}

	p.violations++
	return nil
}

func (p *Patron) DecreaseViolations() error {
	if p.violations == 0 {
		return errors.ErrPatronDecreaseViolationLessThan0()
	}

	p.violations--
	return nil
}

func (p *Patron) BorrowBook(book *Book) error {
	if p.violations >= 3 {
		return errors.ErrPatronViolationsLimit()
	}
	if len(p.books) >= 2 {
		return errors.ErrPatronBooksLimit()
	}
	if !book.IsAvailable() {
		return errors.ErrBookIsNotAvailable()
	}

	for _, b := range p.books {
		if b.GetID() == book.GetID() {
			return errors.ErrPatronBookAlreadyBorrowed()
		}
	}

	err := book.Borrow()
	if err != nil {
		return err
	}

	p.books = append(p.books, *book)
	return nil
}

func (p *Patron) ReturnBook(book *Book) error {
	for i, b := range p.books {
		if b.GetID() == book.GetID() {
			if time.Now().After(time.UnixMilli(b.GetExpiredBorrowDate())) {
				err := p.IncreaseViolations()
				if err != nil {
					return err
				}
			} else {
				if p.violations > 0 {
					err := p.DecreaseViolations()
					if err != nil {
						return err
					}
				}
			}

			err := book.Return()
			newSlice := slices.Delete(p.books, i, 1)
			p.books = newSlice
			if err != nil {
				return err
			}
		}
	}

	return nil
}
