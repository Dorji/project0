package main

import (
	"context"
	"fmt"

	api "atomyzeServer/api/transactions"
)

func main() {
}

var ctx = context.Background()

func ServerCallerTrue(c api.TransactionsGRPCClient, key []byte) error {

	res, err := c.CreateTransaction(ctx, &api.CreateRequest{
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
	res1, err := c.VerifyTransaction(ctx, &req1)

	if err != nil || res1.GetSignatureValid() != true {
		return fmt.Errorf("signatureInvalid:%v; %w", res1.SignatureValid, err)
	}

	return nil
}
func ServerCallerFalseCreating(c api.TransactionsGRPCClient, key []byte) error {
	_, err := c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: nil,
	})
	if err == nil {
		return fmt.Errorf("empty Key is work")
	}
	_, err = c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       nil,
		PrivateKey: nil,
	})
	if err == nil {

		return fmt.Errorf("empty Key and Args work")
	}
	_, err = c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       nil,
		PrivateKey: key,
	})
	if err == nil {
		return fmt.Errorf("empty Args work")
	}
	_, err = c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       [][]byte{},
		PrivateKey: key,
	})
	if err == nil {
		return fmt.Errorf("zero Args work")
	}
	_, err = c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: []byte{},
	})
	if err == nil {
		return fmt.Errorf("zero Key work")
	}
	return nil
}

func ServerCallerFalseVerification(c api.TransactionsGRPCClient, key []byte) error {

	res, err := c.CreateTransaction(ctx, &api.CreateRequest{
		Args:       [][]byte{[]byte("abcd"), []byte("bcde"), []byte("1234"), []byte("4321")},
		PrivateKey: key,
	})
	if err != nil {
		return err
	}

	tests := map[string]struct {
		data     api.VerifyRequest
		response interface{}
	}{
		"nil in verification1": {
			api.VerifyRequest{
				Hash:      nil,
				PublicKey: res.GetPublicKey(),
				Signature: res.GetSignature(),
			},
			nil,
		},
		"nil in verification2": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: nil,
				Signature: res.GetSignature(),
			},
			nil,
		},
		"nil in verification3": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: res.GetPublicKey(),
				Signature: nil,
			},
			nil,
		},
		"len=0 in verification1": {
			api.VerifyRequest{
				Hash:      []byte{},
				PublicKey: res.GetPublicKey(),
				Signature: res.GetSignature(),
			},
			nil,
		},
		"len=0 in verification2": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: []byte{},
				Signature: res.GetSignature(),
			},
			nil,
		},
		"len=0 in verification3": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: res.GetPublicKey(),
				Signature: []byte{},
			},
			nil,
		},

		"bad public key in verification": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: []byte("bad public key"),
				Signature: res.GetSignature(),
			},
			false,
		},
	}
	for name, v := range tests {

		res1, error := c.VerifyTransaction(ctx, &v.data)
		if res1.GetSignatureValid() == true {
			return fmt.Errorf("%v is work!? Response: %v Error:%w", name, res1.SignatureValid, error)
		}
	}

	tests1 := map[string]struct {
		data     api.VerifyRequest
		response interface{}
	}{
		"bad hash in verification": {
			api.VerifyRequest{
				Hash:      []byte("bad hash"),
				PublicKey: res.GetPublicKey(),
				Signature: res.GetSignature(),
			},
			false,
		},
		"bad signature in verification": {
			api.VerifyRequest{
				Hash:      res.GetHash(),
				PublicKey: res.GetPublicKey(),
				Signature: []byte("bad signature"),
			},
			false,
		},
	}
	for name1, v1 := range tests1 {

		res1, error1 := c.VerifyTransaction(ctx, &v1.data)
		if res1.GetSignatureValid() == true {
			return fmt.Errorf("%v is work!? Response!!!: %v Error:%w", name1, res1.GetSignatureValid(), error1)
		}
		if error1 != nil {
			return fmt.Errorf("%v is work!? Response: %v Error!!!:%w", name1, res1.GetSignatureValid(), error1)
		}
	}

	return nil
}
