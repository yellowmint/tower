package account

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/model"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

var ErrAccountNotFound = tower.Err{
	ErrorValue:     repository.ErrAccountNotFound,
	EndUserMessage: "account not found",
}

var ErrAccountAlreadyCreated = tower.Err{
	ErrorValue:     repository.ErrAccountAlreadyCreated,
	EndUserMessage: "account already created",
}

type Svc struct {
	app  *tower.App
	repo repository.AccountRepo
}

func NewService(app *tower.App, repo repository.AccountRepo) *Svc {
	s := Svc{
		app,
		repo,
	}
	return &s
}

func (s *Svc) Get(ctx context.Context, accountId uuid.UUID) (model.Account, error) {
	res, err := s.repo.GetAccountById(ctx, accountId)

	if err == repository.ErrAccountNotFound {
		return model.Account{}, ErrAccountNotFound
	}
	if err != nil {
		return model.Account{}, tower.UnhandledError(err)
	}

	return model.AccountFromRepo(res), nil
}

func (s *Svc) Create(ctx context.Context, authUserId, name string) error {
	accountId, err := s.repo.CreateAccount(ctx, authUserId, name)
	if err == repository.ErrAccountAlreadyCreated {
		return ErrAccountAlreadyCreated
	}
	if err != nil {
		return tower.UnhandledError(err)
	}

	accountClaims := map[string]interface{}{"accountId": accountId.String()}
	err = s.app.FirebaseClients.Auth.SetCustomUserClaims(ctx, authUserId, accountClaims)
	if err != nil {
		return tower.UnhandledError(err)
	}

	return nil
}

func (s *Svc) Delete(ctx context.Context, accountId uuid.UUID, authUserId string) error {
	err := s.repo.DeleteAccountById(ctx, accountId)
	if err != nil {
		return err
	}

	err = s.app.FirebaseClients.Auth.DeleteUser(ctx, authUserId)
	if err != nil {
		return err
	}

	return nil
}
