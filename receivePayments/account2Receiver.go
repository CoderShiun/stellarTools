package main

import (
	"context"
	"fmt"
	"github.com/CoderShiun/stellarTools/gettingAccountDetails"
	"github.com/stellar/go/clients/horizon"
)

var Seed2 = "SCBQ7X2ZYESAYCC2C3V3HMAGLFIW5OGW5MOK4F4N27U7G5HVPJQGZG2P"
var Key2 = "GBU5MUNK3HJGF6FPUCVU6XWZEIW22T6LRAPE23EABUIFOJYDMQQZCKKU"

func main()  {
	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	err := horizon.DefaultTestNetClient.StreamPayments(ctx, Key2, &cursor, func(payment horizon.Payment) {
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

		gettingAccountDetails.GetBalance(Key2)
	})
	//fmt.Println(esrr)

	if err != nil {
		panic(err)
	}
}
