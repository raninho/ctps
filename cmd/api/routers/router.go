package routers

import (
	"github.com/gorilla/mux"

	apiHandler "github.com/raninho/ctps/cmd/api/handlers"
)

//Router ...
func Router(h *apiHandler.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ctps/new", h.HandleNewContract).Methods("GET")
	router.HandleFunc("/ctps/{idContract}", h.HandleViewContract).Methods("GET")
	router.HandleFunc("/ctps/{idContract}/transactions", h.HandleViewTransactions).Methods("GET")

	return router
}
