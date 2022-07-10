package entity

import (
	"strconv"
	"time"
)

type User struct {
	ID         int64     `json:"id" db:"id"`
	Age        int64     `json:"age" db:"age"`
	Username   string    `json:"username" db:"username"`
	Password   string    `json:"-" db:"-"`
	DateOfBorn time.Time `json:"dateOfBorn" db:"date_of_born"`
	Fullname   string    `json:"fullName" db:"fullname"`
	Job        string    `json:"job" db:"job"`
	Notes      string    `json:"notes" db:"notes"`
	Status     int64     `json:"status" db:"status"`
	Flag       int64     `json:"flag" db:"flag"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt  time.Time `json:"deletedAt" db:"deleted_at"`
}

type UserParams struct {
	ID       int64  `form:"id" db:"id"`
	Age      int64  `form:"age" db:"age"`
	Fullname string `form:"fullName" db:"fullname"`
	Job      string `form:"job" db:"job"`
	Limit    int64  `form:"limit" db:"-"`
	Page     int64  `form:"page" db:"-"`
}

func (up *UserParams) GenerateQuery(query string, paginate bool) string {
	if up.ID > 0 {
		query = query + " AND id=" + strconv.Itoa(int(up.ID))
	}
	if up.Age > 0 {
		query = query + " AND age=" + strconv.Itoa(int(up.Age))
	}
	if up.Fullname != "" {
		query = query + " AND full_name LIKE '%" + up.Fullname + "%'"
	}
	if up.Job != "" {
		query = query + " AND job LIKE '%" + up.Job + "%'"
	}
	if up.Page < 1 {
		up.Page = 1
	}
	if paginate {
		offset := (up.Page - 1) * up.Limit
		query = query + " LIMIT " + strconv.Itoa(int(up.Limit))
		query = query + " OFFSET " + strconv.Itoa(int(offset))
	}

	return query
}
