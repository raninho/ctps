package handlers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleViewContract(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idContract := vars["idContract"]

	tx, pending, _ := h.EthClient.TransactionByHash(ctx, common.HexToHash("0x"+idContract))
	if !pending {
		fmt.Println(tx)
	}
}
