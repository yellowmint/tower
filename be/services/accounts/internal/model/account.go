package model

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/validate"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

type Account struct {
	AccountId uuid.UUID `validate:"required"`
	Name      string    `validate:"required,gte=6,lte=16"`
}

func AccountFromRepo(model repository.AccountRecord) Account {
	return Account{
		AccountId: uuid.MustParse(model.AccountId),
		Name:      model.Name,
	}
}

func (a *Account) ToRepoRecord() repository.AccountRecord {
	return repository.AccountRecord{
		AccountId: a.AccountId.String(),
		Name:      a.Name,
	}
}

func (a *Account) Validate() error {
	return validate.Validate(a)
}
