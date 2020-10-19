package handler

import (
	"io/ioutil"
	"log"
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

func (h *Handler) AddTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) webserver.Response {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("failed read body", err)
		return webserver.Failed(err.Error())
	}
	defer r.Body.Close()

	var trx = &model.Transaction{}
	err = jsoni.Unmarshal(body, trx)
	if err != nil {
		log.Println("failed unmarshal json", string(body), err)
		return webserver.Failed(err.Error())
	}

	err = h.repository.InsertOneTransaction(trx)
	if err != nil {
		return webserver.Failed(err.Error())
	}

	return webserver.Success(nil)
}
