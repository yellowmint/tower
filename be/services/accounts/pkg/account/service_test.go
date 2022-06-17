package account_test

import (
	"context"
	"fmt"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower/towermock"
	"git.jetbrains.space/artdecoction/wt/tower/lib/validate"
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

		record, err := repositoryMock.GetAccountByAuthUserId(ctx, authUserId)
		assert.NoError(t, err)
		assert.Equal(t, name, record.Name)

		model, err := svc.Get(ctx, uuid.MustParse(record.AccountId))
		assert.NoError(t, err)
		assert.Equal(t, record.Name, model.Name)

		err = svc.Delete(ctx, model.AccountId, authUserId)
		assert.NoError(t, err)

		_, err = svc.Get(ctx, model.AccountId)
		assert.Equal(t, account.ErrAccountNotFound, err)
	})
}
