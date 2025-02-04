package repository

import (
	"database/sql"
	"errors"
	"github.com/SchunckLeonardo/library-management-system/internal/domain/entity"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book entity.BookInterface) error {
	stmt, err := r.db.Prepare(`INSERT INTO books (id, title, author, description, available, expired_borrow_date) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	dateFormatted := book.GetExpiredBorrowDate().Format(time.RFC3339)

	_, err = stmt.Exec(book.GetID(), book.GetTitle(), book.GetAuthor(), book.GetDescription(), book.IsAvailable(), dateFormatted)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) GetById(id string) (entity.BookInterface, error) {
	var book entity.Book
	stmt, err := r.db.Prepare(`SELECT id, title, author, description, available, expired_borrow_date FROM books WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var dateString string

	err = stmt.QueryRow(id).Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Available, &dateString)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	book.ExpiredBorrowDate, _ = time.Parse(time.RFC3339, dateString)

	return &book, nil
}
