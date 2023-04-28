package token

import (
	"Job/src/beans"
	"bytes"
	"github.com/kataras/iris/v12"
	jwtm "github.com/kataras/iris/v12/middleware/jwt"
	"strconv"
	"time"
)

const (
	accessTokenMaxAge  = 5 * time.Hour
	refreshTokenMaxAge = 10 * time.Hour
)

var (
	Secret   = []byte("56A5A3GH5CBGJ6")
	Signer   = jwtm.NewSigner(jwtm.HS256, Secret, accessTokenMaxAge)
	Verifier = jwtm.NewVerifier(jwtm.HS256, Secret)
)

func init() {
	Verifier.WithDefaultBlocklist()
	Verifier.Extractors = []jwtm.TokenExtractor{jwtm.FromJSON("access_token"), jwtm.FromJSON("refresh_token")}
}

func GenerateTokenPair(user *beans.User) (*jwtm.TokenPair, error) {
	refreshClaims := jwtm.Claims{Subject: strconv.Itoa(user.Id)}

	accessClaims := beans.SecurityUser{
		ID:       user.Id,
		Username: user.Username,
	}

	tokenPair, err := Signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		return nil, err
	}

	return &tokenPair, nil
}

func RefreshToken(user *beans.User, tokenPair *jwtm.TokenPair) (*jwtm.TokenPair, error) {
	tokenPair.RefreshToken = bytes.ReplaceAll(tokenPair.RefreshToken, []byte("\""), []byte(""))

	_, err := Verifier.VerifyToken(tokenPair.RefreshToken, jwtm.Expected{Subject: strconv.Itoa(user.Id)})
	if err != nil {
		return nil, err
	}

	return GenerateTokenPair(user)
}

func GetAccessToken(ctx iris.Context) *beans.SecurityUser {
	return jwtm.Get(ctx).(*beans.SecurityUser)
}
