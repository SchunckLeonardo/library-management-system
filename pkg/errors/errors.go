package errors

import "errors"

func ErrBookIsNotAvailable() error {
	return errors.New("book is not available")
}

func ErrBookAlreadyInLibrary() error {
	return errors.New("book is already in library")
}

func ErrStringIsEmpty(kind string) error {
	return errors.New(kind + " is empty")
}

func ErrPatronViolationsLimit() error {
	return errors.New("patron reached the limit of violations")
}

func ErrPatronDecreaseViolationLessThan0() error {
	return errors.New("patron violations cannot be less than 0")
}

func ErrPatronBooksLimit() error {
	return errors.New("patron reached the limit of books")
}

func ErrPatronBookAlreadyBorrowed() error {
	return errors.New("patron already borrowed this book")
}

func ErrPatronBookListIsEmpty() error {
	return errors.New("patron book list is empty")
}
