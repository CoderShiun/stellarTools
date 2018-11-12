package main

import (
	"fmt"
	"stellarTools/getFromKyeboard"
	"stellarTools/gettingAccountDetails"
)

func main() {
	fmt.Println("Please enter your PublicKey: ")
	publicKey := getFromKyeboard.FromKeyboard()
	gettingAccountDetails.GetBalance(publicKey)
}