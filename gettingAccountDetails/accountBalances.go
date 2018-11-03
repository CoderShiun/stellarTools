package gettingAccountDetails

import (
	"fmt"
	"log"
	"github.com/stellar/go/clients/horizon"
)

func GetBalance(key string) {
	//pair := createSeed.GetSeed()
	account, err := horizon.DefaultTestNetClient.LoadAccount(key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account:", key)

	for _, balance := range account.Balances {
		log.Println(balance)
	}

	//balance := account.Balances

}