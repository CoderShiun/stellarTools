package main

import (
	"fmt"
	"github.com/CoderShiun/stellarTools/buildingTransaction"
	"github.com/CoderShiun/stellarTools/getFromKyeboard"
	"github.com/CoderShiun/stellarTools/gettingAccountDetails"
)

func main() {
	fmt.Println("Please enter your Account: ")
	publicKey := getFromKyeboard.FromKeyboard()

	fmt.Println("Please enter your PrivateKey: ")
	privateKey := getFromKyeboard.FromKeyboard()

	fmt.Println("How much MXC you would like to send? ")
	amount := getFromKyeboard.FromKeyboard()

	fmt.Println("Please enter your second account or the account you want to send: ")
	destinationAcc := getFromKyeboard.FromKeyboard()

	buildingTransaction.SendMXC(publicKey, privateKey, amount, destinationAcc)

	gettingAccountDetails.GetBalance(destinationAcc)
}