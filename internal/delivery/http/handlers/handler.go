package handlers

import (
    "github.com/labstack/echo/v4"
    "tracker/internal/domain/usecase"
)

type Handler struct {
    adminUsecase    usecase.AdminUsecase
    locationUsecase usecase.LocationUsecase
}

func NewHandler(e *echo.Echo, au usecase.AdminUsecase, lu usecase.LocationUsecase) *Handler {
    h := &Handler{
        adminUsecase:    au,
        locationUsecase: lu,
    }

    // Setup routes
    api := e.Group("/api/v1")
    
    // Auth routes
    api.POST("/admin/login", h.Login)
    
    // Admin routes
    admin := api.Group("/admin")
    admin.GET("/visitors", h.GetVisitors)
    admin.POST("/visitors", h.CreateVisitor)
    admin.GET("/stores", h.GetStores)
    admin.POST("/stores", h.CreateStore)
    admin.GET("/routes", h.GetRoutes)
    admin.POST("/routes", h.CreateRoute)

    // Location routes
    api.POST("/location", h.TrackLocation)
    api.GET("/visitor/routes", h.GetVisitorRoutes)

    return h
}