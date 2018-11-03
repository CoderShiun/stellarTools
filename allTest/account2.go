package main

import (
	"github.com/CoderShiun/stellar/buildingTransaction"
	"github.com/CoderShiun/stellar/gettingAccountDetails"
)

var Seed2 = "SCBQ7X2ZYESAYCC2C3V3HMAGLFIW5OGW5MOK4F4N27U7G5HVPJQGZG2P"
var Key2 = "GBU5MUNK3HJGF6FPUCVU6XWZEIW22T6LRAPE23EABUIFOJYDMQQZCKKU"

func main() {
	//Check the Balance for myAccount
	gettingAccountDetails.GetBalance(Key2)

	//	send Transaction
	buildingTransaction.SendTransaction(Seed2, "GB5BEOHMH73KMOA7XAGV4PFJ7INQFMV3PNNXSSM2XYDTWSWYVGXFV6QZ")

}