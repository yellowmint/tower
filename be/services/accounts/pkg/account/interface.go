package account

import (
	"context"
	"git.jetbrains.space/artdecoction/wt/tower/services/accounts/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, accountId uuid.UUID) (model.Account, error)
	GetByAuthUserId(ctx context.Context, authUserId string) (model.Account, error)
	Create(ctx context.Context, authUserId, name string) error
	DeleteId(ctx context.Context, accountId uuid.UUID) error
}
