package domain

import (
	"github.com/Beelzebub0/Investod/src/business/domain/user"
	db "github.com/Beelzebub0/Investod/src/lib/database"
)

type Domains struct {
	User user.Interface
}

func Init(db db.SQL) *Domains {
	d := &Domains{
		User: user.Init(db),
	}

	return d
}
