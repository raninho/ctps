package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	apiHandler "github.com/raninho/ctps/cmd/api/handlers"
	apiRouter "github.com/raninho/ctps/cmd/api/routers"
)

var (
	priKey string
	ethURI string
)

func init() {
	flag.StringVar(&priKey, "privateKey", os.Getenv("ETH_PRIVATE_KEY"), "-privateKey=...")
	flag.StringVar(&ethURI, "ethURI", os.Getenv("ETH_URI"), "-ethURI=localhost:9595")
}

func main() {
	flag.Parse()

	client, err := ethclient.Dial(ethURI)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Fatal(err)
	}

	h := &apiHandler.Handler{EthClient: client, PrivateKey: privateKey}
	routes := apiRouter.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", "9999"),
		Handler:      routes,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 40,
	}

	log.Println("Initializing ListenAndServe")
	if err := s.ListenAndServe(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
