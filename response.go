package main

import (
    "log"
    "math/big"

    "github.com/gin-gonic/gin"
)

func errorResponse(err error, c *gin.Context) {
    if err != nil {
        c.JSON(400, gin.H{
            "error": err,
        })    
    }
}

func successResponseString(key string, value string, c *gin.Context) {
    c.JSON(200, gin.H{
        key: value,
    })
}

func successResponseBigInt(key string, value *big.Int, c *gin.Context) {
    c.JSON(200, gin.H{
        key: value,
    })
}

func successResponseUint8(key string, value uint8, c *gin.Context) {
    c.JSON(200, gin.H{
        key: value,
    })
}

func errorHandler(err error) {
    if err != nil {
        log.Fatal(err)
    }
}