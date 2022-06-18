package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountAlreadyCreated = errors.New("account already created")

var ErrNameCounterNotFound = errors.New("name counter not found")

type AccountRepo interface {
	GetAccountById(ctx context.Context, accountId uuid.UUID) (AccountRecord, error)
	GetAccountByAuthUserId(ctx context.Context, authUserId string) (AccountRecord, error)
	GetNameCounter(ctx context.Context, name string) (res NameCounterRecord, err error)
	CreateAccount(ctx context.Context, authUserId string, account AccountRecord) error
	DeleteAccountById(ctx context.Context, accountId uuid.UUID) error
}

type AccountRecord struct {
	recordId   string
	AccountId  string
	AuthUserId string
	Name       string
	NameNumber uint32
	CreatedAt  time.Time
	DeletedAt  *time.Time
}

type NameCounterRecord struct {
	Name  string
	Count uint32
}
