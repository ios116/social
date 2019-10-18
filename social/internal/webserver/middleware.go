package webserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (s *HttpServer) Log(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
func (s *HttpServer) SessionMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userSession, err := s.SessionProvider.GetSession(r)
		if err != nil {
			s.Logger.Info(err.Error())
		}
		ctx := context.WithValue(r.Context(), s.HttpConfig.ContextKey, userSession)
		r = r.WithContext(ctx)
		inner.ServeHTTP(w, r)
	})
}
