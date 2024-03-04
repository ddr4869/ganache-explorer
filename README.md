# GANACHE or SEPOLIA go clinet 

## select Ganache or Sepolia network

```python
# --> select this
# TEST_NET=GANACHE
# TEST_NET=Sepolia

# Gin config {test release debug}
Gin_MODE=test

# Server config
SERVER_HOST=127.0.0.1
SERVER_PORT=8081

# Sepolia config
SEPOLIA_URL=wss://sepolia.infura.io/ws/v3/{{API_KEY}}

# Ganache config 
GANACHE_ADDRESS0=0xdb7C688d060Fa51B3eaE1f2A6ce4B5ca86C3C7E9
GANACHE_ADDRESS1=0x4F7baC91f6D6bb8ac54c1168Af68596Be0339D69
GANACHE_ADDRESS2=0x0fa725B06CB235D23270b2433a8dC78B1A51Ad47
GANACHE_ADDRESS3=0x78DF26eD4FE5c755818C7017dc6454D6D44876D2
GANACHE_ADDRESS4=0xa9fA6E44d08dFEDC32bf450B867D57D1acfd9ee1
GANACHE_ADDRESS5=0x2a11b196aF748b82716cB2B3a029BFc8AECb1829
GANACHE_ADDRESS6=0x3034767AA90Ac6DD2FcF99ae375825Ce5DA583CB
GANACHE_ADDRESS7=0xd19120090Ba3d84E69D27a559Fbe802faF3813C9
GANACHE_ADDRESS8=0xE91dad6d47168EC042E39D0659Cf13bb5345883F
GANACHE_ADDRESS9=0x3b90CD2eF2F65F569B59B32e1a4939D1c91DF55e
```

## Abigen

- go를 이용하여 이더리움과 쉽게 상호작용 할 수 있는 ABI 바인딩 생성기
- abi를 통해 go 패키지를 만든다.
    
    `go install github.com/ethereum/go-ethereum/cmd/abigen@latest`
    

## Solidity 소스 파일로 abi → go 파일생성 및 배포

require

- solc : `sudo npm install -g solc`
- go abigen :
    - `go get -u github.com/ethereum/go-ethereum`
    - `cd ~/go/pkg/mod/github.com/ethereum/go-ethereum@v1.13.14`
        - go-ethereum의 패키지 경로
    - `make devtools`  → make 파일로 abigen 설치

1. solidity 코드 생성
    
    ```solidity
    pragma solidity ^0.8.24;
    
    contract Store {
      event ItemSet(bytes32 key, bytes32 value);
    
      string public version;
      mapping (bytes32 => bytes32) public items;
    
      constructor(string memory _version) public {
        version = _version;
      }
    
      function setItem(bytes32 key, bytes32 value) external {
        items[key] = value;
        emit ItemSet(key, value);
      }
    }
    ```
    
2. solc로 abi, bin 파일 생성
    - `solc --abi Store.sol -o Store`
        - Store/Store.abi에 파일 생성
    - `solc --bin Store.sol -o Store`
        - Store/Store.bin에 파일 생성
3. `abigen` 으로, abi → go 파일로 변환
    - `abigen --abi=Store/Store.abi --pkg=store --out=**Store.go**`
4. abi로 스마트 컨트랙트 배포
    - 컨트랙트를 배포하려면, 컴파일된 바이트코드 형식의 몇가지 추가 정보가 필요하다.
    - `abigen --bin=Store/Store.bin --abi=Store/Store.abi --pkg=store -out=**Store.go**`

<aside>
💡 3번으로 얻은 `**Store.go`** 파일과, 4번으로 얻은 `**Store.go` 파일은 `DeployStorage` 기능의 유무 차이가 있다.**

</aside>

```solidity
// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}
```

- `DeployStorage()` 함수를 사용하여 go 애플리케이션에서 이더리움 테스트넷에 컨트랙트를 배포할 수 있다.

### DeployStorage를 통한 배포

- 컨트랙트를 배포하려면, 다음이 필요하다.
    - 이더리움 테스트넷에 연결된 실행중인 Geth 노드
    - 컨트랙트 배포 및 상호작용에 필요한 Gas만큼 충분한 이더가 적립된 account
1. transfer 때처럼 nonce, private key, gas price, chainc id등을 불러온 뒤
2. `NewKeyedTransactorWithChainID` 로 새로운 auth(key 트랜잭션 옵션) 생성
3. auth에 표준 트랜잭션 옵션을 설정해준 후
4. abigen으로 생성한 go 파일을 import 하여
5. `DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string)`

```go

func (s *Server) DeployContract(c *gin.Context) {
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, s.config.Client, input)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "Failed to deploy contract")
		return
	}

	c.JSON(http.StatusOK, dto.TransactionHashResponse{
		Hash: tx.Hash(),
	})
	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}
```

### Contracd에 write

1. Deploy 할때처럼, auth 생성
2. Abigen 으로 생성한 go 파일의, NewStore (go package가 store라고 가정)
    
    `instance, err := store.NewStore(address, s.config.Client)`
    
3. 컨트랙트에 정의된대로, 메서드 실행

```go
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "Failed to set item")
		return
	}
```
