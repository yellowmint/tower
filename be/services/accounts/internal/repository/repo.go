package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountAlreadyCreated = errors.New("account already created")

type AccountRepo interface {
	GetAccountById(ctx context.Context, accountId uuid.UUID) (AccountRecord, error)
	GetAccountByAuthUserId(ctx context.Context, authUserId string) (AccountRecord, error)
	CreateAccount(ctx context.Context, authUserId string, name string) (uuid.UUID, error)
	DeleteAccountById(ctx context.Context, accountId uuid.UUID) error
}

type AccountRecord struct {
	recordId   string
	AccountId  string
	AuthUserId string
	Name       string
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
