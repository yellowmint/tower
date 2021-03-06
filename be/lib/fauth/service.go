package fauth

import (
	"context"
	"firebase.google.com/go/v4/auth"
)

type ClaimsService interface {
	ClaimsFromToken(token *auth.Token) interface{}
}

const ClaimsContextKey string = "auth-claims"

func setClaimsInCtx(ctx context.Context, token *auth.Token, claimsService ClaimsService) context.Context {
	claims := claimsService.ClaimsFromToken(token)
	return context.WithValue(ctx, ClaimsContextKey, claims)
}
