package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/jwtToken/userToken"
)

func UserMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		errResponse := response.Error(http.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	claims, err := verifyUserToken(token)
	if err != nil {
		errResponse := response.Error(http.StatusUnauthorized, "Invalid token", "Invalid token", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	ctx.Locals("phone_number", claims.PhoneNumber)
	ctx.Locals("user_id", claims.UserID)
	return ctx.Next()
}

func verifyUserToken(tokenString string) (*userToken.UserClaims, error) {
	//userToken.SecretKey
	token, err := jwt.ParseWithClaims(tokenString, &userToken.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return userToken.SecretUserKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*userToken.UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
