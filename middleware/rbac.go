package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"
	"github.com/labstack/echo/v4"
)

type RBAC struct {
	Token tokenizer.JWT
}

func (r *RBAC) Validate(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("claimss")
			bearer := c.Request().Header.Get("Authorization")
			claims, err := r.Token.GetClaimsFromToken(bearer)
			if err != nil {
				log.Println("error get claims from token", err)
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
				log.Println("Access Denied")
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"Message": "Access Denied",
				})
			}

			// Proceed to the next handler if access is granted
			return next(c)
		}
	}
}
