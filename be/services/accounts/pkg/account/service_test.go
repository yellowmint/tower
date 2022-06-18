package account_test

import (
	"context"
	"fmt"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower/towermock"
	"git.jetbrains.space/artdecoction/wt/tower/lib/validate"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository/repositorymock"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/pkg/account"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountService(t *testing.T) {
	t.Parallel()

	validate.Init()

	appToolMock := towermock.NewTowerMockApp()
	repositoryMock := repositorymock.NewRepositoryMock()
	svc := account.NewService(appToolMock, repositoryMock)

	t.Run("Get no found", func(t *testing.T) {
		_, err := svc.Get(context.Background(), uuid.New())
		assert.Equal(t, account.ErrAccountNotFound, err)
	})

	t.Run("Create, Get and Delete", func(t *testing.T) {
		ctx := context.Background()
		authUserId := fmt.Sprintf("%d", time.Now().UTC().UnixNano())
		name := "tomek123"

		err := svc.Create(ctx, authUserId, name)
		assert.NoError(t, err)

		accountId := getAccountId(t, ctx, repositoryMock, authUserId)

		model, err := svc.Get(ctx, accountId)
		assert.NoError(t, err)
		assert.Equal(t, name, model.Name.Base)

		err = svc.Delete(ctx, model.AccountId, authUserId)
		assert.NoError(t, err)

		_, err = svc.Get(ctx, model.AccountId)
		assert.Equal(t, account.ErrAccountNotFound, err)
	})

	t.Run("Create assigns correct name numbers", func(t *testing.T) {
		ctx := context.Background()
		authUser1Id := fmt.Sprintf("%d", time.Now().UTC().UnixNano())
		authUser2Id := fmt.Sprintf("%d", time.Now().UTC().UnixNano())
		randStr := fmt.Sprintf("%d", time.Now().UTC().UnixNano())
		name := "sam" + randStr[6:19]

		err := svc.Create(ctx, authUser1Id, name)
		assert.NoError(t, err)

		account1Id := getAccountId(t, ctx, repositoryMock, authUser1Id)

		model1, err := svc.Get(ctx, account1Id)
		assert.NoError(t, err)
		assert.Equal(t, name, model1.Name.Base)
		assert.Equal(t, uint32(1), model1.Name.Number)

		err = svc.Create(ctx, authUser2Id, name)
		assert.NoError(t, err)

		account2Id := getAccountId(t, ctx, repositoryMock, authUser2Id)

		model2, err := svc.Get(ctx, account2Id)
		assert.NoError(t, err)
		assert.Equal(t, name, model2.Name.Base)
		assert.Equal(t, uint32(2), model2.Name.Number)
	})
}

func getAccountId(t *testing.T, ctx context.Context, repositoryMock repository.AccountRepo, authUserId string) uuid.UUID {
	record, err := repositoryMock.GetAccountByAuthUserId(ctx, authUserId)
	assert.NoError(t, err)

	return uuid.MustParse(record.AccountId)
}
