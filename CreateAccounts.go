package main

import (
	"fmt"
	"github.com/stellar/go/keypair"
	"io/ioutil"
	"log"
	"net/http"
)

var account1_Pub, account1_Pri, account1 string
var account2_Pub, account2_Pri, account2 string

func main() {
	/**
Get the Seed and Key for account1
 */
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Account1 Public Key: ", pair.Seed())
	log.Println("Account1 Private Key: ", pair.Address())
	log.Println()

	account1_Pub = pair.Seed()
	account1_Pri = pair.Address()

	/**
   Get the Seed and Key for account2
	*/
	pair2, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Account2 Public Key: ", pair2.Seed())
	log.Println("Account2 Private Key: ", pair2.Address())
	log.Println()

	account2_Pub = pair.Seed()
	account2_Pri = pair.Address()

	/*#############################################################################################*/
	/**
	Create account1
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp, err := http.Get("https://friendbot.stellar.org/?addr=" + account1_Pri)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp.Body.Close()
	body, err := ioutil.ReadAll(rresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account1 info: ", string(body))
	account1 = string(body)
	fmt.Println()

	/**
	Create account2
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp2, err := http.Get("https://friendbot.stellar.org/?addr=" + account2_Pri)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp2.Body.Close()
	body2, err := ioutil.ReadAll(rresp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account2 info: ", string(body2))
	account2 = string(body2)
	fmt.Println()
	/*
	##############################################################
	 */
}
