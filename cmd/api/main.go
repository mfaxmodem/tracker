package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "tracker/internal/config"
    "tracker/internal/delivery/http/handlers"
    "tracker/internal/domain/usecase"
    "tracker/internal/repository/postgres"
)

func main() {
    e := echo.New()

    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        e.Logger.Fatal(err)
    }

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
    handlers.NewHandler(e, adminUsecase, locationUsecase)

    // Start server
    e.Logger.Fatal(e.Start(":" + cfg.APIPort))
}
