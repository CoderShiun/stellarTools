package main

import (
	"github.com/CoderShiun/stellar/buildingTransaction"
	"github.com/CoderShiun/stellar/gettingAccountDetails"
)

var Seed1 = "SD64I4OCIUPT6UONUCOT7WVIVOVQNFJI6FLO2NIO5CUDBPXQMUUZSDJA"
var Key1 = "GDZPSKJSFLW4HF77OJFYB3ANCG3CF2PHUZ3YRG5PHPTNF7ZLYGNV7EPV"

func main() {
	//Check the Balance for myAccount
	gettingAccountDetails.GetBalance(Key1)

	//	send Transaction
	buildingTransaction.SendTransaction(Seed1, "GBU5MUNK3HJGF6FPUCVU6XWZEIW22T6LRAPE23EABUIFOJYDMQQZCKKU")

}



