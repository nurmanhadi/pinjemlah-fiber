package middlewares

import (
	"os"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	if token != os.Getenv("SECRET_TOKEN") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired JWT"})
	}

	return c.Next()
}

func ExctractToken(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return token, err
}

func AuthMiddleware(c *fiber.Ctx) error {
	token, err := ExctractToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	databases.DB.First(&user, claims["user_id"])
	if user.UserID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	c.Locals("user", user)
	return c.Next()
}
func AdminAuthMiddleware(c *fiber.Ctx) error {
	token, err := ExctractToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)
	var admin models.Admin
	databases.DB.First(&admin, claims["admin"])
	if admin.AdminID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	c.Locals("admin", admin)
	return c.Next()
}
