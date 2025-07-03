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

var Cache *redis.Client
var Logger *slog.Logger

func CacheUnaryInterceptor(cache *redis.Client, logger *slog.Logger) grpc.UnaryServerInterceptor {
	Cache = cache
	Logger = logger
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
		if data, err := getFromCache(ctx, key, info.FullMethod); err == nil && data != nil {
			return data, nil
		}
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}

		_ = setCache(ctx, resp, key, info.FullMethod)
		return resp, nil

	}
}

func getFromCache(ctx context.Context, key string, fullMethod string) (data interface{}, err error) {
	cachedStr, er := Cache.Get(key, ctx)
	if er == nil && cachedStr != "" {
		respType := getResponseTypeForMethod(fullMethod)
		if err := json.Unmarshal([]byte(cachedStr), respType); err == nil {
			Logger.Info("cache hit", "method", fullMethod)
			return respType, nil
		}
	}
	return "", er
}

func setCache(ctx context.Context, data interface{}, key string, fullMethod string) error {
	data, err := json.Marshal(data)
	if err == nil {
		err = Cache.Set(key, data, ctx)
		if err != nil {
			Logger.Info("error caching data", "method", fullMethod, "key", key)
		}
		Logger.Info("cache set", "method", fullMethod)
		return nil
	}
	return err
}

// TODO: Add deleting on Update and delete, or add deleting from service itself
//func deleteCache(ctx context.Context, key string) error {
//
//}

// TODO: Make key simple (Id:123, username: 123)
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
	"/users.v3.UserService/UpdateUser": func() interface{} { return &usersv3.UpdateUserResponse{} },
	"/users.v3.UserService/DeleteUser": func() interface{} { return &usersv3.DeleteUserResponse{} },
}

func getResponseTypeForMethod(method string) interface{} {
	if f, ok := responseTypeMap[method]; ok {
		return f()
	}
	return nil
}
