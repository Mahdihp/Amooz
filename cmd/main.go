package main

import (
	"Amooz/internal/user/delivery/router"
	"Amooz/pkg/common"
	"Amooz/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	env := os.Getenv("APP_ENV")
	if strings.ToLower(env) == common.KEY_DEVELOPMENT || len(env) <= 0 {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error  loading .env file: %v", err)
		}
	}
}
func main() {
	cfg := config.LoadConfig()
	setupHttpServer(cfg)
}

func setupHttpServer(cfg config.Config) {
	// ایجاد برنامه Fiber
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Mahdi & Mojtaba 👋",
		AppName:       "Amooz",
	})

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	file := setupLogFile()
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			c.JSON(fiber.Map{"status": time.Now().Format(time.RFC3339)})
			return true
		},
		LivenessEndpoint: "/healthcheck",
	}))

	//app.Use(shared.JwtMiddleware)
	app.Use(recover.New())

	//Setup Service
	router.SetupService(cfg, app)

	// شروع سرور
	//log.Println("Starting server on :" + strconv.Itoa(cfg.AppServer.Port) + "...")
	if err := app.Listen(":" + strconv.Itoa(cfg.AppServer.Port)); err != nil {
		log.Fatal(err)
	}
}

func setupLogFile() *os.File {

	file, err := os.OpenFile("./log/log_request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return file
}
