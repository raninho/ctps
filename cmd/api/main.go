package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jinzhu/gorm/dialects/postgres"


	apiHandler "github.com/raninho/ctps/cmd/api/handlers"
	apiRouter "github.com/raninho/ctps/cmd/api/routers"
	"github.com/raninho/ctps/internal/ctps"
)

var (
	postgresURI string
	priKey string
	ethURI string
)

func init() {
	flag.StringVar(&postgresURI, "postgresURI", os.Getenv("POSTGRES_URI"), "-postgresURI=postgres://postgres:@localhost/postgres?sslmode=disable")
	flag.StringVar(&priKey, "privateKey", os.Getenv("ETH_PRIVATE_KEY"), "-privateKey=...")
	flag.StringVar(&ethURI, "ethURI", os.Getenv("ETH_URI"), "-ethURI=localhost:9595")
}

func main() {
	flag.Parse()

	db, err := gorm.Open("postgres", postgresURI)
	if err != nil {
		log.Fatal("can't initialize Postgres. " + err.Error())
		panic("Error DB: " + err.Error())
	}
	defer db.Close()

	log.Println("Initializing AutoMigrate CTPSBlock into Postgres")
	if err := db.AutoMigrate(&ctps.CTPSBlock{}).Error; err != nil {
		log.Println("can't AutoMigrate CTPSBlock into Postgres. " + err.Error())
	}

	log.Println("Initializing AutoMigrate CTPSTransaction into Postgres")
	if err := db.AutoMigrate(&ctps.CTPSTransaction{}).Error; err != nil {
		log.Println("can't AutoMigrate CTPSTransaction into Postgres. " + err.Error())
	}

	client, err := ethclient.Dial(ethURI)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Fatal(err)
	}

	h := &apiHandler.Handler{DB: db, EthClient: client, PrivateKey: privateKey}
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
