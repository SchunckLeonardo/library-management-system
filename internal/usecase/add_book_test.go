package usecase_test

import (
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"github.com/SchunckLeonardo/library-management-system/internal/usecase"
	"github.com/SchunckLeonardo/library-management-system/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddBookUseCase_Execute(t *testing.T) {
	mockRepository := mocks.BookRepositoryInterface{}
	sut := usecase.NewAddBookUseCase(&mockRepository)

	book := entity.NewBook("Book 1", "Author 1", "Description 1")

	mockRepository.On("Create", mock.AnythingOfType("entity.Book")).Return(
		nil).Times(1)

	err := sut.Execute(*book)
	assert.Nil(t, err)
}
