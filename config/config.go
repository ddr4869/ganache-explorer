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
	Gin     GinConf
}

type GanacheConf struct {
	Client    *ethclient.Client
	Addresses []string
}

type GinConf struct {
	Mode string
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	var ganache_Addresses []string
	// Ganache_Address from 0 to 9, os.GetEnv("GANACHE_ADDRESS0") to os.GetEnv("GANACHE_ADDRESS9)
	for i := 0; i < 10; i++ {
		ganache_Addresses = append(ganache_Addresses, os.Getenv("GANACHE_ADDRESS"+strconv.Itoa(i)))
	}

	return &Config{
		Ganache: GanacheConf{
			Client:    client,
			Addresses: ganache_Addresses,
		},
		Gin: GinConf{
			Mode: os.Getenv("GIN_MODE"),
		},
	}
}

func SetGanacheAddresses() {
	os.Getenv("GANACHE_ADDRESS0")
}
