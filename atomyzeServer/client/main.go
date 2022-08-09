package main

import (
	"context"
	"log"

	api "atomyzeServer/api/transactions"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ServerCaller()
}
func ServerCaller() {
	result := make([][]byte, 0)
	conn, err := grpc.Dial(":1337", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewTransactionsGRPCClient(conn)
	res, err := c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       [][]byte{},
		PrivateKey: nil,
	})
	result = append(result, res.GetSignature())
	result = append(result, res.GetPublicKey())
	result = append(result, res.GetHash())

	res1, err := c.VerifyTransaction(context.Background(), &api.VerifyRequest{
		Hash:      res.GetSignature(),
		PublicKey: res.GetPublicKey(),
		Signature: res.GetHash(),
	})

	if res1.GetSignatureValid() != true {

	}
}
