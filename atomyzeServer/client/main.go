package main

import (
	"context"
	"fmt"

	api "atomyzeServer/api/transactions"
)

func main() {
}

func ServerCallerTrue(c api.TransactionsGRPCClient) error {

	res, err := c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: []byte("MY_PRIVATE_KEY"),
	})
	if err != nil {
		return err
	}
	req1 := api.VerifyRequest{
		Hash:      res.GetHash(),
		PublicKey: res.GetPublicKey(),
		Signature: res.GetSignature(),
	}

	res1, err := c.VerifyTransaction(context.Background(), &req1)

	if res1.GetSignatureValid() != true {
		if err != nil {
			return fmt.Errorf("signatureInvalid")
		}
	}
	return nil
}
