package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"log"
	"os"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"userId"`
}

func EncodeJwt(id int64) string {
	errFunc := func(err error) {
		log.Panic(err)
	}

	key := []byte(os.Getenv("SECRET"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	errFunc(err)

	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: fmt.Sprintf("%d", id),
		},
		UserId: id,
	}

	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	errFunc(err)

	return token.String()
}

func DecodeJwt(token string) (*UserClaims, error) {
	errFunc := func(err error) {
		log.Panic(err)
	}
	key := []byte(os.Getenv("SECRET"))
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	errFunc(err)

	newToken, _ := jwt.Parse([]byte(token), verifier)
	if err := verifier.Verify(newToken); err != nil {
		return nil, fmt.Errorf("jwd error")
	}

	claims := &UserClaims{}
	_ = json.Unmarshal(newToken.Claims(), claims)

	return claims, nil
}

func GetUserId(ctx context.Context) int64 {
	claims := ctx.Value("claims").(*UserClaims)
	return claims.UserId
}

func GetClaims(ctx context.Context) *UserClaims {
	return ctx.Value("claims").(*UserClaims)
}
