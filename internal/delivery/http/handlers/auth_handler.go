package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt"
    "time"
)

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *Handler) Login(c echo.Context) error {
    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    user, err := h.adminUsecase.GetUserByEmail(req.Email)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
    }

    // TODO: Add proper password hashing and verification
    if req.Password != "admin123" || user.Role != "admin" {
        return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
    }

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["email"] = user.Email
    claims["role"] = user.Role
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token
    t, err := token.SignedString([]byte("your-super-secret-key-change-this-in-production"))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}