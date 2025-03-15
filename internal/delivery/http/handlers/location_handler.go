package handlers

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/mfaxmodem/tracker/internal/domain/models"
    "github.com/mfaxmodem/tracker/internal/domain/usecase"
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

// Add these methods to handle location routes

func (h *Handler) GetVisitorLocations(c echo.Context) error {
    visitorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid visitor ID")
    }
    
    locations, err := h.locationUsecase.GetVisitorLocations(visitorID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    return c.JSON(http.StatusOK, locations)
}

func (h *Handler) GetVisitorRoutesHandler(c echo.Context) error {
    visitorID, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid visitor ID")
    }
    
    routes, err := h.locationUsecase.GetVisitorRoutes(visitorID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    
    return c.JSON(http.StatusOK, routes)
}