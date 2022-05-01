package model

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"go-community/util/gcerror"
	"log"
)

func (u User) CountEmail(ctx context.Context, email string) int {
	emailCount := 0
	err := Db.NewSelect().
		ColumnExpr("count(*)").
		Table(UserTable).
		Where("email=(?)", email).
		Scan(ctx, &emailCount)

	if err != nil {
		return 0
	}

	return emailCount
}

func (u *User) Insert(tx bun.Tx, ctx context.Context, nick string, email string, pw string) (*int64, error) {
	*u = User{
		NickName: nick,
		Email:    email,
		Password: pw,
	}

	if _, err := Db.NewInsert().Model(u).Returning("id").Exec(ctx); err != nil {
		err := tx.Rollback()
		log.Printf("tx error : %+v \n", err)
		return nil, gcerror.Error{Message: "계정 저장 시도중 문제가 발생했습니다."}
	}

	if err := tx.Commit(); err != nil {
		return nil, gcerror.Error{Message: "계정 저장 시도중 문제가 발생했습니다.2"}
	}

	return &u.Id, nil
}

type UserService struct {
	*User
	UserRepository UserRepository
}

func (u UserService) Create(ctx context.Context, nick string, email string, pw string) (*int64, error) {
	if ctx == nil {
		return nil, nil
	}
	emailCount := u.UserRepository.CountEmail(ctx, email)

	if emailCount > 0 {
		return nil, gcerror.Error{Message: "already be used"}
	}

	tx, err := Db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Panic(err)
	}

	return u.UserRepository.Insert(tx, ctx, nick, email, pw)
}
