package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// BasicAuth function basic auth
func BasicAuth(user, pass string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {

		validate := func(u, p string) bool {
			if user == u && pass == p {
				return true
			}
			return false
		}

		if validate(username, password) {
			return true, nil
		}
		return false, nil
	})
}
