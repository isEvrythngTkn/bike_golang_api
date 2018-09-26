package main

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    bike "./contracts/bike"
    bikeCoin "./contracts/bike_coin"
)

func getAddress(hexAddress string) common.Address {
    return common.HexToAddress(hexAddress)
}

func getBikeCoinContractInstance(hexAddress string, client *ethclient.Client) *bikeCoin.BikeCoin {
    address := getAddress(hexAddress)
    instance, err := bikeCoin.NewBikeCoin(address, client)
    errorHandler(err);
    return instance
}

func getBikeContractInstance(hexAddress string, client *ethclient.Client) *bike.Bike {
    address := getAddress(hexAddress)
    instance, err := bike.NewBike(address, client)
    errorHandler(err);
    return instance
}