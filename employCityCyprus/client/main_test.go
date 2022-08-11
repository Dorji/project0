package main

import (
	"log"
	"testing"

	api "employCity/api/grpc_memcached"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = []byte{}

func Test_ServerCallerTrue(t *testing.T) {

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewMyGRPCClient(conn)
	t.Run("test right grpc calls", func(t *testing.T) {
		assert.NoError(t, ServerCallerSetGetDelete(c))
	})

	t.Run("test wrong grpc calls", func(t *testing.T) {
		assert.NoError(t, ServerWrongCaller(c))
	})
}
