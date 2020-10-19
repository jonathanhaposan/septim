package webserver

import "github.com/jonathanhaposan/septim/septim-backend/internal/model"

type GetTransactionResponse struct {
	Summary      model.Summary        `json:"summary"`
	Transactions []*model.Transaction `json:"transactions"`
}
