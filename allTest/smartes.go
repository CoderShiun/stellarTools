package main

import (
	"fmt"
	"github.com/CoderShiun/stellar/createSeed"
	"github.com/CoderShiun/stellar/createTestAccount"
	"github.com/CoderShiun/stellar/gettingAccountDetails"
)

func main()  {
	fmt.Print("My seed ist: ", createSeed.GetSeed())
	fmt.Print("My account ist: ", createTestAccount.GetAccount())
	fmt.Print("My account balance ist: ", interface{}(gettingAccountDetails.GetBalance()))
}
