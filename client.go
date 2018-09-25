package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/gin-gonic/gin"

    bike "./contracts/bike"
    bikeCoin "./contracts/bike_coin"
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
        if err != nil {
            c.JSON(400, gin.H{
                "error": err,
            })    
        }
        c.JSON(200, gin.H{
            "value": rentalFee,
        })
    })

    r.GET("/rentalTimeInMinutes/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeContractInstance(hexAddress, client)
        rentalTimeInMinutes, err := instance.RentalTimeInMinutes(nil)
        if err != nil {
            c.JSON(400, gin.H{
                "error": err,
            })    
        }
        c.JSON(200, gin.H{
            "value": rentalTimeInMinutes,
        })
    })

    r.GET("/symbol/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        symbol, err := instance.Symbol(nil)
        if err != nil {
            c.JSON(400, gin.H{
                "error": err,
            })    
        }
        c.JSON(200, gin.H{
            "value": symbol,
        })
    })

    r.GET("/decimals/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        decimals, err := instance.Decimals(nil)
        if err != nil {
            c.JSON(400, gin.H{
                "error": err,
            })    
        }
        c.JSON(200, gin.H{
            "value": decimals,
        })
    })

    r.GET("/name/:address", func(c *gin.Context) {
        hexAddress := c.Param("address")
        instance := getBikeCoinContractInstance(hexAddress, client)
        name, err := instance.Name(nil)
        if err != nil {
            c.JSON(400, gin.H{
                "error": err,
            })    
        }
        c.JSON(200, gin.H{
            "value": name,
        })
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}

func getBikeCoinContractInstance(hexAddress string, client *ethclient.Client) *bikeCoin.BikeCoin {
    address := common.HexToAddress(hexAddress)
    instance, err := bikeCoin.NewBikeCoin(address, client)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("bike coin contract is loaded")
    return instance
}

func getBikeContractInstance(hexAddress string, client *ethclient.Client) *bike.Bike {
    address := common.HexToAddress(hexAddress)
    instance, err := bike.NewBike(address, client)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("bike contract is loaded")
    return instance
}