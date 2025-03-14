package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tracker/internal/domain/models"
    "tracker/internal/domain/usecase"
)

type AdminHandler struct {
    adminUsecase usecase.AdminUsecase
}

func NewAdminHandler(au usecase.AdminUsecase) *AdminHandler {
    return &AdminHandler{
        adminUsecase: au,
    }
}

func (h *Handler) GetVisitors(c echo.Context) error {
    visitors, err := h.adminUsecase.GetAllVisitors()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, visitors)
}

func (h *Handler) CreateVisitor(c echo.Context) error {
    var visitor models.User
    if err := c.Bind(&visitor); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    if err := h.adminUsecase.CreateVisitor(&visitor); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, visitor)
}

func (h *Handler) GetStores(c echo.Context) error {
    stores, err := h.adminUsecase.GetAllStores()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, stores)
}

func (h *Handler) CreateStore(c echo.Context) error {
    var store models.Store
    if err := c.Bind(&store); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    if err := h.adminUsecase.CreateStore(&store); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, store)
}

func (h *Handler) GetRoutes(c echo.Context) error {
    routes, err := h.adminUsecase.GetAllRoutes()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, routes)
}

func (h *Handler) CreateRoute(c echo.Context) error {
    var route models.Route
    if err := c.Bind(&route); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    
    if err := h.adminUsecase.CreateRoute(&route); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, route)
}