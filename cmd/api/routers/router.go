package routers

import (
	"github.com/gorilla/mux"

	apiHandler "github.com/raninho/ctps/cmd/api/handlers"
)

//Router ...
func Router(h *apiHandler.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/status/{idContract}", h.HandleViewContract).Methods("GET")

	return router
}
