package claims

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fauth"
)

type BasicClaims struct {
	UserId      string
	DisplayName string
}

func (BasicClaims) ClaimsFromToken(token *auth.Token) interface{} {
	name, ok := token.Claims["name"].(string)
	if !ok {
		name = ""
	}

	return BasicClaims{
		UserId:      token.UID,
		DisplayName: name,
	}
}

func GetClaimsFromCtx(ctx context.Context) (BasicClaims, bool) {
	claims, ok := ctx.Value(fauth.ClaimsContextKey).(BasicClaims)
	if ok {
		return claims, true
	}

	return BasicClaims{}, false
}

func GetUserIdFromCtx(ctx context.Context) string {
	claims, ok := GetClaimsFromCtx(ctx)
	if !ok {
		return ""
	}

	return claims.UserId
}
