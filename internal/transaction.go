package internal

import (
	"context"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTransactionByHash(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionByHashRequest)
	txHash := common.HexToHash(req.Hash)
	tx, _, err := s.config.Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetTransactionByHashResponse{
		Transaction: dto.ConvertTransaction(tx),
	})
}

func (s *Server) GetTransactions(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionRequest)
	blockNumber := int64(req.BlockNumber)

	block, err := s.config.Client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, dto.GetTransactionResponse{
		Transactions: dto.ConvertTransactions(block.Transactions()),
		Tx_length:    len(block.Transactions()),
	})
}

func (s *Server) GetTransactionReceipt(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionReceiptRequest)
	blockNumber, err := s.CheckBlockNumber(req.BlockNumber)
	if err != nil {
		log.Fatal(err)
		return
	}
	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]types.Receipt, 0)

	if req.Hash != "" {
		receipt, err := s.config.Client.TransactionReceipt(context.Background(), common.HexToHash(req.Hash))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *receipt)
		c.JSON(http.StatusOK, dto.GetTransactionReceiptResponse{
			Receipts: result,
		})
		return
	}

	for _, tx := range block.Transactions() {
		// chainID, err := s.config.Client.NetworkID(context.Background())
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
		// 	fmt.Println(from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		// }
		receipt, err := s.config.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *receipt)
	}
	log.Printf("Receipts len !: %v\n", len(result))
	c.JSON(http.StatusOK, dto.GetTransactionReceiptResponse{
		Receipts: result,
	})
}

func (s *Server) GetTransactionHashes(c *gin.Context) {
	req := c.MustGet("req").(dto.GetTransactionHashesRequest)
	blockNumber := int64(req.BlockNumber)

	block, err := s.config.Client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal(err)
	}

	hashes := make([]common.Hash, 0)
	for _, v := range block.Transactions() {
		hashes = append(hashes, v.Hash())
	}

	c.JSON(http.StatusOK, dto.GetTransactionHashesResponse{
		Hashes:    hashes,
		Tx_length: len(hashes),
	})
}
