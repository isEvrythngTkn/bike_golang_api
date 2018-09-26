package main

import (
    "log"
    "context"
    "fmt"
    "strings"

    "encoding/hex"
    "math/big"
    "crypto/ecdsa"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/gin-gonic/gin"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/rentalFee/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeContractInstance(hexAddress, client)
        rentalFee, err := instance.RentalFee(nil)
        errorResponse(err, c)
        successResponseBigInt("rentalFee", rentalFee, c)
    })

    r.GET("/rentalTimeInMinutes/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeContractInstance(hexAddress, client)
        rentalTimeInMinutes, err := instance.RentalTimeInMinutes(nil)
        errorResponse(err, c)
        successResponseBigInt("rentalTimeInMinutes", rentalTimeInMinutes, c)
    })

    r.GET("/symbol/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        symbol, err := instance.Symbol(nil)
        errorResponse(err, c)
        successResponseString("symbol", symbol, c)
    })

    r.GET("/decimals/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        decimals, err := instance.Decimals(nil)
        errorResponse(err, c)
        successResponseUint8("decimals", decimals, c)
    })

    r.GET("/name/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        name, err := instance.Name(nil)
        errorResponse(err, c)
        successResponseString("name", name, c)
    })

    r.GET("/viewBalance/:tokenContractAddress/:accountAddress", func(c *gin.Context) {
        accountAddressString := c.Param("accountAddress")
        accountAddress := getAddress(accountAddressString)

        tokenContractAddress := c.Param("tokenContractAddress")
        instance := getBikeCoinContractInstance(tokenContractAddress, client)
        
        balance, err := instance.BalanceOf(&bind.CallOpts{}, accountAddress)
        errorResponse(err, c)
        successResponseBigInt("balance", balance, c)
    });

    r.POST("/purchaseTokens/:crowdSaleAddress/:privateKey", func(c *gin.Context) {
        crowdSaleAddress := c.Param("crowdSaleAddress")
        privateKeyHex := c.Param("privateKey")
        if strings.HasPrefix(privateKeyHex, "0x") {
            a := []rune(privateKeyHex)
            privateKeyHex = string(a[2:])
        }
        decoded, err := hex.DecodeString(privateKeyHex)
        if err != nil {
            //log.Fatal("here")
            log.Fatal(err)
        }

        privateKey, err := crypto.ToECDSA(decoded)
        if err != nil {
            log.Fatal(err)
        }

        publicKey := privateKey.Public()
        publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
        if !ok {
            log.Fatal("error casting public key to ECDSA")
        }

        fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
        nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
        if err != nil {
            log.Fatal(err)
        }

        value := big.NewInt(1000000000000000000) // in wei (1 eth)
        gasLimit := uint64(3000000)
        gasPrice, err := client.SuggestGasPrice(context.Background())
        if err != nil {
            log.Fatal(err)
        }

        var data []byte
        toAddress := getAddress(crowdSaleAddress)
        tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

        chainID, err := client.NetworkID(context.Background())
        if err != nil {
            log.Fatal(err)
        }

        signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
        if err != nil {
            log.Fatal(err)
        }

        err = client.SendTransaction(context.Background(), signedTx)
        if err != nil {
            log.Fatal(err)
        }

        transactionString := signedTx.Hash().Hex()
        fmt.Printf(transactionString)

        successResponseString("transaction", transactionString, c)
    });

    r.Run() // listen and serve on 0.0.0.0:8080
}





