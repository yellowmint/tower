package repositorymock

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

type RepositoryMock struct {
	data map[string]repository.AccountRecord
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
		data: map[string]repository.AccountRecord{},
	}
}

func (r *RepositoryMock) GetAccountById(_ context.Context, accountId uuid.UUID) (repository.AccountRecord, error) {
	record, ok := r.data[accountId.String()]
	if !ok {
		return repository.AccountRecord{}, repository.ErrAccountNotFound
	}

	return record, nil
}

func (r *RepositoryMock) GetAccountByAuthUserId(_ context.Context, authUserId string) (repository.AccountRecord, error) {
	for _, record := range r.data {
		if record.AuthUserId == authUserId {
			return record, nil
		}
	}

	return repository.AccountRecord{}, repository.ErrAccountNotFound
}

func (r *RepositoryMock) CreateAccount(ctx context.Context, authUserId string, account repository.AccountRecord) error {
	_, err := r.GetAccountByAuthUserId(ctx, authUserId)
	if err == nil {
		return repository.ErrAccountAlreadyCreated
	}

	account.AuthUserId = authUserId
	r.data[account.AccountId] = account

	return nil
}

func (r *RepositoryMock) DeleteAccountById(_ context.Context, accountId uuid.UUID) error {
	delete(r.data, accountId.String())
	return nil
}
