package interceptors

import (
	"buf.build/go/protovalidate"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log/slog"
)

func NewValidationInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	op := "NewValidationInterceptor"
	validator, err := protovalidate.New()
	if err != nil {
		panic("failed to create validator: " + err.Error())
	}

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		msg, ok := req.(proto.Message)
		if ok {
			if err := validator.Validate(msg); err != nil {
				logger.Error(op, err.Error())
				return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
			}
		}

		return handler(ctx, req)
	}
}
