package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"log"
	"testing"

	api "atomyzeServer/api/transactions"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = []byte{}

func Test_ServerCallerTrue(t *testing.T) {

	conn, err := grpc.Dial(":1337", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	privateKeyOriginal, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKeyBytes, _ := x509.MarshalECPrivateKey(privateKeyOriginal)
	c := api.NewTransactionsGRPCClient(conn)
	t.Run("test right grpc calls", func(t *testing.T) {
		assert.NoError(t, ServerCallerTrue(c, privateKeyBytes))
	})
	t.Run("test wrong grpc calls", func(t *testing.T) {
		assert.NoError(t, ServerCallerFalse(c, privateKeyBytes))
	})
}
