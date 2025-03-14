package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tracker/internal/domain/models"
    "tracker/internal/domain/usecase"
)

type LocationHandler struct {
    locationUsecase usecase.LocationUsecase
}

func NewLocationHandler(lu usecase.LocationUsecase) *LocationHandler {
    return &LocationHandler{
        locationUsecase: lu,
    }
}

func (h *Handler) TrackLocation(c echo.Context) error {
    var location models.Location
    if err := c.Bind(&location); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    if err := h.locationUsecase.TrackLocation(&location); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "Location tracked successfully",
    })
}

func (h *Handler) GetVisitorRoutes(c echo.Context) error {
    visitorID := c.Get("user_id").(int64)
    
    routes, err := h.locationUsecase.GetVisitorRoutes(visitorID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, routes)
}