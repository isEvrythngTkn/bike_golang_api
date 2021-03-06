# Bike Contract API

This is an API designed to interact with the Solidity contracts built for the Bike Challenge.

## Requirements: 
- github.com/ethereum/go-ethereum/ethclient
- github.com/gin-gonic/gin

You'll need this repo too: https://github.com/isEvrythngTkn/bike-challenge \
And you'll need to migrate those contracts onto your local Ethereum blockchain (ganache) and ensure that you have the addresses to those contracts.

## Instructions

Install dependencies: 
`go get github.com/ethereum/go-ethereum/ethclient && go get github.com/gin-gonic/gin`

To start the server:
`go run client.go response.go contract.go`

## Routes

Bike Rental Contract\
`/rentalFee/[address]`\
`/rentalTimeInMinutes/[address]`

Bike Coin Contract\
`/symbol/[tokenContractAddres]`\
`/decimals/[tokenContractAddress]`\
`/name/[tokenContractAddress]`\
`/viewBalance/[tokenContractAddress]/[accountAddress]`

Crowdsale Contract\
`/purchaseTokens/[crowdSaleAddress]/[privateKey]`

The server runs on port 8080 by default. An example url is:
http://localhost:8080/viewBalance/0x4b4eca1e3672a29125befbc9fd6469d2f49c4294/0x1ba2cbe1bdae495c72563590410413446d208596

## Sources
** A lot of the code for this is based off https://goethereumbook.org/ **
