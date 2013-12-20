package main

//https://openexchangerates.org/api/latest.json?app_id=38f89b9bd9db4540bb12de3794f11977

import (
	"fmt"
	//"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Currency struct {
	rate   float64
	symbol string
}

func main() {

	getCurrentExchangeRates()
	os.Exit(0)

	//arg1 should contain the amount of currency you want to convert
	arg1 := os.Args[1]
	amount, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		panic(err)
	}

	//arg2 should contain the curreny code GBP, USD etc
	arg2 := os.Args[2]
	fromCurrency := strings.Trim(arg2, " ")

	if fromCurrency == "" {
		panic("No curreny code has been set, please provide GPB, USD etc as the 2nd argument")
	}

	//arg3 should contain the curreny code GBP, USD etc
	arg3 := os.Args[3]
	toCurrency := strings.Trim(arg3, " ")

	if toCurrency == "" {
		panic("No curreny code has been set, please provide GPB, USD etc as the 3nd argument")
	}

	total, symbol := calculateAmount(amount, fromCurrency, toCurrency)

	fmt.Printf("%s%6.2f\n", symbol, total)
}

func calculateAmount(amount float64, from string, to string) (float64, string) {
	currencies := make(map[string]Currency)
	currencies["GBP"] = Currency{rate: 1.0, symbol: "£"}
	currencies["USD"] = Currency{rate: 1.63, symbol: "$"}
	currencies["EUR"] = Currency{rate: 1.2, symbol: "E"}
	currencies["YEN"] = Currency{rate: 170.75, symbol: "Y"}

	baseCurrency := "GBP"

	total := 0.0

	if from == baseCurrency {
		total = amount * currencies[to].rate
	} else {
		total = amount / currencies[from].rate * currencies[to].rate
	}

	return total, currencies[to].symbol
}

func getCurrentExchangeRates() {
	data, err := ioutil.ReadFile("config/currencies.json")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(data)
}
