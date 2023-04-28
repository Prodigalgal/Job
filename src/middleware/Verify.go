package middleware

import (
	"Job/src/beans"
	"Job/src/token"
	"github.com/kataras/iris/v12"
)

func VerifyMiddleware() iris.Handler {
	return token.Verifier.Verify(func() interface{} {
		return new(beans.SecurityUser)
	})
}
