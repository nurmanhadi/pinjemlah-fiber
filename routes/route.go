package routes

import (
	"pinjemlah-fiber/handlers"
	"pinjemlah-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	admin := auth.Group("/admin", middlewares.AuthToken)
	admin.Post("/login", handlers.AdminLogin)
	admin.Post("/logout", handlers.AdminLogout)
	admin.Post("/register", handlers.RegisterAdmin)

	adm := api.Group("/admin", middlewares.AuthToken, middlewares.AdminAuthMiddleware)
	adm.Get("/users", handlers.AdminGetUsers)
	adm.Get("/loans", handlers.AdminGetLoans)
	adm.Get("/payments", handlers.GetPayments)
	adm.Get("/user/:id", handlers.AdminGetUserByID)
	adm.Get("/loan/lunas", handlers.AdminGetLoanStatus)
	adm.Get("/loan/cicilan", handlers.AdminGetLoanCicilan)
	adm.Get("/loan/:id", handlers.AdminGetLoanByID)
	adm.Put("/loan/status/:id", handlers.AdminUpdateStatusLoan)
	adm.Delete("/delete/:id", handlers.DeleteUser)

	// auth
	auth.Post("/register", handlers.RegisterUser)
	auth.Post("/login", handlers.UserLogin)
	auth.Post("/logout", handlers.UserLogout)

	// user
	user := api.Group("/user", middlewares.AuthToken, middlewares.AuthMiddleware)
	user.Post("/loan", handlers.AddLoan)
	user.Get("/loans", handlers.GetLoans)
	user.Get("/loan/:id", handlers.GetLoanByID)
	user.Post("/loan/payment/:id", handlers.LoanPayment)
	user.Get("/Profile", handlers.UserProfile)
	user.Put("/update-Profile", handlers.UserUpdateProfile)
}
