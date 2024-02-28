package main

import (
	"github.com/ddr4869/ether-go/config"
	"github.com/ddr4869/ether-go/internal"
)

func main() {
	cfg := config.Init()
	router := internal.NewRestController(cfg)
	router.Start()
}
