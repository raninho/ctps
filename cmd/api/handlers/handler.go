package handlers

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Handler struct {
	EthClient  *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
}
