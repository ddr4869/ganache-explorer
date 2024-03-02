package internal

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"

	"github.com/ddr4869/ether-go/internal/dto"
	"github.com/ddr4869/ether-go/solidity/store"

	token "github.com/ddr4869/ether-go/solidity/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func (s *Server) DeployStoreContract(c *gin.Context) {
	req := c.MustGet("req").(dto.DeployContractRequest)
	nonce, privateKey, err := s.GetNonceAndPKfromKeyString(req.From_private)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get nonce")
		return
	}

	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Gas suggestion failed")
		return
	}

	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get chainID")
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to bind auth")
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units -> 0.003 ether
	auth.GasPrice = gasPrice

	version := "1.0"
	address, tx, instance, err := store.DeployStore(auth, s.config.Client, version)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to deploy contract")
		return
	}
	_ = instance
	c.JSON(http.StatusOK, dto.DeployContractResponse{
		Address: address.Hex(),
		Hash:    tx.Hash(),
	})
}

// load contract
func (s *Server) LoadStoreContract(c *gin.Context) {
	req := c.MustGet("req").(dto.LoadStoreContractRequest)
	contractAddress := common.HexToAddress(req.Contract_address)
	instance, err := store.NewStore(contractAddress, s.config.Client)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to create store")
		return
	}
	key := [32]byte{}
	copy(key[:], []byte(req.Key))
	result, err := instance.Items(nil, key)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get item")
		return
	}
	log.Printf("result : %v", string(result[:]))
}

// write contract
func (s *Server) WriteStoreContract(c *gin.Context) {
	req := c.MustGet("req").(dto.WriteContractRequest)
	nonce, privateKey, err := s.GetNonceAndPKfromKeyString(req.From_private)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get nonce")
		return
	}

	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Gas suggestion failed")
		return
	}
	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get chainID")
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to bind auth")
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(req.Contract_address)
	instance, err := store.NewStore(address, s.config.Client)
	if err != nil {
		log.Print(err)
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create store")
		return
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte(req.Key))
	copy(value[:], []byte(req.Value))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Print(err)
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to set item")
		return
	}
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Print(err)
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to get item")
		return
	}
	fmt.Println(string(result[:])) // "bar"
	c.JSON(http.StatusOK, dto.TransactionHashResponse{
		Hash: tx.Hash(),
	})
}

func (s *Server) ReadByteContract(c *gin.Context) {
	req := c.MustGet("req").(dto.ReadByteContractRequest)
	addresses := common.HexToAddress(req.Contract_address)
	bytecode, err := s.config.Client.CodeAt(context.Background(), addresses, nil) // nil is latest block
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get bytecode")
		return
	}
	c.JSON(http.StatusOK, dto.ReadByteContractResponse{
		Bytecode: hex.EncodeToString(bytecode),
	})
}

func (s *Server) DeployErc20Contract(c *gin.Context) {
	req := c.MustGet("req").(dto.DeployContractRequest)
	nonce, privateKey, err := s.GetNonceAndPKfromKeyString(req.From_private)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get nonce")
		return
	}

	gasPrice, err := s.config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Gas suggestion failed")
		return
	}

	chainID, err := s.config.Client.NetworkID(context.Background())
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to get chainID")
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to bind auth")
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units -> 0.003 ether
	auth.GasPrice = gasPrice

	version := "1.0"
	address, tx, instance, err := store.DeployStore(auth, s.config.Client, version)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to deploy contract")
		return
	}
	_ = instance
	c.JSON(http.StatusOK, dto.DeployContractResponse{
		Address: address.Hex(),
		Hash:    tx.Hash(),
	})
}

func (s *Server) ReadErc20Contract(c *gin.Context) {
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, s.config.Client)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create token 1")
		return
	}

	address := common.HexToAddress("0x29fF15c2bF1a5Ac3F528A0E0920c9886b41E7245")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create token 2")
		return
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create token 3")
		return
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create token 4")
		return
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to create token 5")
		return
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
