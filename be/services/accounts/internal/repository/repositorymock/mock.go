package repositorymock

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/repository"
	"github.com/google/uuid"
)

type RepositoryMock struct {
	accounts     map[string]repository.AccountRecord
	nameCounters map[string]repository.NameCounterRecord
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
		accounts:     map[string]repository.AccountRecord{},
		nameCounters: map[string]repository.NameCounterRecord{},
	}
}

func (r *RepositoryMock) GetAccountById(_ context.Context, accountId uuid.UUID) (repository.AccountRecord, error) {
	record, ok := r.accounts[accountId.String()]
	if !ok {
		return repository.AccountRecord{}, repository.ErrAccountNotFound
	}

	return record, nil
}

func (r *RepositoryMock) GetAccountByAuthUserId(_ context.Context, authUserId string) (repository.AccountRecord, error) {
	for _, record := range r.accounts {
		if record.AuthUserId == authUserId {
			return record, nil
		}
	}

	return repository.AccountRecord{}, repository.ErrAccountNotFound
}

func (r *RepositoryMock) GetNameCounter(_ context.Context, name string) (repository.NameCounterRecord, error) {
	record, ok := r.nameCounters[name]
	if !ok {
		return repository.NameCounterRecord{}, repository.ErrNameCounterNotFound
	}

	return record, nil
}

func (r *RepositoryMock) UpdateNameCounter(nameCounter repository.NameCounterRecord) {
	r.nameCounters[nameCounter.Name] = nameCounter
}

func (r *RepositoryMock) CreateAccount(ctx context.Context, authUserId string, record repository.AccountRecord) error {
	r.UpdateNameCounter(repository.NameCounterRecord{Name: record.Name, Count: record.NameNumber})

	_, err := r.GetAccountByAuthUserId(ctx, authUserId)
	if err == nil {
		return repository.ErrAccountAlreadyCreated
	}

	record.AuthUserId = authUserId
	r.accounts[record.AccountId] = record

	return nil
}

func (r *RepositoryMock) DeleteAccountById(_ context.Context, accountId uuid.UUID) error {
	delete(r.accounts, accountId.String())
	return nil
}
