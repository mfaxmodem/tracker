package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/mfaxmodem/tracker/internal/config"
    "github.com/mfaxmodem/tracker/internal/delivery/http/handlers"
    "github.com/mfaxmodem/tracker/internal/domain/usecase"
    "github.com/mfaxmodem/tracker/internal/repository/postgres"
    "github.com/mfaxmodem/tracker/pkg/validator"
)

func main() {
    e := echo.New()

    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        e.Logger.Fatal(err)
    }

    // Set up validator
    e.Validator = validator.NewValidator()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    // Database connection
    db, err := sql.Open("postgres", cfg.GetDBConnString())
    if err != nil {
        e.Logger.Fatal(err)
    }
    defer db.Close()

    // Initialize components
    repo := postgres.NewRepository(db)
    adminUsecase := usecase.NewAdminUsecase(repo)
    locationUsecase := usecase.NewLocationUsecase(repo)
    
    // Initialize handlers and register routes
    handler := handlers.NewHandler(e, adminUsecase, locationUsecase)
    handler.RegisterRoutes(e)  // اضافه کردن این خط برای ثبت مسیرها

    // Start server
    e.Logger.Fatal(e.Start(":" + cfg.APIPort))
}
