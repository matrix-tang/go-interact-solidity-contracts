package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-interact-solidity-contracts/contracts"
)

func main() {
	//服务器地址
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("local dial err", err)
		return
	}
	defer conn.Close()

	//创建合约对象
	dead, err := contracts.NewDeadline(common.HexToAddress("0x7287360507f0450577471F4987c2F26014B3729D"), conn)
	if err != nil {
		fmt.Println("new dead error", err)
	}
	result, err := dead.Show(&bind.CallOpts{})
	if err != nil {
		return
	}
	fmt.Println(result)

	bsc, err := ethclient.Dial("https://bsc-dataseed2.binance.org:443")
	if err != nil {
		fmt.Println("bsc dial err", err)
	}
	defer bsc.Close()
	router, err := contracts.NewPancakeRouter(common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"), bsc)
	if err != nil {
		fmt.Println("new pancake router error", err)
	}
	weth, err := router.WETH(&bind.CallOpts{})
	if err != nil {
		return
	}

	fmt.Println(weth)
}
