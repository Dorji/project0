package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"

	pb "atomyzeServer/api/transactions"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeTransactionsGRPCServer
}

func (s server) CreateTransaction(ctx context.Context, request *pb.CreateRequest) (*pb.CreateReply, error) {
	if request.PrivateKey == nil || request.Args == nil ||
		len(request.PrivateKey) == 0 || len(request.Args) == 0 {
		return nil, fmt.Errorf("transaction server: nil argument")
	}

	args := request.Args

	msg := fmt.Sprintf("%#v", args)

	hash := sha256.Sum256([]byte(msg))
	PrivateKey, _ := x509.ParseECPrivateKey(request.PrivateKey)
	signature, err := ecdsa.SignASN1(rand.Reader, PrivateKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("transaction server: %w", err)
	}

	return &pb.CreateReply{
		Hash:      hash[:],
		PublicKey: generatePublicKeyFromPrivate(*PrivateKey),
		Signature: signature,
	}, nil
}
func generatePublicKeyFromPrivate(privateKey ecdsa.PrivateKey) []byte {
	return elliptic.Marshal(privateKey.Curve, privateKey.X, privateKey.Y)
}

func (s server) VerifyTransaction(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyReplyGet, error) {
	if request.Hash == nil || request.Signature == nil || request.PublicKey == nil ||
		len(request.Hash) == 0 || len(request.Signature) == 0 || len(request.PublicKey) == 0 {
		return nil, fmt.Errorf("transaction server: nil or len=0 argument incoming")
	}
	X, Y := elliptic.Unmarshal(elliptic.P256(), request.GetPublicKey())

	if X == nil {
		return nil, fmt.Errorf("transaction server: bad public key")
	}

	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     X,
		Y:     Y,
	}

	valid := ecdsa.VerifyASN1(&publicKey, request.GetHash(), request.GetSignature())
	return &pb.VerifyReplyGet{SignatureValid: valid}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1337))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionsGRPCServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
