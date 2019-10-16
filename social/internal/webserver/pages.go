package webserver

import "net/http"

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(r.Context().Value("id").(string)))
}

func (s *HttpServer) loginForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.RenderTemplate(ctx, w, "login", nil)
}

func (s *HttpServer) registrationForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.RenderTemplate(ctx, w, "registration", nil)
}
func (s *HttpServer) userProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.RenderTemplate(ctx, w, "profile", nil)
}
