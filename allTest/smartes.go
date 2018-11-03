package main

import (
	"fmt"
	"github.com/CoderShiun/stellarTools/createSeed"
	"github.com/CoderShiun/stellarTools/createTestAccount"
)

func main()  {
	fmt.Print("My seed ist: ", createSeed.GetSeed())
	fmt.Print("My account ist: ", createTestAccount.GetAccount())
	//fmt.Print("My account balance ist: ", interface{}(gettingAccountDetails.GetBalance()))
}
