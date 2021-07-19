package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"

	v1 "testRPC/v1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := v1.NewGreeterClient(conn)

	req := new(v1.HelloRequest)
	req.Name = "kratos grpc"
	r, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
