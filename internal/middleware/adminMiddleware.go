package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-Ecommerce-backend/internal/utils/response"
)

func AdminMiddleware(ctx *fiber.Ctx) error {
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
	if role == "super_admin" || role == "admin" {
		return ctx.Next()
	}
	return fmt.Errorf("Permission denied")
}
