package main

import (
	"fmt"
	"github.com/CoderShiun/stellarTools/getFromKyeboard"
	"github.com/CoderShiun/stellarTools/gettingAccountDetails"
)

func main() {
	fmt.Println("Please enter your PublicKey: ")
	publicKey := getFromKyeboard.FromKeyboard()
	gettingAccountDetails.GetBalance(publicKey)
}