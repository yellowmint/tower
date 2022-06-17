package support

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	rpcpublicv1 "git.jetbrains.space/artdecoction/wt/tower/contracts/accounts/rpcpublic/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"time"
)

type Authorization struct {
	AuthUserId string
	AccountId  uuid.UUID
}

func (s *Support) NewFakeAuthorization() Authorization {
	return Authorization{
		AuthUserId: s.NewFakeAuthUserId(),
		AccountId:  uuid.New(),
	}
}

func (s *Support) NewAuthorization() Authorization {
	authUser := s.CreateTestAuthUser(context.Background())

	authorization := Authorization{
		AuthUserId: authUser.UID,
		AccountId:  uuid.UUID{},
	}

	ctx := s.AuthorizeInContext(context.Background(), authorization)

	request := &rpcpublicv1.CreateMyAccountRequest{
		Name: "bubble" + authUser.Email[16:26], // bubble.1655486512806060832@decoct.dev -> bubble2806060832
	}

	cc := s.NewGrpcClientConn("accounts")
	client := rpcpublicv1.NewAccountsServiceClient(cc)

	_, err := client.CreateMyAccount(ctx, request)
	if err != nil {
		panic(err)
	}

	authorization.AccountId = s.GetAccountIdByAuthUserId(context.Background(), authUser.UID)

	return authorization
}

func (s *Support) NewFakeAuthUserId() string {
	return fmt.Sprintf("%d", time.Now().UTC().Unix())
}

func (s *Support) AuthorizeInContext(ctx context.Context, auth Authorization) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "xxx-auth-user-id", auth.AuthUserId)
	ctx = metadata.AppendToOutgoingContext(ctx, "xxx-account-id", auth.AccountId.String())

	return ctx
}

func (s *Support) CreateTestAuthUser(ctx context.Context) *auth.UserRecord {
	email := fmt.Sprintf("bubble.%d@decoct.dev", time.Now().UTC().UnixNano())

	userParams := &auth.UserToCreate{}
	userParams.Email(email)

	user, err := s.FirebaseClients.Auth.CreateUser(ctx, userParams)
	if err != nil {
		panic(err)
	}

	return user
}

func (s *Support) GetAccountIdByAuthUserId(ctx context.Context, authUserId string) uuid.UUID {
	docs, err := s.FirebaseClients.Firestore.Collection("accounts:accounts").
		Where("AuthUserId", "==", authUserId).
		Where("DeletedAt", "==", nil).
		Limit(1).
		Documents(ctx).
		GetAll()

	if err != nil {
		panic(err)
	}
	if len(docs) == 0 {
		panic("account not found by authUserId")
	}

	var record struct {
		AccountId string
	}

	err = docs[0].DataTo(&record)
	if err != nil {
		panic(err)
	}

	return uuid.MustParse(record.AccountId)
}
