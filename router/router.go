package router

import (
	"github.com/B-P-Coders/backend-template/config"
	"github.com/B-P-Coders/backend-template/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *fiber.App {
	a := fiber.New()

	// Middleware
	a.Use(logger.New())
	a.Use(recover.New())
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(cors.New(cors.Config{
		AllowOrigins: config.Conf.SrvConfig.AllowOrigins,
		AllowMethods: "GET, OPTIONS",
		AllowHeaders: "Origin, Host, Content-Type, Accept",
	}))
	a.Get("/monitor", monitor.New(monitor.Config{
		Title: "6obcy people monitor stats",
	}))

	// Controller
	api := controllers.Controller{Db: db}

	// Routing
	// Connect routes here
	_ = api // Remove later

	return a
}
