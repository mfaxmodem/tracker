package handlers

import (
    "github.com/labstack/echo/v4"
    "github.com/mfaxmodem/tracker/internal/domain/usecase"
)

type Handler struct {
    adminUsecase    usecase.AdminUsecase
    locationUsecase usecase.LocationUsecase
}

func NewHandler(e *echo.Echo, au usecase.AdminUsecase, lu usecase.LocationUsecase) *Handler {
    return &Handler{
        adminUsecase:    au,
        locationUsecase: lu,
    }
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
    // Admin routes
    admin := e.Group("/api/v1/admin")
    admin.POST("/register", h.RegisterAdmin)
    admin.POST("/login", h.Login)
    admin.GET("/visitors", h.GetVisitors)
    admin.POST("/visitors", h.CreateVisitor)
    admin.GET("/stores", h.GetStores)
    admin.POST("/stores", h.CreateStore)
    admin.GET("/routes", h.GetRoutes)
    admin.POST("/routes", h.CreateRoute)

    // Location routes
    location := e.Group("/api/v1/location")
    location.POST("/track", h.TrackLocation)
    location.GET("/visitor/:id", h.GetVisitorLocations)
    location.GET("/visitor/:id/routes", h.GetVisitorRoutesHandler)
}