package router

import (
	"net/http"

	"github.com/jonathanhaposan/septim/septim-backend/common/webserver"
	"github.com/jonathanhaposan/septim/septim-backend/handler"
	jsoniter "github.com/json-iterator/go"

	"github.com/julienschmidt/httprouter"
)

var (
	jsoni = jsoniter.ConfigCompatibleWithStandardLibrary
	r     *httprouter.Router
	hndlr *handler.Handler
)

type handlerFunc func(http.ResponseWriter, *http.Request, httprouter.Params) webserver.Response

func InitializeRoute(handler *handler.Handler) *httprouter.Router {
	r = httprouter.New()

	hndlr = handler

	allowCORS()
	registerRoutes()

	return r
}

func registerRoutes() {
	register(http.MethodGet, "/transaction", hndlr.GetTransactionList)
	register(http.MethodGet, "/transaction/:id", nil)
	register(http.MethodPut, "/transaction/:id", nil)
	register(http.MethodPost, "/transaction", hndlr.AddTransaction)
}

func allowCORS() {
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func register(method, path string, fn handlerFunc) {
	r.Handle(method, path, handle(fn))
}

func handle(fn handlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := fn(w, r, ps)

		b, err := jsoni.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.Header().Add("Content-Type", "application/json")

		if result.Success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)

		}

		w.Write(b)
	}
}
