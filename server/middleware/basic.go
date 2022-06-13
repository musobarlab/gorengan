package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/musobarlab/gorengan/shared"
)

type basicAuthconfig struct {
	username string
	password string
}

func NewBasicAuthConfig(username, password string) *basicAuthconfig {
	return &basicAuthconfig{username: username, password: password}
}

func BasicAuth(config *basicAuthconfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		validate := func(username, password string) bool {
			if config.username == username && config.password == password {
				return true
			}
			return false
		}

		auth := strings.SplitN(req.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			shared.BuildJSONResponse(res, shared.Response[shared.EmptyJSON] {
				Success: false,
				Code: http.StatusUnauthorized,
				Message: "authorization failed",
				Data: shared.EmptyJSON{},
			}, http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			shared.BuildJSONResponse(res, shared.Response[shared.EmptyJSON] {
				Success: false,
				Code: http.StatusUnauthorized,
				Message: "authorization failed",
				Data: shared.EmptyJSON{},
			}, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(res, req)
	})
}