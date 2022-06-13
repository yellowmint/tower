package claims

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/fauth"
	"github.com/google/uuid"
)

type BasicClaims struct {
	AuthUserId string
	AccountId  string
}

func (BasicClaims) ClaimsFromToken(token *auth.Token) interface{} {
	accountId, ok := token.Claims["accountId"].(string)
	if !ok {
		accountId = ""
	}

	return BasicClaims{
		AuthUserId: token.UID,
		AccountId:  accountId,
	}
}

func GetClaimsFromCtx(ctx context.Context) (BasicClaims, bool) {
	claims, ok := ctx.Value(fauth.ClaimsContextKey).(BasicClaims)
	if ok {
		return claims, true
	}

	return BasicClaims{}, false
}

func GetAuthUserIdFromCtx(ctx context.Context) string {
	claims, ok := GetClaimsFromCtx(ctx)
	if !ok {
		return ""
	}

	return claims.AuthUserId
}

func GetAccountIdFromCtx(ctx context.Context) uuid.UUID {
	claims, ok := GetClaimsFromCtx(ctx)
	if !ok {
		return uuid.UUID{}
	}

	return uuid.MustParse(claims.AccountId)
}
