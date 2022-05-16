package middleware

import (
	"clickcash_backend/config"
	"clickcash_backend/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"net/http"
	"time"
)
var (
	JwtSecretKey =  []byte(config.GetEnv("jwt.key","jwtcc"))
)
func NewGenerateAccessToken(user string)(string,error)  {
	standardClaims := jwt.StandardClaims{
		Id:        user,
		Issuer:    user,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)
	signedString, err := newWithClaims.SignedString(JwtSecretKey)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return signedString,nil
}
func NewAuthentication(ctx *fiber.Ctx) error {
	return jwtMiddleware.New(jwtMiddleware.Config{
		SigningMethod: "HS256",
		SigningKey:    JwtSecretKey,
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return ctx.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
			})
		},
	})(ctx)
}
