package model_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	m "go-community/model"
	"log"
	"testing"
)

func TestUser_Create(t *testing.T) {
	t.Run("model / user / create", setup(func() {
		// given
		ctx := context.Background()

		nick := "nick"
		email := "email"
		pw := "pw1234"

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepository := m.NewMockUserRepository(ctrl)

		mockUserRepository.EXPECT().
			CountEmail(gomock.Any(), gomock.Any()).
			Return(0)

		u := m.User{
			Id: int64(10),
		}

		mockUserRepository.EXPECT().
			Insert(gomock.Any(), gomock.Any(), gomock.Eq(nick), gomock.Eq(email), gomock.Eq(pw)).
			Return(&u.Id, nil)

		userService := m.UserService{
			User: &u, UserRepository: mockUserRepository,
		}

		// when
		create, err := userService.Create(ctx, nick, email, pw)
		if err != nil {
			log.Panic(err)
		}

		// then
		assert.True(t, *create > int64(0))
	}))
}
