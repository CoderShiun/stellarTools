package setOptions

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func SetSecureMasterKey(pubKey, priKey string) {
	tx, err := build.Transaction(
		build.SourceAccount{pubKey},
		build.AutoSequence{horizon.DefaultTestNetClient},
		//build.Sequence{1},
		build.TestNetwork,
		build.SetOptions(
			//build.InflationDest(),
			//build.SetAuthRequired(),
			//build.SetAuthRevocable(),
			//build.SetAuthImmutable(),
			//build.ClearAuthRequired(),
			//build.ClearAuthRevocable(),
			//build.ClearAuthImmutable(),
			build.MasterWeight(1),
			build.SetThresholds(0, 0, 2),
			//build.HomeDomain(),
			build.AddSigner("GCMXW6TWZIT4GHMEN2YE33MJT4OF3R4PKCVR77Z7MHDD4C5SWUDJCA2L", 2),
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(priKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	fmt.Println()

	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Master security key successful: ")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}
