package middleware

import (
	"blog-server/config"
	"blog-server/errs"
	"blog-server/utils"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	jwtUtil *utils.JWTUtil
}

func NewAuthMiddleware(cfg config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: utils.NewJWTUtil(cfg.JWT),
	}
}

func (m *AuthMiddleware) JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    errs.CodeUnauthorized,
				"message": "Authorization header is required",
			})
		}

		tokenString := ""
		if len(authHeader) >= 7 && strings.ToUpper(authHeader[0:7]) == "BEARER " {
			tokenString = authHeader[7:]
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    errs.CodeUnauthorized,
				"message": "Authorization header must start with Bearer",
			})
		}

		claims, err := m.jwtUtil.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    errs.CodeUnauthorized,
				"message": "Invalid token",
			})
		}

		// 将用户信息存储到上下文中，供后续处理程序使用
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}

// 获取当前用户ID
func GetUserID(c *fiber.Ctx) uint {
	if userID, ok := c.Locals("user_id").(uint); ok {
		return userID
	}
	return 0
}

// 获取当前用户邮箱
func GetUserEmail(c *fiber.Ctx) string {
	if email, ok := c.Locals("email").(string); ok {
		return email
	}
	return ""
}

// 获取当前用户名
func GetUserUsername(c *fiber.Ctx) string {
	if username, ok := c.Locals("username").(string); ok {
		return username
	}
	return ""
}