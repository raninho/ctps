package routers

import (
	"github.com/gorilla/mux"

	apiHandler "github.com/raninho/ctps/cmd/api/handlers"
)

//Router ...
func Router(h *apiHandler.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ctps", h.HandleNewCTPS).Methods("POST")
	router.HandleFunc("/ctps/{ctpsId}", h.HandleChangeCTPS).Methods("PUT")
	router.HandleFunc("/ctps/{ctpsId}", h.HandleViewCTPS).Methods("GET")

	router.HandleFunc("/ctps/{ctpsId}/transaction", h.HandleViewAllTransactions).Methods("GET")
	router.HandleFunc("/ctps/{ctpsId}/transaction/{transacationId}", h.HandleViewTransaction).Methods("GET")

	return router
}
