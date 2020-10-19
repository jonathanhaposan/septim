package handler

import (
	"github.com/jonathanhaposan/septim/septim-backend/internal/repository"
	jsoniter "github.com/json-iterator/go"
)

type Handler struct {
	repository *repository.Repository
}

var (
	jsoni = jsoniter.ConfigFastest
)

func Initialize(repository *repository.Repository) *Handler {

	return &Handler{
		repository: repository,
	}
}
