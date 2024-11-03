package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/kritpi/arom-web-services/configs"
	"github.com/kritpi/arom-web-services/domain/usecases"
	"github.com/kritpi/arom-web-services/internal/adapters/pg"
	"github.com/kritpi/arom-web-services/internal/adapters/rest"
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	cfg := configs.NewConfig()

	// Initialize database connection
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Set up Fiber app
	app := fiber.New()
	setupMiddleware(app)
	setupRoutes(app, db, cfg)

	// Start server with graceful shutdown
	startServer(app)
}

func initDatabase(cfg *configs.Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME,
	))
	if err != nil {
		return nil, err
	}
	log.Println("✅ Database Connected!")
	return db, nil
}

func setupMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowOrigins: "http://localhost:3000",
	}))
}

func setupRoutes(app *fiber.App, db *sqlx.DB, cfg *configs.Config) {
	// Repositories and Use Cases

	// Event Repo
	eventRepo := pg.NewEventPGRepository(db)
	eventService := usecases.ProvideEventService(eventRepo, cfg)
	eventHandler := rest.NewEventHandler(eventService)

	diaryRepo := pg.NewDiaryPGRepository(db)
	diaryService := usecases.ProvideDiaryService(diaryRepo, cfg)
	diaryHandler := rest.NewDiaryHandler(diaryService)

	// User Repo
	userRepo := pg.NewUserPGRepository(db)
	// Pass the entire cfg instead of just the JWT secret
	userUsecase := usecases.ProvideUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userUsecase)

	// Routes
	// Event Routes
	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello, World! test test 2")
		return c.SendStatus(200)
	})

	// Routes
	// Event Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/event", eventHandler.CreateEvent)

	app.Post(`/event`, eventHandler.CreateEvent)

	app.Post(`/diary`, diaryHandler.CreateDiary)
	app.Get(`/diary`, diaryHandler.GetAllDiary)
	app.Get(`/diary/date/:date`, diaryHandler.GetDiaryByDate)
	app.Get(`/diary/:id`, diaryHandler.GetDiaryByID)
	app.Get(`/diary/user/:userID`, diaryHandler.GetDiaryByUserID)
	app.Put(`/diary/:id`, diaryHandler.UpdateDiary)

	app.Listen(":8000")
	// User Routes
	app.Post("/user/register", userHandler.Register) //todo: Insert Successfully but there's error catched ["error": "sql: no rows in result set"]
	app.Post("/user/login", userHandler.Login)
}

func startServer(app *fiber.App) {
	// Run server in a goroutine to allow graceful shutdown
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalf("❌ Error starting server: %v", err)
		}
	}()
	log.Println("🚀 Server running on port 8000")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}
	log.Println("✅ Server exited gracefully")
}
