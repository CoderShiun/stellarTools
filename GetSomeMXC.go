package main

import (
	"fmt"
	"stellarTools/getFromKyeboard"
	"stellarTools/gettingAccountDetails"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/network"
)

var distributorPri = "SD64I4OCIUPT6UONUCOT7WVIVOVQNFJI6FLO2NIO5CUDBPXQMUUZSDJA"
var distributorPub = "GDZPSKJSFLW4HF77OJFYB3ANCG3CF2PHUZ3YRG5PHPTNF7ZLYGNV7EPV"

func main() {
	fmt.Println("Please enter your PublicKey: ")
	publicKey := getFromKyeboard.FromKeyboard()
	//buildingTransaction.SendMXC(publicKey)

	// Make sure destination account exists
	if _, err := horizon.DefaultTestNetClient.LoadAccount(publicKey); err != nil {
		panic(err)
	}

	passphrase := network.TestNetworkPassphrase
	fmt.Println(passphrase)

	/*tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{source},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{destination},
			//build.PayWith(MXC01,"100"),
			//build.Amount("50"),
			//build.CreditAmount{"MXC01",MXCissuer,"50"},
			build.NativeAmount{"50"},
		),
	)*/

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{  distributorPub},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{publicKey},
			//build.SourceAccount{"GAPYC3DNGCYC4TJYCVPSLC476WYVNDTDJ2XOKDJWITCF3XYDSDT2FLXL"},
			//build.Asset{Code:"MXC01",Issuer:"GAPYC3DNGCYC4TJYCVPSLC476WYVNDTDJ2XOKDJWITCF3XYDSDT2FLXL"},
			//build.CreditAsset("MXC01","GDXFS34FXPGM3DNQIHQBPFP3DP32XNJX25FJI7OPICA7IQFQGJHFVMJT"),
			//build.CreditAsset("MXC","GD4HUK74YYBT72FDNTIVF3OBHGICU74AWRFUGZOM7CUJV7FLASFYAXXO"),
			build.CreditAmount{Code:"MXC",Issuer:"GD4HUK74YYBT72FDNTIVF3OBHGICU74AWRFUGZOM7CUJV7FLASFYAXXO",Amount:"200"},
			//build.Amount(50),
			//build.NativeAmount{"50"},
		),
	)

	if err != nil {
		panic(err)
	}

	// Sign the transaction to prove you are actually the person sending it.
	txe, err := tx.Sign(distributorPri)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)
	fmt.Println()

	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)


	fmt.Println("Your Account Balance: ")
	gettingAccountDetails.GetBalance(publicKey)
}
