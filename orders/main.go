package main

import "context"

func main() {
	// we will initiate everytighin

	store := NewStore()
	svc := NewService(store)
	svc.CreateOrder(context.Background())

}