package main

import (
	"context"
	"fmt"
	"github.com/CoderShiun/stellar/gettingAccountDetails"
	"github.com/stellar/go/clients/horizon"
)

var Seed1 = "SD64I4OCIUPT6UONUCOT7WVIVOVQNFJI6FLO2NIO5CUDBPXQMUUZSDJA"
var Key1 = "GDZPSKJSFLW4HF77OJFYB3ANCG3CF2PHUZ3YRG5PHPTNF7ZLYGNV7EPV"

func main()  {
	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	err := horizon.DefaultTestNetClient.StreamPayments(ctx, Key1, &cursor, func(payment horizon.Payment) {
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

		gettingAccountDetails.GetBalance(Key1)
	})
	//fmt.Println(esrr)

	if err != nil {
		panic(err)
	}
}
