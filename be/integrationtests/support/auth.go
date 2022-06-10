package support

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func (s *Support) AuthorizeInContext(ctx context.Context, authUserId string, accountId uuid.UUID) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "xxx-auth-user-id", authUserId)
	ctx = metadata.AppendToOutgoingContext(ctx, "xxx-account-id", accountId.String())

	return ctx
}

const AccountEmail = "bubble@decoct.dev"

func GetUser(app tower.App) *auth.UserRecord {
	ctx := context.Background()

	user, err := app.FirebaseClients.Auth.GetUserByEmail(ctx, AccountEmail)
	if err != nil {
		panic(err)
	}

	return user
}

func CreateTestUser(ctx context.Context, app tower.App) *auth.UserRecord {
	params := (&auth.UserToCreate{}).Email(AccountEmail)

	user, err := app.FirebaseClients.Auth.CreateUser(ctx, params)
	if err != nil {
		panic(err)
	}

	return user
}
