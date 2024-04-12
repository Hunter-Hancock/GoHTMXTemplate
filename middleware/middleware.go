package mw

import (
	"context"
	"net/http"
)

type MiddleWare struct{}

func (mw *MiddleWare) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userID", 123)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
