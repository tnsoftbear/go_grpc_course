package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/cshep4/grpc-course/module5-exercise/internal/token"
	"github.com/cshep4/grpc-course/module5-exercise/proto"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

const bindAddr = ":50051"

// const claimID = "ID"
const ctxKeyClaims = "claims"

var JWTSecret string
var ErrInvalidToken = errors.New("invalid token")

func intrcptr(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("info.FullMethod:", info.FullMethod)
	claims, err := parseToken(ctx)
	if err != nil {
		if errors.Is(err, ErrInvalidToken) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, status.Error(codes.Internal, "error validating token: "+err.Error())
	}

	// add the user ID to the context
	ctx = context.WithValue(ctx, ctxKeyClaims, claims)

	resp, err = handler(ctx, req)
	return resp, err
}

func parseToken(ctx context.Context) (map[string]string, error) {
	authToken, err := getTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(authToken, ".")
	if len(parts) == 3 {
		header, _ := base64.RawURLEncoding.DecodeString(parts[0])
		payload, _ := base64.RawURLEncoding.DecodeString(parts[1])
		fmt.Printf("Header: %s\nPayload: %s\n", header, payload)
	}

	t, err := jwt.Parse(authToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, fmt.Errorf("jwt parse failed: %w: %w", ErrInvalidToken, err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, ErrInvalidToken
	}
	fmt.Println(claims)

	claimMap := make(map[string]string)
	for k, v := range claims {
		s, ok := v.(string)
		if ok {
			claimMap[k] = s
			continue
		}
		ts, ok := v.(float64)
		if ok {
			claimMap[k] = fmt.Sprintf("%.0f", ts)
			continue
		}
		claimMap[k] = fmt.Sprintf("%v", v)
	}
	fmt.Println(claimMap)

	return claimMap, nil
}

func getTokenFromMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(meta["authorization"]) != 1 {
		return "", errors.New("token not found in metadata")
	}

	return meta["authorization"][0], nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	var ok bool
	JWTSecret, ok = os.LookupEnv("JWT_SECRET")
	if !ok {
		panic("JWT_SECRET is required")
	}
	fmt.Println("JWTSecret:", JWTSecret)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(intrcptr))
	tokenService := token.Service{}
	proto.RegisterTokenServiceServer(grpcServer, tokenService)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", bindAddr)
		if err != nil {
			return err
		}
		log.Printf("Starting gRPC server on address: %s", bindAddr)
		err = grpcServer.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return nil
	})

	err := g.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Println("Closing server gracefully")
}
