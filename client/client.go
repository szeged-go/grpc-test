package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/szeged-go/grpc-test/registration"
	"google.golang.org/grpc"
)

var address string

func init() {
	flag.StringVar(&address, "address", "", "")
}
func main() {
	conn, err := grpc.Dial("localhost:4543", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := registration.NewRegistrationClient(conn)
	for i := 0; i < 10; i++ {
		resp, err := client.Register(context.Background(), &registration.Request{
			Id:   "12",
			Name: "Emese",
			Age:  27,
		})
		if err != nil {
			log.Print(err)
		}
		log.Print(resp)
	}
}
