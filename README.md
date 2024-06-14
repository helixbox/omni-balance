# Omni Balance
Omni Balance is an intelligent liquidity management tool designed to achieve seamless interoperability between decentralized exchanges (DEX), centralized exchanges (CEX), and bridging services. It integrates automated cross-chain operations to help users effortlessly balance their liquidity pools while ensuring token balances across different addresses are replenished.

## Table of Contents
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Configuration](#configuration)

### Features

* **Multi-address Monitoring**: Real-time monitoring of specific token balances in different addresses, ensuring continuous supervision.
* **Automatic Token Balance Adjustment**: Allows setting custom balance thresholds, and once below the threshold, it automatically adjusts through the best route, including transfers from DEX, CEX, and bridges.
* **Cross-chain Support**: Enables seamless cross-chain transactions, expanding liquidity options.
* **Address Isolation**: Allows separation of monitoring addresses and operational addresses.
* **Supports Most EVM-compatible Chains** (depending on the provider you choose).

#### Supported Providers

| name                                                                             | swap | cross-chain |
|----------------------------------------------------------------------------------|------|-------------|
| [UniSwap](https://uniswap.org/)                                                  | ✅    |             |
| [Gate.io](https://gate.io)                                                       | ✅    | ✅           |
| [Helix Bridge](https://helixbridge.app/)                                         | ✅    |             |
| [Darwinia Bridge](https://bridge.darwinia.network/)                              | ✅    |             |
| [OKX web3](https://www.okx.com/zh-hans/web3/build/docs/waas/dex-crosschain-swap) | ✅    | ✅           |


### Installation

#### Download Binary Files from GitHub Releases

1. Visit the [Releases](https://github.com/darwinia-network/omni-balance/releases) page.
2. Download the latest version of the binary file for your operating system.

#### Using Docker

1. Pull the Docker image: `docker pull ghcr.io/darwinia-network/omni-balance:latest`

#### Building from Source

1. Ensure Golang 1.22 or later is installed.
2. Clone the repository: `git clone https://github.com/darwinia-network/omni-balance && cd omni-balance`
3. Build the project: `go build ./cmd`

#### Using Golang Install

1. Run the command: `go install github.com/darwinia-network/omni-balance@latest`

### Usage

#### Supported Commands

```sh
COMMANDS:
   version, v  show version
   list        list supported providers and docs
   tasks       list supported tasks
   example     create a example config file
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:

    --conf value, -c value              (default: "./config.yaml")
          config file path

    --placeholder, -p                   (default: false)
          enable placeholder, you can use placeholder to replace private key, Example: Fill '{{privateKey}}' in config.yaml. Run with -p to enable placeholder: SERVER_PORT=:8080 omni-balance -c ./config.yaml -p. Waiting for 'waiting for placeholder...' log, send placeholder
          data according to the prompt.

    --port value                        (default: ":8080")
          When the placeholder parameter is set to true, you can specify and set the
          listening address of the HTTP server that receives the placeholder.

   MISC
    --help, -h   show help
```
#### Configuration
```
# Debug mode
debug: true
# Chains
chains:
    - # Chain id
      id: 1
      # Chain name
      name: etnereum
      # Native token name, if not set, use the 0x0000000000000000000000000000000000000000
      nativetoken: "ETH"
      # RPC endpoints
      rpc_endpoints:
        - https://api.tatum.io/v3/blockchain/node/ethereum-mainnet
        - https://ethereum-rpc.publicnode.com
      # Tokens
      tokens:
        - # Token name
          name: ETH
          # Token contract address
          contract_address: "0x0000000000000000000000000000000000000000"
          # Token decimals
          decimals: 18
        - # Token name
          name: RING
          # Token contract address
          contract_address: 0x9469D013805bFfB7D3DEBe5E7839237e535ec483
          # Token decimals
          decimals: 18
        - # Token name
          name: USDT
          # Token contract address
          contract_address: 0xdAC17F958D2ee523a2206206994597C13D831ec7
          # Token decimals
          decimals: 6
    - # Chain id
      id: 42161
      # Chain name
      name: arbitrum
      # Native token name, if not set, use the 0x0000000000000000000000000000000000000000
      nativetoken: "ETH"
      # RPC endpoints
      rpc_endpoints:
        - https://1rpc.io/arb
        - https://arbitrum.llamarpc.com
      # Tokens
      tokens:
        - # Token name
          name: ETH
          # Token contract address
          contract_address: "0x0000000000000000000000000000000000000000"
          # Token decimals
          decimals: 18
        - # Token name
          name: USDT
          # Token contract address
          contract_address: 0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9
          # Token decimals
          decimals: 6
        - # Token name
          name: RING
          # Token contract address
          contract_address: 0x9e523234D36973f9e38642886197D023C88e307e
          # Token decimals
          decimals: 18
# Source token used to buy other tokens
source_token:
    - # Token name
      name: USDT
      # Chain name
      chains:
        - etnereum
        - arbitrum
    - # Token name
      name: ETH
      # Chain name
      chains:
        - etnereum
        - arbitrum
# Liquidity providers
liquidity_providers:
    - # Type liquidity provider type
      type: CEX
      # LiquidityName liquidity provider name
      liquidity_name: gate.io
      # Config liquidity provider config, depend on the type
      config:
        key: <gate_api_key>
        secret: <gate_api_secret>
    - # Type liquidity provider type
      type: DEX
      # LiquidityName liquidity provider name
      liquidity_name: uniswap
    - # Type liquidity provider type
      type: Bridge
      # LiquidityName liquidity provider name
      liquidity_name: helixbridge
    - # Type liquidity provider type
      type: Bridge
      # LiquidityName liquidity provider name
      liquidity_name: darwinia-bridge
# Wallets need to rebalance
wallets:
    - # Wallet address
      address: 0x43Ef13E84D9992d1461a1f90CAc4653658CEA4FD
      # If set, when the 'address' balance is insufficient, the operator address will be used to rebalance
      operator: ""
      # Tokens to be monitored
      tokens:
        - # Token name
          name: ETH
          # The number of each rebalance
          amount: "1"
          # Threshold when the token balance is less than the threshold, the rebalance will be triggered
          threshold: "2"
          # The chains need to be monitored
          chains:
            - ethereum
      # Wallet private key. If operator is not empty, private_key is the operator's private key
      private_key: <wallet1_private_key>
    - # Wallet address
      address: 0x43Ef13E84D9992d1461a1f90CAc4653658CEA4FD
      # If set, when the 'address' balance is insufficient, the operator address will be used to rebalance
      operator: "0x178D8546C5f78e01133858958355B06EC3406A1A"
      # Tokens to be monitored
      tokens:
        - # Token name
          name: RING
          # The number of each rebalance
          amount: "20000"
          # Threshold when the token balance is less than the threshold, the rebalance will be triggered
          threshold: "10000"
          # The chains need to be monitored
          chains:
            - ethereum
            - arbitrum
      # Wallet private key. If operator is not empty, private_key is the operator's private key
      private_key: <wallet2_private_key>
# Database config, must set one of them
db:
    # MYSQL config
    MYSQL:
        host: ""
        port: ""
        user: ""
        password: ""
        database: ""
    # POSTGRESQL config
    POSTGRESQL:
        host: ""
        port: ""
        user: ""
        password: ""
        database: ""
    # SQLITE config
    SQLITE:
        path: /data/omni-balance.db
task_interval:
    cross_chain: 1m
    get_token_price_in_usdt: 1m
    monitor_wallet_balance: 1m
    rebalance: 1m
```

#### Run
##### Using Docker
1. Run the Docker image:
```
docker run -d --name omni-balance -p 8080:8080 \
    -v <you-host-path>:/data \
    --restart=always \
    --name omni-balance \
     omni-balance:latest -c /data/config.yaml -p
```
2. Check logs: `docker logs omni-balance`

##### Using command
1. Run the command: `omni-balance -c <your-config-path> -p`

##### Replacing Placeholders (Optional)
**For security reasons, if you do not want to write private keys in plain text in the configuration file, you can replace placeholders via a POST request**

1. Run the command:
```shell
curl -X POST http://localhost:8080 \
  -d '{"<wallet1_private_key>":"xxxx", "<wallet2_private_key>":"xxxx"}'
```
2. Clear history
```shell
history -c
```
