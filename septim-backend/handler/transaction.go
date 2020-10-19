package handler

import (
	"net/http"

	"github.com/jonathanhaposan/septim/septim-backend/common/consts"
	"github.com/jonathanhaposan/septim/septim-backend/common/webserver"
	"github.com/jonathanhaposan/septim/septim-backend/internal/model"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetTransactionList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) webserver.Response {

	result := h.repository.GetAllTransaction()

	summary := model.Summary{}
	for _, tx := range result {
		if tx.Type == consts.BUY {
			summary.TotalBuyStock += tx.Price
		}

		if tx.Type == consts.SELL {
			summary.TotalSellStock += tx.Price
		}
	}

	return webserver.Success(&webserver.GetTransactionResponse{
		Summary:      summary,
		Transactions: result,
	})
}
