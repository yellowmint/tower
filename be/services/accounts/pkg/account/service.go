package account

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/model"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

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
		return model.Account{}, tower.Err{
			ErrorValue:     err,
			EndUserMessage: "account not found for given accountId",
		}
	}
	if err != nil {
		return model.Account{}, tower.Err{
			ErrorValue:     err,
			EndUserMessage: "unhandled error",
		}
	}

	return model.AccountFromRepo(res), nil
}

func (s *Svc) GetByUserId(ctx context.Context, userId uuid.UUID) (model.Account, error) {
	res, err := s.repo.GetAccountByAuthUserId(ctx, userId)
	if err == repository.ErrAccountNotFound {
		return model.Account{}, tower.Err{
			ErrorValue:     err,
			EndUserMessage: "account not found for given userId",
		}
	}
	if err != nil {
		return model.Account{}, tower.Err{
			ErrorValue:     err,
			EndUserMessage: "unhandled error",
		}
	}

	return model.AccountFromRepo(res), nil
}

func (s *Svc) Create(ctx context.Context, userId uuid.UUID, name string) error {
	err := s.repo.CreateAccount(ctx, userId, name)
	if err == repository.ErrAccountAlreadyCreated {
		return tower.Err{
			ErrorValue:     err,
			EndUserMessage: "account already exist for given userId",
		}
	}
	if err != nil {
		return tower.Err{
			ErrorValue:     err,
			EndUserMessage: "unhandled error",
		}
	}

	return nil
}

func (s *Svc) DeleteByUserId(ctx context.Context, userId uuid.UUID) error {
	return s.repo.DeleteAccountByAuthUserId(ctx, userId)
}
