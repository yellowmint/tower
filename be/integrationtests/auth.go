package integrationtests

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"google.golang.org/grpc/metadata"
)

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

func AuthUser(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "xxx-user-id", id)
}
