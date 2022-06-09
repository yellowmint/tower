package model

import (
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

type Account struct {
	AccountId uuid.UUID
	Name      string
}

func AccountFromRepo(model repository.AccountRecord) Account {
	return Account{
		AccountId: model.AccountId,
		Name:      model.Name,
	}
}
