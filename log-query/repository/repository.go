package repository

import (
	"github.com/monzo/gocassa"
)

type Repository struct {
	DB gocassa.KeySpace
}

func New(db gocassa.KeySpace) *Repository {
	return &Repository{db}
}
