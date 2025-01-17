package main

import (
  "context"
  "fmt"
)

//price fetcher is an interface that can fetch a price
type PriceFetcher interface {
  FetchPrice(context.Context,string) (float64,error)
}


// priceFetcher implements the pricefetcher interface
type priceFetcher struct {

}


func (s *priceFetcher) FetchPrice (ctx context.Context, ticker string) (float64,error) {
  // this is our business logic 
  // don't use types in business logic 
  // don't use json representations here 
  // business logic needs to be clean and 
  // only handle the business logic 
  return  MockPriceFetcher(ctx,ticker) 

}

var priceMocks = map[string]float64 {
  "BTC":20_000.0,
  "ETH": 200.0,
 "GG": 100_000.0,
}


func MockPriceFetcher(ctx context.Contex,ticker string) (float64,error) {

  price, ok := priceMocks[ticker]

   if !ok {
    return price,fmt.Errorf("The given ticker (%s) is not supported",ticker)
  }

  return price,nil

}

// let's mimic an api call 
