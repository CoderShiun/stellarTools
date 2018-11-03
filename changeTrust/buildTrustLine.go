package main

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)


func main() {


	tx, err := build.Transaction(
		build.SourceAccount{"GCM3B4YCWQJFNJWGWNYECV6IU4F7WFHH6CGUCTLBAW4XXIKRORP53HXI"},
		build.AutoSequence{horizon.DefaultTestNetClient},
		//build.Sequence{1},
		build.TestNetwork,
		build.Trust("MXC", "GD4HUK74YYBT72FDNTIVF3OBHGICU74AWRFUGZOM7CUJV7FLASFYAXXO"),
		//build.ChangeTrust(),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign("SBOT4Q7JJ33FKFLJMBWGLSCXE3FY6J2AEBJ7P2MFPEBVKJB2LRGHFY3P")
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

	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}
