package main

import (
    "log"

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

    r.Run() // listen and serve on 0.0.0.0:8080
}





