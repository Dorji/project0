package main

import (
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
	c := api.NewTransactionsGRPCClient(conn)
	t.Run("test right grpc calls", func(t *testing.T) {
		assert.NoError(t, ServerCallerTrue(c))
	})
}
