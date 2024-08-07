package handlers

import (
	"os"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	type inputUser struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	var input inputUser
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed hash password"})
	}

	user := models.User{
		Status:      "active",
		PhoneNumber: input.PhoneNumber,
		Password:    string(hash),
	}
	if err := databases.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "registrasi gagal"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "registrasi berhasil",
		"data":    user,
	})
}

func UserLogin(c *fiber.Ctx) error {
	type inputUser struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	var input inputUser
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	var user models.User
	databases.DB.Where("phone_number =?", input.PhoneNumber).First(&user)
	if user.UserID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "wrong password"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "gagal generate token JWT"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "202",
		"message": "berhasil login",
		"token":   tokenString,
	})
}

func UserLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "logout success",
	})
}

func AdminLogin(c *fiber.Ctx) error {
	type inputUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input inputUser
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	var admin models.Admin
	databases.DB.Where("username =?", input.Username).First(&admin)
	if admin.AdminID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "wrong password"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin_id"] = admin.AdminID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "gagal generate token JWT"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "202",
		"message": "berhasil login",
		"token":   tokenString,
	})
}
func AdminLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "200",
		"message": "logout success",
	})
}

func RegisterAdmin(c *fiber.Ctx) error {
	type inputAdmin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input inputAdmin
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "cannot parse JSON"})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed hash password"})
	}

	admin := models.Admin{
		Username: input.Username,
		Password: string(hash),
	}
	if err := databases.DB.Create(&admin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "registrasi gagal"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "registrasi berhasil",
		"data":    admin,
	})
}
