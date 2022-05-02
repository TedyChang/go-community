package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

var UserTable string = "gc_user"
var UserAlias string = "u"

type UserRepository interface {
	Insert(tx bun.Tx, ctx context.Context, nick string, email string, pw string) (*int64, error)
	CountEmail(ctx context.Context, email string) int
}

type User struct {
	bun.BaseModel `bun:"table:gc_user,alias:u"`
	Id            int64      `bun:"id,pk,autoincrement"`
	NickName      string     `bun:"nickname"`
	Email         string     `bun:"email"`
	Password      string     `bun:"password"`
	CreateAt      *time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeleteAt      *time.Time `bun:",soft_delete"`
}
