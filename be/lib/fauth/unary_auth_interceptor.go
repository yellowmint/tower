package fauth

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func UnaryAuthInterceptor(
	authMockEnabled bool,
	claimsService ClaimsService,
	authClient *auth.Client,
	skipAuthMethods []string,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if authMockEnabled {
			testingAuthValue := getHeader(ctx, "xxx-user-id")
			if testingAuthValue == "" {
				return nil, status.Error(codes.Unauthenticated, "unauthenticated")
			}

			ctx = setUserInCtx(ctx, &auth.Token{UID: testingAuthValue}, claimsService)
			return handler(ctx, req)
		}

		ctx, ok := addDataToContext(ctx, claimsService, authClient)
		if !ok && !shouldSkipAuth(info.FullMethod, skipAuthMethods) {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated")
		}

		return handler(ctx, req)
	}
}

func addDataToContext(ctx context.Context, claimsService ClaimsService, authClient *auth.Client) (context.Context, bool) {
	jwt, ok := getJwtFromHeader(ctx)
	if !ok {
		return ctx, false
	}

	token, err := authClient.VerifyIDToken(ctx, jwt)
	if err != nil {
		return ctx, false
	}

	ctx = setUserInCtx(ctx, token, claimsService)

	return ctx, true
}

func getJwtFromHeader(ctx context.Context) (token string, ok bool) {
	authHeaderValue := getHeader(ctx, "authorization")
	if len(authHeaderValue) <= 7 || strings.ToLower(authHeaderValue[0:6]) != "bearer" {
		return "", false
	}

	return authHeaderValue[7:], true
}

func getHeader(ctx context.Context, name string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	headerMetadata, ok := md[name]
	if !ok || len(headerMetadata) == 0 {
		return ""
	}

	return headerMetadata[0]
}

func shouldSkipAuth(method string, exceptions []string) bool {
	for _, exception := range exceptions {
		if method == exception {
			return true
		}
	}

	return false
}
