package gettingAccountDetails

import (
	"fmt"
	"github.com/stellar/go/clients/horizon"
	horizon2 "github.com/stellar/go/protocols/horizon"
	"log"
)

func GetBalance(key string) (amount []horizon2.Balance){
	//pair := createSeed.GetSeed()
	account, err := horizon.DefaultTestNetClient.LoadAccount(key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account:", key)


	for _, balance := range account.Balances {
		log.Println(balance)
	}

	return account.Balances

	//balance := account.Balances

}