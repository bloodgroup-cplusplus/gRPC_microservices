package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bloodgroup-cplusplus/gRPC_microservices/common"
	pb "github.com/bloodgroup-cplusplus/gRPC_microservices/common/api"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR",":3000")
	orderServiceAddr = "localhost:3000"
)

func main() {
	conn,err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err !=nil {
		log.Fatalf("Failed to dial server: %v",err)
	}
	defer conn.Close()
	log.Println("dialing orders service at",orderServiceAddr)
	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Http server running at %s",httpAddr)
	if err:= http.ListenAndServe(httpAddr,mux); err !=nil {
		log.Fatal("Failed to start http server")
	}

}