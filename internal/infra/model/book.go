package model

import "time"

type Book struct {
	ID                string
	Title             string
	Author            string
	Description       string
	Available         bool
	ExpiredBorrowDate time.Time
}
