package model

import (
	"fmt"
	"git.jetbrains.space/artdecoction/wt/tower/lib/validate"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

type Account struct {
	AccountId uuid.UUID `validate:"required"`
	Name      AccountName
}

type AccountName struct {
	Base   string `validate:"required,gte=6,lte=16"`
	Number uint32 `validate:"required,gte=1"`
}

func AccountFromRepo(model repository.AccountRecord) Account {
	return Account{
		AccountId: uuid.MustParse(model.AccountId),
		Name: AccountName{
			Base:   model.Name,
			Number: model.NameNumber,
		},
	}
}

func (a *Account) ToRepoRecord() repository.AccountRecord {
	return repository.AccountRecord{
		AccountId:  a.AccountId.String(),
		Name:       a.Name.Base,
		NameNumber: a.Name.Number,
	}
}

func (a *Account) Validate() error {
	return validate.Validate(a)
}

func (n *AccountName) GetFullName() string {
	return fmt.Sprintf("%s/%d", n.Base, n.Number)
}
