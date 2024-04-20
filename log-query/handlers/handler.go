// Package handler defines the API and DB interface
package handler

import (
	"github.com/mhope-2/log-query/repository"
	"github.com/monzo/gocassa"
)

type Handler struct {
	Repo *repository.Repository
}

func New(db gocassa.KeySpace) *Handler {
	repo := repository.New(db)

	return &Handler{
		Repo: repo,
	}
}
