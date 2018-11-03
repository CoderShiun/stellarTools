package main

import (
	"context"
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"io/ioutil"
	"log"
	"net/http"
)

var mySeed, myAddress, myAccount string
var yourSeed, yourAddress, yourAccount string

func main() {
	/**
	Get the Seed and Key for myAccount
	 */
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mySeed: ", pair.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	log.Println("myAddress: ", pair.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB
	log.Println()

	mySeed = pair.Seed()
	myAddress = pair.Address()

	/*
	####################################
	 */
	/**
   Get the Seed and Key for yourAccount
	*/
	pair2, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("yourSeed: ", pair2.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	log.Println("yourAddress: ", pair2.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB
	log.Println()

	yourSeed = pair.Seed()
	yourAddress = pair.Address()

	/*#############################################################################################*/
	/**
	Create myAccount
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp, err := http.Get("https://friendbot.stellar.org/?addr=" + myAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp.Body.Close()
	body, err := ioutil.ReadAll(rresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("myAccount ist: ", string(body))
	myAccount = string(body)
	fmt.Println()
	/*#############################################################################################*/

	/**
	Create yourAccount
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp2, err := http.Get("https://friendbot.stellar.org/?addr=" + yourAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp2.Body.Close()
	body2, err := ioutil.ReadAll(rresp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("yourAccount ist: ", string(body2))
	yourAccount = string(body2)
	fmt.Println()
	/*
	##############################################################
	 */
	/**
	Check the Balance for myAccount
	 */
	account, err := horizon.DefaultTestNetClient.LoadAccount(myAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for myAccount:", myAddress)

	for _, balance := range account.Balances {
		log.Println("myAccount balance: ", balance)
		//fmt.Printf("Balance Type = %T", &balance)
	}
	fmt.Println()

	/*
	##############################################################
	 */
	/**
	Check the Balance for yourAccount
	 */
	account2, err := horizon.DefaultTestNetClient.LoadAccount(yourAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for yourAccount:", yourAddress)

	for _, balance := range account2.Balances {
		log.Println("yourAccount balance: ", balance)
		//fmt.Printf("Balance Type = %T", &balance)
	}
	fmt.Println()

	/*#############################################################################################*/
	//	send Transaction
	//source := "SCZANGBA5YHTNYVVV4C3U252E2B6P6F5T3U6MM63WBSBZATAQI3EBTQ4"
	//destination := "GCWMDLBZV5NZQHUJ3Z4OEHRLAHSLUUUIM3BITFGQVJMSYKJ5DGR4UEGM"

	// Make sure destination account exists
	if _, err := horizon.DefaultTestNetClient.LoadAccount(myAddress); err != nil {
		panic(err)
	}

	passphrase := network.TestNetworkPassphrase
	fmt.Println(passphrase)

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{mySeed},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{yourAddress},
			build.NativeAmount{"10"},
		),
	)

	if err != nil {
		panic(err)
	}

	// Sign the transaction to prove you are actually the person sending it.
	txe, err := tx.Sign(mySeed)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)

	/*################################################################################*/
	//Receive Transaction
	//address := "SACHH2LKB5WYPZL4KE4SCAWF443R5KPOIWO3OT3ZEXS6Y5D35TYJ62WC"
	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	esrr := horizon.DefaultTestNetClient.StreamPayments(ctx, yourAddress, &cursor, func(payment horizon.Payment) {
		fmt.Println("Payment type", payment.Type)
		fmt.Println("Payment Paging Token", payment.PagingToken)
		fmt.Println("Payment From", payment.From)
		fmt.Println("Payment To", payment.To)
		fmt.Println("Payment Asset Type", payment.AssetType)
		fmt.Println("Payment Asset Code", payment.AssetCode)
		fmt.Println("Payment Asset Issuer", payment.AssetIssuer)
		fmt.Println("Payment Amount", payment.Amount)
		fmt.Println("Payment Memo Type", payment.Memo.Type)
		fmt.Println("Payment Memo", payment.Memo.Value)
	})
	//fmt.Println(esrr)

	if err != nil {
		panic(esrr)
	}

	/*
	############################################
	 */
	/*Check the Balance of myAccount again*/
	account3, err := horizon.DefaultTestNetClient.LoadAccount(myAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:", myAccount)

	for _, balance := range account3.Balances {
		log.Println("The balance: ", balance)
		fmt.Printf("Balance Type = %T", &balance)
	}

	/*
	############################################
	 */
	/*Check the Balance of myAccount again*/
	account4, err := horizon.DefaultTestNetClient.LoadAccount(myAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:", yourAccount)

	for _, balance := range account4.Balances {
		log.Println("The balance: ", balance)
		fmt.Printf("Balance Type = %T", &balance)
	}

}
