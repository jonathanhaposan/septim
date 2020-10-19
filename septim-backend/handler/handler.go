package handler

import "github.com/jonathanhaposan/septim/septim-backend/internal/repository"

type Handler struct {
	repository *repository.Repository
}

func Initialize(repository *repository.Repository) *Handler {

	return &Handler{
		repository: repository,
	}
}
