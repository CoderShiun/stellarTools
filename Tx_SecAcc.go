package main

import (
	"bytes"
	"fmt"
	"stellarTools/buildingTransaction"
	"stellarTools/getFromKyeboard"
	"stellarTools/gettingAccountDetails"
	"io/ioutil"
	"net/http"
)

func CheckLimit() (limit string){
	url := "http://127.0.0.1:8080/create"
	var jsonStr = []byte(`{"title":"0 or 1"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}

func main() {
	fmt.Println("Please enter your Account: ")
	publicKey := getFromKyeboard.FromKeyboard()

	fmt.Println("Please enter your PrivateKey: ")
	privateKey := getFromKyeboard.FromKeyboard()

	fmt.Println("How much MXC you would like to send? ")
	amount := getFromKyeboard.FromKeyboard()

	fmt.Println("Please enter the account that you want to send: ")
	destinationAcc := getFromKyeboard.FromKeyboard()

	limit := CheckLimit()
	if limit == "0" {
		fmt.Println("Sorry, you are unable to trans more MXC today!")
		gettingAccountDetails.GetBalance(destinationAcc)

		return
	}else {
		buildingTransaction.SendMXC(publicKey, privateKey, amount, destinationAcc)

		gettingAccountDetails.GetBalance(destinationAcc)
	}
}
