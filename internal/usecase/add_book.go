package usecase

import (
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/internal/infra/repository"
)

type AddBookUseCase struct {
	bookRepository repository.BookRepositoryInterface
}

func NewAddBookUseCase(bookRepository repository.
	BookRepositoryInterface) *AddBookUseCase {
	return &AddBookUseCase{bookRepository: bookRepository}
}

func (u *AddBookUseCase) Execute(book entity.Book) error {
	err := u.bookRepository.Create(book)
	if err != nil {
		return err
	}

	return nil
}
