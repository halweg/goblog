package middlewares

import "net/http"

func ForceHTMLMiddleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "text/html; charset=utf-8")
		next.ServeHTTP(writer, request)
	})
}

