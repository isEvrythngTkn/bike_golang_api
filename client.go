package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/gin-gonic/gin"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")
    _ = client

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}