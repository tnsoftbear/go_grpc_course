package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/grpc-course/02-exercise-solution/internal/todo"
	"github.com/cshep4/grpc-course/02-exercise-solution/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
)

const addr = ":50051"

func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()
	todoService := todo.New()
	proto.RegisterTodoServiceServer(grpcServer, todoService)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("cannot listen %s: %w", addr, err)
		}
		slog.Info("Starting gRPC server on address", slog.String("address", lis.Addr().String()))
		err = grpcServer.Serve(lis)
		if err != nil {
			return fmt.Errorf("cannot serve gRPC: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return nil
	})

	return g.Wait()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := run(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("gRPC server run failed", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("closing gRPC server gracefully")
}
