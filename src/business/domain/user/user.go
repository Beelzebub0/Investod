package user

import (
	"context"

	"github.com/Beelzebub0/Investod/src/business/entity"
)

type Interface interface {
	Get(ctx context.Context, uid int64) (entity.User, error)
	GetList(ctx context.Context, params entity.UserParams) (entity.User, error)
}
