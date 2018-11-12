package createTestAccount

import (
	"stellarTools/createSeed"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

func GetAccount() string{
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	pair := createSeed.GetSeed()
	address := pair.Address()
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	return string(body)
}