package interceptors

import (
	"context"
	usersv3 "crm/proto/gen/go/users/v3"
	"crm/services/users/pkg/redis"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"strings"
)

func CacheUnaryInterceptor(cache *redis.Client, logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// for GET type methods only
		if !isSafeMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		key := generateCacheKey(info.FullMethod, req)

		if cachedStr, err := cache.Get(key, ctx); err == nil && cachedStr != "" {
			respType := getResponseTypeForMethod(info.FullMethod)
			if err := json.Unmarshal([]byte(cachedStr), respType); err == nil {
				logger.Info("cache hit", "method", info.FullMethod)
				return respType, nil
			}
		}

		resp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}

		if data, err := json.Marshal(resp); err == nil {
			_ = cache.Set(key, data, ctx)
			logger.Info("cache set", "method", info.FullMethod)
		}
		return resp, nil
	}
}

func generateCacheKey(method string, req interface{}) string {
	reqBytes, _ := json.Marshal(req)
	return fmt.Sprintf("grpc_cache:%s:%s", method, string(reqBytes))
}

func isSafeMethod(method string) bool {
	return contains(method, []string{"Get", "By", "Search", "Users"})
}

func contains(str string, substrs []string) bool {
	for _, s := range substrs {
		if strings.Contains(str, s) {
			return true

		}
	}
	return false
}

var responseTypeMap = map[string]func() interface{}{
	"/users.v3.UserService/GetUser": func() interface{} {
		return &usersv3.GetUserResponse{}
	},
	"/users.v3.UserService/GetUsers": func() interface{} {
		return &usersv3.GetUsersResponse{}
	},
	"/users.v3.UserService/Search": func() interface{} {
		return &usersv3.SearchUsersResponse{}
	},
}

func getResponseTypeForMethod(method string) interface{} {
	if f, ok := responseTypeMap[method]; ok {
		return f()
	}
	return nil
}
