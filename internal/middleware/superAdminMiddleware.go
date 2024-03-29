package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
	"github.com/hudayberdipolat/go-Ecommerce-backend/pkg/jwtToken/adminToken"
)

func SuperAdminMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		errResponse := response.Error(http.StatusUnauthorized, "Unauthorized", "Unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	claims, err := verifyAdminToken(token)
	if err != nil {
		errResponse := response.Error(http.StatusUnauthorized, "Invalid token", "Invalid token", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(errResponse)
	}
	ctx.Locals("admin_id", claims.AdminID)
	ctx.Locals("phone_number", claims.PhoneNumber)
	ctx.Locals("admin_role", claims.AdminRole)
	ctx.Locals("admin_status", claims.AdminStatus)
	role := ctx.Locals("admin_role")
	if role != "super_admin" {
		return fmt.Errorf("Permission denied")
	}
	return ctx.Next()
}

func verifyAdminToken(tokenString string) (*adminToken.AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &adminToken.AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return adminToken.SecretAdminKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*adminToken.AdminClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
