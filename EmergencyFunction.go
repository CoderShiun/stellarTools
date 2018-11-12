package main

import (
	"fmt"
	"stellarTools/buildingTransaction"
	"stellarTools/getFromKyeboard"
	"stellarTools/gettingAccountDetails"
)

func main() {
	fmt.Println("Enter the Public Key: ")
	account2_publicKey := getFromKyeboard.FromKeyboard()

	balance := gettingAccountDetails.GetBalance(account2_publicKey)
	fmt.Println()
	fmt.Println()
	//fmt.Println(balance[0].Balance)
	//fmt.Println(len(balance[0].Balance))
	//fmt.Println(balance[0].Balance[0:3])

	// you have to turn it back the to string, otherwise is unicode
	len := len(balance[0].Balance)
	var amount string
	for i := 0; i< len; i++ {
		if string(balance[0].Balance[i]) == "." {
			break
		}
		//fmt.Println(balance[0].Balance[i])
		amount += string(balance[0].Balance[i])
	}

	//fmt.Println("Enter the amount: ")
	//amount := getFromKyeboard.FromKeyboard()

	buildingTransaction.SendMXC(account2_publicKey, "SCYWWO3FSPBHAYB5PTO35X2SU65MOLVTDOGO4MZAFQ2UQ2Y76PWV6S34", amount, "GDNDAZ4Q56GWNM3U3QQZZQXZCC5LPCB4EO2DTB6QLLC2ADVHM4TEKRJ2")
	//
	gettingAccountDetails.GetBalance(account2_publicKey)
}
