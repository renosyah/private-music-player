package auth

import (
	"context"
	"net/http"
)

type (
	Option struct {
		Enable bool
	}
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Cache-Control", "max-age=0")

		if !option.Enable {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "POST")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, session, X-Requested-With")
		}

		if option.Enable {
			sessionID := r.Header.Get("session")
			if sessionID == "" {
				RespondError(w, "Unauthorized Access", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "session", sessionID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
