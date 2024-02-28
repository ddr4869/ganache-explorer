package config

import (
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Config struct {
	Ganache GanacheConf
	Sepolia SepoliaConf
	Gin     GinConf
	Client  *ethclient.Client
}

type GanacheConf struct {
	Client    *ethclient.Client
	Addresses []string
}

type SepoliaConf struct {
	Client *ethclient.Client
}

type GinConf struct {
	Mode string
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ganache_client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	sepolia_client, err := ethclient.Dial(os.Getenv("SEPOLIA_URL"))
	if err != nil {
		log.Fatal(err)
	}

	var ganache_Addresses []string
	// Ganache_Address from 0 to 9, os.GetEnv("GANACHE_ADDRESS0") to os.GetEnv("GANACHE_ADDRESS9)
	for i := 0; i < 10; i++ {
		ganache_Addresses = append(ganache_Addresses, os.Getenv("GANACHE_ADDRESS"+strconv.Itoa(i)))
	}
	client := ethclient.Client{}
	if os.Getenv("TEST_NET") == "Sepolia" {
		client = *sepolia_client
	} else {
		client = *ganache_client
	}
	return &Config{
		Ganache: GanacheConf{
			Client:    ganache_client,
			Addresses: ganache_Addresses,
		},
		Sepolia: SepoliaConf{
			Client: sepolia_client,
		},
		Gin: GinConf{
			Mode: os.Getenv("GIN_MODE"),
		},
		Client: &client,
	}
}

func SetGanacheAddresses() {
	os.Getenv("GANACHE_ADDRESS0")
}
