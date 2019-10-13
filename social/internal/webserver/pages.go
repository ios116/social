package webserver

import "net/http"

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(r.Context().Value("id").(string)))
}
