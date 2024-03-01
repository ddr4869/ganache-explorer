package internal

import (
	"context"
	"log"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

// block_number를 받지 않으면, 가장 최신 블록을 가져온다.
func (s *Server) GetBlockByNumber(c *gin.Context) {
	req := c.MustGet("req").(dto.GetBlockHeaderByNumberRequest)
	blockNumber, err := s.CheckBlockNumber(req.BlockNumber)
	if err != nil {
		return
	}
	block, err := s.config.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, dto.ConvertBlockToResponse(*block))
}

func (s *Server) CheckBlockNumber(BlockNumber int) (*big.Int, error) {
	var blockNumber *big.Int
	if BlockNumber == 0 {
		header, err := s.config.Client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		blockNumber = header.Number
	} else {
		blockNumber = big.NewInt(int64(BlockNumber))
	}
	return blockNumber, nil
}

func (s *Server) SubscribeBlock(c *gin.Context) {

	headers := make(chan *types.Header) // 최신 블록 헤더를 받을 chan 생성
	sub, err := s.config.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to subscribe new head")
		return
	}

	for {
		select {
		case err := <-sub.Err():
			dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Received error from subscription")
			return
		case header := <-headers:
			log.Printf("block hash: %s", header.Hash().Hex())
			log.Printf("block number: %s", header.Number.String())
			block, err := s.config.Client.BlockByHash(context.Background(), header.Hash())
			b, err := s.entClient.Block.Create().
				SetNumber(int(block.Number().Int64())).
				SetGasLimit(int(block.GasLimit())).
				SetGasUsed(int(block.GasUsed())).
				SetDifficulty(int(block.Difficulty().Int64())).
				SetTime(block.Time()).
				SetNumberU64(block.NumberU64()).
				SetMixDigest(block.MixDigest().Hex()).
				SetNonce(int(block.Nonce())).
				SetCoinbase(block.Coinbase().Hex()).
				SetRoot(block.Root().Hex()).
				SetParentHash(block.ParentHash().Hex()).
				SetTxHash(block.TxHash().Hex()).
				SetReceiptHash(block.ReceiptHash().Hex()).
				SetUncleHash(block.UncleHash().Hex()).
				SetTxCount(block.Transactions().Len()).
				Save(context.Background())
			if err != nil {
				dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to save block to db")
				return
			}
			log.Printf("block id: %d", b.ID)
			for _, tx := range block.Transactions() {
				// if i > 1 {
				// 	break
				// }
				// log.Printf("tx hash: %s", tx.Hash().Hex())
				// log.Printf("tx nonce: %d", tx.Nonce())
				// log.Printf("tx gas: %d", tx.Gas())
				// log.Printf("tx gas price: %s", tx.GasPrice().String())
				// log.Printf("tx gas tip cap: %s", tx.GasTipCap().String())
				// log.Printf("tx gas fee cap: %s", tx.GasFeeCap().String())
				// log.Printf("tx value: %s", tx.Value().String())
				// log.Printf("tx to: %s", tx.To().Hex())
				// log.Printf("tx data: %s", string(tx.Data()))
				// log.Printf("tx type: %d", tx.Type())
				// log.Printf("tx chain id: %d", tx.ChainId().Int64())

				_, err := s.entClient.Transaction.Create().
					SetType(int(tx.Type())).
					SetChainID(int(tx.ChainId().Int64())).
					SetNonce(int(tx.Nonce())).
					SetTo(tx.To().Hex()).
					SetGas(int(tx.Gas())).
					SetGasPrice(tx.GasPrice().String()).
					SetGasTipCap(tx.GasTipCap().String()).
					SetGasFeeCap(tx.GasFeeCap().String()).
					SetValue(tx.Value().String()).
					//SetData(string(tx.Data())).
					SetHash(tx.Hash().Hex()).
					Save(context.Background())
				if err != nil {
					dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to save transaction to db")
					return
				}
			}
			if err != nil {
				dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get block by hash")
				return
			}
			log.Printf("block tx len: %d", block.Transactions().Len()) // 3
		}
	}
}
