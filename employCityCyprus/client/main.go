package main

import (
	"context"
	"fmt"

	api "employCity/api/grpc_memcached"
)

func main() {

}

var ctx = context.Background()

func ServerCallerSetGetDelete(c api.MyGRPCClient) error {

	res, err := c.Set(ctx, &api.SetRequest{
		Id:   "1212",
		Body: "BodyBodyBodyBodyBody",
	})
	if err != nil || res.GetStatus() != "created" {
		return fmt.Errorf("call grpcStorage set err:%w", err)
	}
	req1 := api.GetRequest{
		Id: "1212",
	}
	res1, err := c.Get(ctx, &req1)
	if err != nil || res1.GetBody() != "BodyBodyBodyBodyBody" {
		return fmt.Errorf("call grpcStorage get 1 err:%w", err)
	}

	req2 := api.DeleteRequest{
		Id: "1212",
	}

	res2, err := c.Delete(ctx, &req2)
	if err != nil || res2.GetStatus() != "deleted" {
		return fmt.Errorf("call grpcStorage delete err: %w", err)
	}

	res3, err := c.Get(ctx, &req1)
	if err == nil || res3.GetBody() != "" {
		return fmt.Errorf("call grpcStorage get 2 err:%w", err)
	}

	return nil
}
func ServerWrongCaller(c api.MyGRPCClient) error {

	_, err := c.Set(ctx, &api.SetRequest{
		Id:   "111",
		Body: "BodyBodyBodyBodyBody",
	})
	if err != nil {
		return fmt.Errorf("call grpcStorage set err:%w", err)
	}
	req1 := api.GetRequest{
		Id: "2222",
	}
	res1, err := c.Get(ctx, &req1)
	if err == nil || res1.GetBody() == "BodyBodyBodyBodyBody" {
		return fmt.Errorf("call grpcStorage get 1 err:%w", err)
	}

	req2 := api.DeleteRequest{
		Id: "323232",
	}

	res2, err := c.Delete(ctx, &req2)
	if err == nil || res2.GetStatus() == "deleted" {
		return fmt.Errorf("call grpcStorage delete err: %w", err)
	}
	req3 := api.GetRequest{
		Id: "111",
	}
	res3, err := c.Get(ctx, &req3)
	if err != nil || res3.GetBody() != "BodyBodyBodyBodyBody" {
		return fmt.Errorf("call grpcStorage get 2 err:%w", err)
	}

	return nil
}
