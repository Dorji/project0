package main

import (
	"context"
	"fmt"

	api "atomyzeServer/api/transactions"
)

func main() {
}

func ServerCallerTrue(c api.TransactionsGRPCClient, key []byte) error {
	res, err := c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: key,
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
			return fmt.Errorf("signatureInvalid:%v; %w", res1.SignatureValid, err)
		}
	}
	return nil
}
func ServerCallerFalse(c api.TransactionsGRPCClient, key []byte) error {
	res, err := c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: nil,
	})
	if err == nil {
		return fmt.Errorf("empty Key is work")
	}
	_, err = c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       nil,
		PrivateKey: nil,
	})
	if err == nil {
		return fmt.Errorf("empty Key and Args work")
	}
	_, err = c.CreateTransaction(context.Background(), &api.CreateRequest{
		Args:       nil,
		PrivateKey: key,
	})
	if err == nil {
		return fmt.Errorf("empty Args work")
	}

	req1 := api.VerifyRequest{
		Hash:      res.GetHash(),
		PublicKey: res.GetPublicKey(),
		Signature: res.GetSignature(),
	}

	res1, err := c.VerifyTransaction(context.Background(), &req1)

	if res1.GetSignatureValid() == true {
		if err != nil {
			return fmt.Errorf("signatureInvalid")
		}
	}
	return nil
}
