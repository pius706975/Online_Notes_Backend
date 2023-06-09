package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/pius706975/backend/libs"
)

type UserID string
type Role string

func AuthMiddle(role ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")

			var header string
			var valid bool

			if header = r.Header.Get("Authorization"); header == "" {
				libs.Respond("You need to login first", 401, true).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				libs.Respond("Invalid header", 401, true).Send(w)
				return
			}

			tokens := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := libs.TokenCheck(tokens)
			if err != nil {
				libs.Respond(err.Error(), 401, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkToken.Role {
					valid = true
				}
			}

			if !valid {
				libs.Respond("You do not have permission", 401, true).Send(w)
				return
			}

			log.Println("Auth middleware pass")

			// share user id to controller
			ctx := context.WithValue(r.Context(), UserID("user"), checkToken.UserID)

			// serve next middleware
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
