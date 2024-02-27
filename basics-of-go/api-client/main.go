package main

import (
	"fmt"
	"sync"

	"dantejs.com/go/api-client/api"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is USD %v\n", rate.Currency, rate.Price)
	}
}
