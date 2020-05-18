package middleware

import (
	"github.com/omc-college/management-system/pkg/rbac/opa"
	"net/http"
)

type AuthorizationMiddleware struct{}

func (middleware *AuthorizationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var isAccessGranted bool
		var requestRegoInput opa.RegoInput

		requestRegoInput = opa.RegoInput{
			Path:   r.URL.Path,
			Method: r.Method,
			Token:  r.Header.Get("Authorization"),
		}

		isAccessGranted = opa.GetDecision(requestRegoInput)

		if !isAccessGranted {
			http.Error(w, "Access is not granted", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
