package main

import (
	"net/http"
	"log"

	_ "github.com/joho/godotenv/autoload"
    common "github.com/hemanth895/commons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/hemanth895/commons/api"

)

var (
	httpAddr = common.EnvString("HTTP_ADDR",":3000")
	orderServiceAddr = "localhost:3000"
)

func main(){

	conn ,err := grpc.Dial(orderServiceAddr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to dial serever : %v",err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at",orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)


	mux := http.NewServeMux()
	handler := NewHandler(c)

	handler.registerRoutes(mux)

		log.Printf("Starting HTTP server at %s",httpAddr)


	if err :=  http.ListenAndServe(httpAddr,mux);err != nil{
		log.Fatal("Failed to start http server");
	}

}