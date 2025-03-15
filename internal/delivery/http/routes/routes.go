package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/mfaxmodem/tracker/internal/delivery/http/handlers"
    "github.com/mfaxmodem/tracker/internal/delivery/http/middlewares"
)

func SetupRoutes(e *echo.Echo, h *handlers.Handler) {
    // API v1 group
    v1 := e.Group("/api/v1")
    
    // Admin routes
    admin := v1.Group("/admin")
    admin.Use(middlewares.AdminAuth)
    
    // Admin CRUD operations
    admin.GET("/visitors", h.Admin.ListVisitors)
    admin.POST("/visitors", h.Admin.CreateVisitor)
    admin.PUT("/visitors/:id", h.Admin.UpdateVisitor)
    admin.DELETE("/visitors/:id", h.Admin.DeleteVisitor)
    
    admin.GET("/stores", h.Admin.ListStores)
    admin.POST("/stores", h.Admin.CreateStore)
    admin.PUT("/stores/:id", h.Admin.UpdateStore)
    admin.DELETE("/stores/:id", h.Admin.DeleteStore)
    
    admin.GET("/routes", h.Admin.ListRoutes)
    admin.POST("/routes", h.Admin.CreateRoute)
    admin.PUT("/routes/:id", h.Admin.UpdateRoute)
    admin.DELETE("/routes/:id", h.Admin.DeleteRoute)
    
    // Visitor routes
    visitor := v1.Group("/visitor")
    visitor.Use(middlewares.VisitorAuth)
    
    // Real-time location tracking
    visitor.POST("/location", h.Location.TrackLocation)
    visitor.GET("/routes", h.Location.GetAssignedRoutes)
    
    // WebSocket connection for real-time updates
    v1.GET("/ws", h.WebSocket.Connect)
}