package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://kovan.infura.io/v3/550dd3ed604f4342aaf4aa938937a274"

func main() {
	// Creating the keystore/ wallet with standard parameters
	//ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// Creating a new account using password = "password"
	//_, err := ks.NewAccount("password")
	//if err != nil {

	//	log.Fatal(err)
	//}

	//_, err = ks.NewAccount("password")
	//if err != nil {

	//	log.Fatal(err)
	//}

	//"d5e52477cd5a9f940e98fac35ad24320b685fa77"
	//"f013b50694c0e118016826b7b57dbcc4ee5d88d1"

	// Setup our ETH client

	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Getting two addreses
	a1 := common.HexToAddress("d5e52477cd5a9f940e98fac35ad24320b685fa77")
	a2 := common.HexToAddress("f013b50694c0e118016826b7b57dbcc4ee5d88d1")

	// Getting the Balance account 1
	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Getting the Balance account 2
	b2, err := client.BalanceAt(context.Background(), a2, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance of account1 :", b1)
	fmt.Println("Balance of account2:", b2)

	// Getting the nonce number for transaccion en account 1
	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 10000000000000000 wei - sendeing 0.01 ether
	amount := big.NewInt(100000000000000000)

	// Getting suggested gas price of the Network/Kovan in this case
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Creating new TX to send
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Reading from the encrypted wallet
	b, err := ioutil.ReadFile("wallet/UTC--2022-07-03T01-56-55.753964400Z--d5e52477cd5a9f940e98fac35ad24320b685fa77")
	if err != nil {
		log.Fatal(err)
	}
	//  Getting private key Decrypted using password = "password"
	key, err := keystore.DecryptKey(b, "password")

	if err != nil {
		log.Fatal(err)
	}
	// Signing the transaction using the private key
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	// Sending the transaction to the network
	err = client.SendTransaction(context.Background(), tx)

	if err != nil {
		log.Fatal(err)
	}
	// Printing the hash of the transaction
	fmt.Println("tx sent: ", tx.Hash().Hex())

}
