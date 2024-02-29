package internal

import (
	"log"
	"net/http"
	"os"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/api")

	// account
	api.GET("/account/balance", s.GetBalanceValid, s.GetBalance)
	api.GET("/account/balance/pending", s.GetBalanceValid, s.GetPendingBalance)

	// block
	api.GET("/block/header", s.GetBlockByNumberValid, s.GetBlockByNumber)
	api.GET("/block/header/latest", s.GetLatestBlockNumber)
	// transaction
	api.GET("/transaction", s.GetTransaction)
}

func (s *Server) Start() error {

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: s.router,
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	return nil
}
