package account

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, accountId uuid.UUID) (model.Account, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (model.Account, error)
	Create(ctx context.Context, userId uuid.UUID, name string) error
	DeleteByUserId(ctx context.Context, userId uuid.UUID) error
}
