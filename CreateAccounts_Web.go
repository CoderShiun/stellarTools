package main

import (
	"fmt"
	"github.com/CoderShiun/stellarTools/changeTrust"
	"github.com/CoderShiun/stellarTools/setOptions"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/stellar/go/keypair"
	"io/ioutil"
	"log"
	"net/http"
)

var account1_PubW, account1_PriW, account1W string
var account2_PubW, account2_PriW, account2W string

func CreateAccount5(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/**
Get the Seed and Key for account1
 */
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w,"Account1 Public Key: ", pair.Address())
	fmt.Fprint(w,"Account1 Private Key: ", pair.Seed())
	log.Println()

	account1_PriW = pair.Seed()
	account1_PubW = pair.Address()

	/**
   Get the Seed and Key for account2
	*/
	pair2, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w,"Account2 Public Key: ", pair2.Address())
	fmt.Fprint(w,"Account2 Private Key: ", pair2.Seed())
	log.Println()

	account2_PriW = pair2.Seed()
	account2_PubW = pair2.Address()

	/*#############################################################################################*/
	/**
	Create account1
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp, err := http.Get("https://friendbot.stellar.org/?addr=" + account1_PubW)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp.Body.Close()
	body, err := ioutil.ReadAll(rresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w,"Account1 info: ", string(body))
	account1W = string(body)
	fmt.Println()

	/**
	Create account2
	 */
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//address := pair.Address()
	rresp2, err := http.Get("https://friendbot.stellar.org/?addr=" + account2_PubW)
	if err != nil {
		log.Fatal(err)
	}

	defer rresp2.Body.Close()
	body2, err := ioutil.ReadAll(rresp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w,"Account2 info: ", string(body2))
	account2W = string(body2)
	fmt.Println()

	/*
	##############################################################
	 */
	// Build the trust line for account1
	fmt.Fprint(w,"Account1: ")
	changeTrust.BuildTrustLine(account1_PubW, account1_PriW)
	fmt.Println()

	// Build the trust line for account2
	fmt.Fprint(w,"Account2: ")
	changeTrust.BuildTrustLine(account2_PubW, account2_PriW)
	fmt.Println()

	/*
	##############################################################
	 */
	// set up the security key for account2
	setOptions.SetSecureMasterKey(account2_PubW, account2_PriW)


}

func main() {
	router := httprouter.New()
	router.POST("/create", CreateAccount5)

	log.Print("Listening...8080");
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
