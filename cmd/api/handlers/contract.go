package handlers

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"

	"github.com/raninho/ctps/internal/ctps"
)

type NewCTPSRequest struct {
	DataRegistro string `json:"data_registro"`
	Profissao    string `json:"profissao"`
	Numero       string `json:"numero"`
	Livro        string `json:"livro"`
	Folha        string `json:"folha"`
	DRT          string `json:"drt"`
}

type NewCTPSResponse struct {
	CTPSId string `json:"ctpsId"`
}

func (h *Handler) HandleNewCTPS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctpsRequest := new(NewCTPSRequest)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ctpsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publicKey := h.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		http.Error(w, "error casting public key to ECDSA", http.StatusInternalServerError)
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.EthClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gasPrice, err := h.EthClient.SuggestGasPrice(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth := bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := ctps.DeployCtps(auth, h.EthClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newCtps := new(ctps.CTPSBlock)
	newCtps.CTPSId = address.Hash().Hex()

	err = ctps.InsertCTPS(h.DB, newCtps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctpsTx := new(ctps.CTPSTransaction)
	ctpsTx.CTPSId = newCtps.CTPSId
	ctpsTx.TransactionId = tx.Hash().Hex()

	err = ctps.InsertCTPSTransaction(h.DB, ctpsTx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err = h.EthClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gasPrice, err = h.EthClient.SuggestGasPrice(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth = bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	tx2, err := instance.SetProfissao(auth, ctpsRequest.DataRegistro, ctpsRequest.Profissao, ctpsRequest.Numero, ctpsRequest.Livro, ctpsRequest.Folha, ctpsRequest.DRT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctpsTx2 := new(ctps.CTPSTransaction)
	ctpsTx2.CTPSId = newCtps.CTPSId
	ctpsTx2.TransactionId = tx2.Hash().Hex()

	err = ctps.InsertCTPSTransaction(h.DB, ctpsTx2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := new(NewCTPSResponse)
	response.CTPSId = newCtps.CTPSId

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

type ChangeCTPSResponse struct {
	Message string `json:"message"`
	TxId string `json:"txId"`
}

func (h *Handler) HandleChangeCTPS(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctpsId := vars["ctpsId"]

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctpsRequest := new(ChangeCTPSRequest)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ctpsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publicKey := h.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		http.Error(w, "error casting public key to ECDSA", http.StatusInternalServerError)
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress(ctpsId)

	instance, err := ctps.NewCtps(address, h.EthClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := h.EthClient.PendingNonceAt(r.Context(), fromAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gasPrice, err := h.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth := bind.NewKeyedTransactor(h.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SetProfissao(auth, ctpsRequest.DataRegistro, ctpsRequest.Profissao, ctpsRequest.Numero, ctpsRequest.Livro, ctpsRequest.Folha, ctpsRequest.DRT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctpsTx := new(ctps.CTPSTransaction)
	ctpsTx.CTPSId = ctpsId
	ctpsTx.TransactionId = tx.Hash().Hex()

	err = ctps.InsertCTPSTransaction(h.DB, ctpsTx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := new(ChangeCTPSResponse)
	response.Message = "Changed with success"
	response.TxId = tx.Hash().Hex()

	json.NewEncoder(w).Encode(response)
}

type ViewCTPS struct {
	ReceivedAt string `json:"receivedAt"`
}

func (h *Handler) HandleViewCTPS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	ctpsId := vars["ctpsId"]

	transaction, err := ctps.GetLastCTPSTransactionByCTPSID(h.DB, ctpsId)
	if err != nil {
		log.Println("Error GetLastCTPSTransactionByCTPSID", ctpsId, err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println(transaction.TransactionId)
	log.Println(transaction.CTPSId)

	tx, _, err := h.EthClient.TransactionByHash(ctx, common.HexToHash(transaction.TransactionId))
	if err != nil {
		log.Println("Error BlockByHash", transaction.TransactionId, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := new(ViewCTPS)
	response.ReceivedAt = string(tx.Data())

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
		return
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
		return
	}
	transaction := block.Transaction(common.HexToHash(transactionId))

	response, err := transaction.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
