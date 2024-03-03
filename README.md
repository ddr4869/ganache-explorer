- 솔리디티로 만든 스마트 컨트랙트 개ㅐ발시 배포/테스트 환경을 제공한다
- NodeJs 위에서 작동하며, npm으로 설치한다.
- 설치
    - `sudo npm install -g truffle`
- 시작
    - truffle init
    
    ```go
    ieungyu@ieungyuui-MacBookPro truffle % tree
    .
    ├── contracts
    ├── migrations
    ├── test
    └── truffle-config.js
    ```
    
- openzeppelin 설치
    - `npm init -y`
    - `sudo npm install @openzeppelin/contracts`
    
- solidity 코드 작성
    
    ```solidity
    // SPDX-License-Identifier: MIT
    pragma solidity ^0.8.0;
    
    import "**@openzeppelin/contracts/token/ERC20/ERC20.sol**";
    
    contract TomToken is ERC20 {
        constructor() ERC20("QuickNode Coin", "QKC") {
            _mint(msg.sender, 10000 * 10 ** decimals());
        }
    }
    ```
    
    - **@openzeppelin/contracts/token/ERC20/ERC20.sol 부분은** openzeppelin을 설치 하여야 사용할 수 있음
- 
- 테스트 스크립트
    - test 폴더 안에 컨트랙트 테스트 코드를 작성할 수 있다, 여기선 skip

- 컨트랙트 컴파일
    - `sudo truffle compile`
    - build 폴더가 생성되었고, 구조는 다음과 같다.
        - abi, byte 등의 내용이 포함되었음
    
    ```solidity
    ieungyu@ieungyuui-MacBookPro build % tree
    .
    └── contracts
        ├── Context.json
        ├── ERC20.json
        ├── IERC1155Errors.json
        ├── IERC20.json
        ├── IERC20Errors.json
        ├── IERC20Metadata.json
        ├── IERC721Errors.json
        └── TokToken.json
    
    ```
    

- 배포 스크립트
    - migrations 폴더 안에 자바스크립트 파일로 배포 스크립트를 작성할 수 있다.
    - `1_initial_migration.js`
        - 파일명은 [번호]_[내용]_[컨트랙트이름].js 와 같은 형식으로 만든다.
    
    ```jsx
    const tomtoken = artifacts.require("TomToken"); // 컨트랙트 정보를 얻어오는 코드
    
    module.exports = function(deployer) {
        deployer.deploy(tomtoken);
    }
    ```
    
    - `sudo truffle migration`
    
    > 현재 openzeppelin 5.0.x 버전의 경우 solc 0.8.20 이상 버전부터 사용해야 하는데, 이 경우 이더리움 메인넷이 아닌 다른 네트워크의 경우 배포에서 에러가 발생한다.
    → `"@openzeppelin/contracts": "^4.9.0” ,` 
    -> `pragma solidity ^0.8.18;` 으로 배포하면 에러 발생하지 않음
    > 
    
    ```jsx
    ieungyu@ieungyuui-MacBookPro truffle % sudo truffle migration
    
    Compiling your contracts...
    ===========================
    > Everything is up to date, there is nothing to compile.
    
    Starting migrations...
    ======================
    > Network name:    'development'
    > Network id:      1709392599038
    > Block gas limit: 6721975 (0x6691b7)
    
    1_initial_migration.js
    ======================
    
       Deploying 'TomToken'
       --------------------
       > transaction hash:    0xe22322533b2e00d8b2048aa14a3839546a69f3ce67101d3e64a847c89057e9b5
       > Blocks: 0            Seconds: 0
       > contract address:    0x5e5c09a8eC3287c82153274a78178a748e3cC8a5
       > block number:        3
       > block timestamp:     1709394173
       > account:             0x2F8F42ccF8Ef6110C87156641104e880ca8fFE69
       > balance:             99.9666128
       > gas used:            1172676 (0x11e4c4)
       > gas price:           20 gwei
       > value sent:          0 ETH
       > total cost:          0.02345352 ETH
    
       > Saving artifacts
       -------------------------------------
       > Total cost:          0.02345352 ETH
    
    2_deploy_tomtoken.js
    ====================
    
       Replacing 'TomToken'
       --------------------
       > transaction hash:    0x49586c90e97006c77c62b20f9ed3be1eca54c9c1b8e901dd91c6ab474dd3607b
       > Blocks: 0            Seconds: 0
       > contract address:    0x95e6977eeFe64824e925c0682dAbBC29aFA62a29
       > block number:        4
       > block timestamp:     1709394173
       > account:             0x2F8F42ccF8Ef6110C87156641104e880ca8fFE69
       > balance:             99.94315928
       > gas used:            1172676 (0x11e4c4)
       > gas price:           20 gwei
       > value sent:          0 ETH
       > total cost:          0.02345352 ETH
    
       > Saving artifacts
       -------------------------------------
       > Total cost:          0.02345352 ETH
    
    Summary
    =======
    > Total deployments:   2
    > Final cost:          0.04690704 ETH
    
    ```
