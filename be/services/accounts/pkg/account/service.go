package account

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/lib/tower"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/model"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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
	appTool tower.AppTool
	repo    repository.AccountRepo
}

func NewService(appTool tower.AppTool, repo repository.AccountRepo) *Svc {
	s := Svc{
		appTool,
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

func (s *Svc) GetNameNextNumber(ctx context.Context, name string) (uint32, error) {
	res, err := s.repo.GetNameCounter(ctx, name)

	if err == repository.ErrNameCounterNotFound {
		return 1, nil
	}
	if err != nil {
		return 0, tower.UnhandledError(err)
	}

	return res.Count + 1, nil
}

func (s *Svc) Create(ctx context.Context, authUserId, name string) error {
	nameNumber, err := s.GetNameNextNumber(ctx, name)
	if err != nil {
		return tower.UnhandledError(errors.Wrap(err, "get name number"))
	}

	account := model.Account{
		AccountId: uuid.New(),
		Name: model.AccountName{
			Base:   name,
			Number: nameNumber,
		},
	}

	err = account.Validate()
	if err != nil {
		return tower.ValidationError(err)
	}

	err = s.repo.CreateAccount(ctx, authUserId, account.ToRepoRecord())
	if err == repository.ErrAccountAlreadyCreated {
		return ErrAccountAlreadyCreated
	}
	if err != nil {
		return tower.UnhandledError(errors.Wrap(err, "repo create account"))
	}

	accountClaims := map[string]interface{}{"accountId": account.AccountId.String()}
	err = s.appTool.GetAuth().SetCustomUserClaims(ctx, authUserId, accountClaims)
	if err != nil {
		return tower.UnhandledError(errors.Wrap(err, "update custom claims"))
	}

	return nil
}

func (s *Svc) Delete(ctx context.Context, accountId uuid.UUID, authUserId string) error {
	err := s.repo.DeleteAccountById(ctx, accountId)
	if err != nil {
		return err
	}

	err = s.appTool.GetAuth().DeleteUser(ctx, authUserId)
	if err != nil {
		return err
	}

	return nil
}
