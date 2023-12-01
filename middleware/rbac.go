package middleware

import (
	"net/http"

	"github.com/DitoAdriel99/go-monsterdex/pkg/jwt_parse"
	"github.com/labstack/echo/v4"
)

func RBAC(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearer := c.Request().Header.Get("Authorization")
			claims, err := jwt_parse.GetClaimsFromToken(bearer)
			if err != nil {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"Message": err,
				})
			}
			userRole := claims.Role

			// Check if the user's role matches any of the specified roles
			roleMatch := false
			for _, role := range roles {
				if userRole == role {
					roleMatch = true
					break // Exit the loop if there's a match
				}
			}

			// If the user's role doesn't match any of the specified roles, deny access
			if !roleMatch {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"Message": "Access Denied",
				})
			}

			// Proceed to the next handler if access is granted
			return next(c)
		}
	}
}
