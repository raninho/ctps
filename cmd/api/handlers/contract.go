package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	ctps "github.com/raninho/ctps/internal/store"
)

func (h *Handler) HandleViewContract(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idContract := vars["idContract"]

	block, err := h.EthClient.BlockByHash(ctx, common.HexToHash("0x"+idContract))
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, block)
}

func (h *Handler) HandleViewTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idContract := vars["idContract"]

	block, err := h.EthClient.BlockByHash(ctx, common.HexToHash("0x"+idContract))
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	transactions := block.Transactions()
	fmt.Fprint(w, transactions)
}

func (h *Handler) HandleNewContract(w http.ResponseWriter, r *http.Request) {
	publicKey := h.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.EthClient.PendingNonceAt(r.Context(), fromAddress)
	if err != nil {
		log.Fatal("PendingNonceAt err", err)
	}

	gasPrice, err := h.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("SuggestGasPrice err", err)
	}

	auth := bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	parsed, err := abi.JSON(strings.NewReader(ctps.StoreABI))
	if err != nil {
		log.Println("1", err.Error())
	}

	address, _, contract, err := bind.DeployContract(auth, parsed, nil, h.EthClient)
	if err != nil {
		log.Println("2", err.Error())
	}

	fmt.Println(address.Hex())
	fmt.Println(contract)

	fmt.Fprint(w, address.Hex())
}
