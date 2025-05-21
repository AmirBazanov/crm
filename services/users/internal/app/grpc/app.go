package grpcusers

import (
	users "crm/services/users/internal/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	users.Register(gRPCServer)
	return &App{log, gRPCServer, port}
}
func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}
func (a *App) run() error {
	const op = "grpcusers.Run"
	a.log.Info(op, "Starting gRPC server on port", a.port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.Error(op, err)
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		a.log.Error(op, err)
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil

}

func (a *App) Stop() {
	const op = "grpcusers.Stop"
	a.log.Info(op, "Stopping gRPC server on port", a.port)
	a.gRPCServer.GracefulStop()
}
