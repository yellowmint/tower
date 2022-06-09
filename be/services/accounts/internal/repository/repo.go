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
	GetAccountById(ctx context.Context, playerId uuid.UUID) (AccountRecord, error)
	GetAccountByAuthUserId(ctx context.Context, userId uuid.UUID) (AccountRecord, error)
	CreateAccount(ctx context.Context, userId uuid.UUID, name string) error
	DeleteAccountByAuthUserId(ctx context.Context, userId uuid.UUID) error
}

type AccountRecord struct {
	recordId   string
	AccountId  uuid.UUID
	AuthUserId uuid.UUID
	Name       string
	CreatedAt  time.Time
	DeletedAt  *time.Time
}