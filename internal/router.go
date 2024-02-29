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
	api.GET("/block", s.GetBlockByNumberValid, s.GetBlockByNumber)

	// transaction
	api.GET("/transaction", s.GetTransactionByHashValid, s.GetTransactionByHash)
	api.GET("/transactions", s.GetTransactionsValid, s.GetTransactions)
	api.GET("/transactions/hashes", s.GetTransactionHashesValid, s.GetTransactionHashes)
	api.GET("/transactions/receipt", s.GetTransactionReceiptValid, s.GetTransactionReceipt)

	// transfer
	api.POST("/transfer", s.TransferValid, s.Transfer)
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
