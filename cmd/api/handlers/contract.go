package handlers

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"

	ctps "github.com/raninho/ctps/internal/store"
)

type NewCTPSResponse struct {
	BlockID string `json:"blockId"`
}

func (h *Handler) HandleNewCTPS(w http.ResponseWriter, r *http.Request) {
	publicKey := h.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		http.Error(w, "error casting public key to ECDSA", http.StatusInternalServerError)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.EthClient.PendingNonceAt(r.Context(), fromAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	gasPrice, err := h.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	auth := bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := ctps.DeployCtps(auth, h.EthClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("transaction:", tx.Hash().Hex())

	response := new(NewCTPSResponse)
	response.BlockID = address.Hex()

	json.NewEncoder(w).Encode(response)
}

type ChangeCTPSRequest struct {
	DataRegistro string `json:"data_registro"`
	Profissao    string `json:"profissao"`
	Numero       string `json:"numero"`
	Livro        string `json:"livro"`
	Folha        string `json:"folha"`
	DRT          string `json:"drt"`
}

func (h *Handler) HandleChangeCTPS(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctpsId := vars["ctpsId"]

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctpsRequest := new(ChangeCTPSRequest)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ctpsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	publicKey := h.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		http.Error(w, "error casting public key to ECDSA", http.StatusInternalServerError)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress(ctpsId)

	instance, err := ctps.NewCtps(address, h.EthClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	nonce, err := h.EthClient.PendingNonceAt(r.Context(), fromAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	gasPrice, err := h.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	auth := bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SetProfissao(auth, ctpsRequest.DataRegistro, ctpsRequest.Profissao, ctpsRequest.Numero, ctpsRequest.Livro, ctpsRequest.Folha, ctpsRequest.DRT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transaction:", tx.Hash().Hex())
}

type ViewCTPS struct {
	ReceivedAt time.Time `json:"receivedAt"`
}

func (h *Handler) HandleViewCTPS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idCTPS := vars["ctpsId"]

	block, err := h.EthClient.BlockByHash(ctx, common.HexToHash(idCTPS))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := new(ViewCTPS)
	response.ReceivedAt = block.ReceivedAt

	json.NewEncoder(w).Encode(response)
}

type ViewAllTransactions struct {
	Transactions []Transacation `json:"transactions"`
}

type Transacation struct {
	TransactionID string `json:"transactionId"`
}

func (h *Handler) HandleViewAllTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idCTPS := vars["ctpsId"]

	block, err := h.EthClient.BlockByHash(ctx, common.HexToHash(idCTPS))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	transactions := block.Transactions()

	response := new(ViewAllTransactions)
	for _, t := range transactions {
		transaction := new(Transacation)
		transaction.TransactionID = t.Hash().Hex()
		response.Transactions = append(response.Transactions, *transaction)
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleViewTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	transactionId := vars["transactionId"]

	block, err := h.EthClient.BlockByHash(ctx, common.HexToHash(transactionId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	transaction := block.Transaction(common.HexToHash(transactionId))

	response, err := transaction.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(response)
}
