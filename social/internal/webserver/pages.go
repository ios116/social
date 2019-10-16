package webserver

import "net/http"

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(r.Context().Value("id").(string)))
}

func (s *HttpServer) loginForm(w http.ResponseWriter, r*http.Request) {

}

func (s *HttpServer) registrationForm(w http.ResponseWriter, r*http.Request)  {

}
func (s *HttpServer) userProfile(w http.ResponseWriter, r*http.Request)  {

}
