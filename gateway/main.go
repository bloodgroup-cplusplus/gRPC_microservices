package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"github.com/bloodgroup-cplusplus/gRPC_microservices/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR",":3000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Http server running at %s",httpAddr)
	if err:= http.ListenAndServe(httpAddr,mux); err !=nil {
		log.Fatal("Failed to start http server")
	}

}