package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"  // Add this import
    "github.com/mfaxmodem/tracker/internal/domain/models"
    "github.com/mfaxmodem/tracker/internal/domain/usecase"
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

// Remove this method since it's now in handler.go
// func (h *Handler) RegisterRoutes(e *echo.Echo) {
//     admin := e.Group("/api/v1/admin")
//     
//     // Add registration endpoint
//     admin.POST("/register", h.RegisterAdmin)
//     admin.POST("/login", h.Login)
// }

func (h *Handler) RegisterAdmin(c echo.Context) error {
    var input struct {
        Name     string `json:"name" validate:"required"`
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required,min=6"`
        Role     string `json:"role" validate:"required,oneof=admin"`
    }

    if err := c.Bind(&input); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    if err := c.Validate(&input); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
    }

    user := &models.User{
        Name:         input.Name,
        Email:        input.Email,
        PasswordHash: string(hashedPassword),
        Role:         input.Role,
    }

    if err := h.adminUsecase.CreateUser(user); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
    }

    return c.JSON(http.StatusCreated, map[string]string{
        "message": "Admin registered successfully",
    })
}