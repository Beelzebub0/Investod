package user

import (
	"context"
	"math"

	"github.com/Beelzebub0/Investod/src/business/entity"
	db "github.com/Beelzebub0/Investod/src/lib/database"
	"github.com/pkg/errors"
)

type Interface interface {
	Get(ctx context.Context, uid int64) (entity.User, error)
	GetList(ctx context.Context, params entity.UserParams) ([]entity.User, entity.Pagination, error)
}

type user struct {
	db db.SQL
}

func Init(db db.SQL) Interface {
	u := &user{
		db: db,
	}

	return u
}

func (u *user) Get(ctx context.Context, uid int64) (entity.User, error) {
	result := entity.User{}
	db, err := u.db.Connect()
	if err != nil {
		return result, err
	}
	defer db.Close()

	err = db.QueryRowContext(ctx, readUser, uid).Scan(
		&result.ID,
		&result.Username,
		&result.Fullname,
		&result.Job,
		&result.Age,
	)
	if err != nil {
		return result, err
	}

	return result, err
}

func (u *user) GetList(ctx context.Context, params entity.UserParams) ([]entity.User, entity.Pagination, error) {
	var result []entity.User
	var pagination entity.Pagination

	db, err := u.db.Connect()
	if err != nil {
		return result, pagination, err
	}
	defer db.Close()

	query := params.GenerateQuery(readUserList, true)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return result, pagination, err
	}
	defer rows.Close()

	for rows.Next() {
		r := entity.User{}
		err = rows.Scan(
			&r.ID,
			&r.Username,
			&r.Fullname,
			&r.Age,
			&r.Age,
		)
		if err != nil {
			return result, pagination, err
		}
		result = append(result, r)
	}

	queryCount := params.GenerateQuery(userCount, false)
	err = db.QueryRowContext(ctx, queryCount).Scan(&pagination.TotalElements)
	if err != nil {
		return result, pagination, errors.Wrap(err, "can't exec query")
	}

	pagination.CurrentElements = int64(len(result))
	pagination.CurrentPage = params.Page
	pagination.TotalPages = int64(math.Ceil(float64(pagination.TotalElements) / float64(params.Limit)))

	return result, pagination, err
}
