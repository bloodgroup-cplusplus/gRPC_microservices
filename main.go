package main 
import (
  "context"
  "fmt"
  "log"
)


func main () {

  var svc PriceFetcher
  svc = NewLoggingService(&priceFetcher())

  price,err=svc.FetchPrice(context.Background(),"ETH")
  if err !=nil {
    log.Fatal(err)
  }

  fmt.Println(price)

  

}
