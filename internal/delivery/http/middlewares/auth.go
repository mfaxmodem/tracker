package middlewares

import (
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt"
    "net/http"
    "strings"
)

type JWTCustomClaims struct {
    UserID int64  `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func AuthMiddleware(secret string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            authHeader := c.Request().Header.Get("Authorization")
            if authHeader == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
            }

            tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
            claims := &JWTCustomClaims{}

            token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })

            if err != nil || !token.Valid {
                return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
            }

            c.Set("user", claims)
            return next(c)
        }
    }
}