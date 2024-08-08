package configs

import (
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func isLocalNetwork(ip string) bool {
	ip = strings.TrimSpace(ip)
	privateIPBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"::1/128",
		"fe80::/10",
	}
	for _, block := range privateIPBlocks {
		_, cidr, _ := net.ParseCIDR(block)
		if cidr.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}

func SetupCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			allowedOrigins := []string{"http://localhost:5173", "https://pinjemlah-react-admin.vercel.app"}

			for _, o := range allowedOrigins {
				if o == origin {
					return true
				}
			}

			if isLocalNetwork(origin) {
				return true
			}

			return false
		},
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Cookie",
		AllowCredentials: true,
	}))
}
