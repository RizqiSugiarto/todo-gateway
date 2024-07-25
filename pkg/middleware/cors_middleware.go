package middleware

import (
	"net/http"
	"strings"
)

var (
	corsAllowedHeaders = []string{
		"Connection", "User-Agent", "Referer",
		"Accept", "Accept-Language", "Content-Type",
		"Content-Language", "Content-Disposition", "Origin",
		"Content-Length", "Authorization", "ResponseType",
		"X-Requested-With", "X-Forwarded-For",
	}
	corsAllowedMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsAllowedOrigins = []string{"*"}
)

func CORS(h http.Handler, env, allowedOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsAllowedMethods, ", "))
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsAllowedHeaders, ", "))
		w.Header().Set("Content-Security-Policy", "object-src 'none'; child-src 'none'; script-src 'unsafe-inline' https: http: ")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Permissions-Policy", "geolocation=()")
		w.Header().Set("Referrer-Policy", "no-referrer")

		if env != "production" {
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(corsAllowedOrigins, ", "))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		}

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
