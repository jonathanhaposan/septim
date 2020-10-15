package handler

import (
	"net/http"

	"github.com/jonathanhaposan/septim/septim-backend/common/webserver"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetTransactionList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) webserver.Response {

	str := `{"test": 123}`

	return webserver.Success(str)
}
