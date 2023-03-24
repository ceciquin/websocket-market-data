package server

import (
	"math/rand"
)

type Symbol int

const (
	AAPL Symbol = iota
	MELI
	BAC
)

type MarketData struct {
	Symbol          string
	Selling_price   float64
	Buying_price    float64
	Purchase_amount float64
	Selling_amount  float64
}

// function to deliver random market data
func GenerateMarketData() MarketData {

	selling_price := rand.Float64() * 300
	buying_price := rand.Float64() * 300
	randPurchaseAmount := make(chan float64)
	defer close(randPurchaseAmount)
	randSellingAmount := make(chan float64)
	defer close(randSellingAmount)

	go evaluatePurchaseAmount(randPurchaseAmount, selling_price, buying_price)

	go evaluateSellingAmount(randSellingAmount, selling_price, buying_price)

	return MarketData{
		Symbol:          "AAPL",
		Selling_price:   selling_price,
		Buying_price:    buying_price,
		Purchase_amount: <-randPurchaseAmount,
		Selling_amount:  <-randSellingAmount}

}

// function to calculate purchase amount
func evaluatePurchaseAmount(ch chan<- float64, selling_price float64, buying_price float64) {

	purchase_amount := rand.Float64() * 40
	pb_ratio := selling_price / (selling_price * purchase_amount)

	if pb_ratio < 1 {
		purchase_amount = 6
	} else {
		purchase_amount = 15
	}

	ch <- purchase_amount

}

// function to calculate selling amount
func evaluateSellingAmount(ch chan<- float64, selling_price float64, buying_price float64) {

	selling_amount := rand.Float64() * 30
	pb_ratio := selling_price / (selling_price * selling_amount)

	if pb_ratio < 1 {
		selling_amount = rand.Float64() * 7
	} else {
		selling_amount = rand.Float64() * 30
	}

	ch <- selling_amount
}
