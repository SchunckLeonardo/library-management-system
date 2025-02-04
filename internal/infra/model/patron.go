package model

type Patron struct {
	ID           string
	Name         string
	Email        string
	HashPassword string
	Violations   int
	Books        []Book
}
