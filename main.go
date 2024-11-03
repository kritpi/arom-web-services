package main

import (
	"context"
	"fmt"
	"log"

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
	app := fiber.New()

	ctx := context.Background()
	cfg := configs.NewConfig()
	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowOrigins: "http://localhost:3000",
	}))

	defer db.Close()

	eventRepo := pg.NewEventPGRepository(db)
	eventService := usecases.ProvideEventService(eventRepo, cfg)
	eventHandler := rest.NewEventHandler(eventService)

	diaryRepo := pg.NewDiaryPGRepository(db)
	diaryService := usecases.ProvideDiaryService(diaryRepo, cfg)
	diaryHandler := rest.NewDiaryHandler(diaryService)

	//Routing
	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello, World! test test 2")
		return c.SendStatus(200)
	})

	app.Post(`/event`, eventHandler.CreateEvent)

	app.Post(`/diary`, diaryHandler.CreateDiary)
	app.Get(`/diary`, diaryHandler.GetAllDiary)
	app.Get(`/diary/date/:date`, diaryHandler.GetDiaryByDate)
	app.Get(`/diary/:id`, diaryHandler.GetDiaryByID)
	app.Get(`/diary/user/:userID`, diaryHandler.GetDiaryByUserID)
	app.Put(`/diary/:id`, diaryHandler.UpdateDiary)


	app.Listen(":8000")
}
