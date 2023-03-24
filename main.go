package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-interact-solidity-contracts/contracts"
	"math/big"
)

const BscRpcUrl = "https://bsc-dataseed2.binance.org:443"
const PriKey = ""
const PancakeRouterAddress = "0x10ED43C718714eb63d5aA57B78B54704E256024E"

func main() {
	//服务器地址
	/*conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("local dial err", err)
		return
	}
	defer conn.Close()

	/// deadline
	dead, err := contracts.NewDeadline(common.HexToAddress("0x7287360507f0450577471F4987c2F26014B3729D"), conn)
	if err != nil {
		fmt.Println("new dead error", err)
	}
	result, err := dead.Show(&bind.CallOpts{})
	if err != nil {
		return
	}
	fmt.Println(result)

	/// storage
	storage, err := contracts.NewStorage(common.HexToAddress("0xA4981D26DAE32ECE1109e6Be9eb5612B16596f82"), conn)
	if err != nil {
		fmt.Println("new store error", err)
	}

	privateKey, _ := crypto.HexToECDSA("80e1e5e75bd00b049db7c8d7de3d4668a6115699150653c665a07e6cada71848")
	chainId, _ := conn.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	tx, err := storage.Store(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, new(big.Int).SetInt64(2e18))
	if err != nil {
		fmt.Println("store error", err)
	}
	fmt.Println(tx.Hash())*/

	BscSwapExample()
}

func BscSwapExample() {
	/// bsc
	bsc, err := ethclient.Dial(BscRpcUrl)
	if err != nil {
		fmt.Println("bsc dial err", err)
	}
	defer bsc.Close()

	router, err := contracts.NewPancakeRouter(common.HexToAddress(PancakeRouterAddress), bsc)
	if err != nil {
		fmt.Println("new pancake router error", err)
	}

	privateKeyBsc, _ := crypto.HexToECDSA(PriKey)
	chainIdBsc, _ := bsc.ChainID(context.Background())
	authBsc, _ := bind.NewKeyedTransactorWithChainID(privateKeyBsc, chainIdBsc)

	// wbnb 授权给router地址
	/*wbnb, err := contracts.NewWbnb(common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"), bsc)
	wtx, err := wbnb.Approve(&bind.TransactOpts{
		From:   authBsc.From,
		Signer: authBsc.Signer,
	}, common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"), new(big.Int).SetInt64(2e18))
	if err != nil {
		fmt.Println("wbnb approve error: ", err)
	}
	fmt.Println("wbnb approve tx: ", wtx.Hash())*/
	// cake 授权给router地址
	/*cake, err := contracts.NewCake(common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"), bsc)
	ctx, err := cake.Approve(&bind.TransactOpts{
		From:   authBsc.From,
		Signer: authBsc.Signer,
	}, common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"), new(big.Int).SetInt64(2e18))
	if err != nil {
		fmt.Println("cake approve error: ", err)
	}
	fmt.Println("cake approve tx: ", ctx.Hash())*/

	path0 := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	path1 := common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82")
	to := common.HexToAddress("0xD21735fb3D95320d72a434A24b9f60b040c442D6")
	tx1, err := router.SwapExactETHForTokens(&bind.TransactOpts{
		From:   authBsc.From,
		Signer: authBsc.Signer,
		Value:  new(big.Int).SetInt64(2e12),
	}, new(big.Int).SetInt64(2e12), []common.Address{path0, path1}, to, new(big.Int).SetInt64(1680510182))
	if err != nil {
		fmt.Println("swap wbnb to cake error", err)
	}
	fmt.Println("swap wbnb to cake tx: ", tx1.Hash())
}
