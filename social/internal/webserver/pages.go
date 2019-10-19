package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"social/internal/domain/entities"
	"strconv"
	"strings"
)

func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := s.UserService.GetUsersWithLimitAndOffset(ctx, 200, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := map[string]interface{}{
		"Users": users,
	}
	s.RenderTemplate(ctx, w, "index", data)
}

func (s *HttpServer) loginForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.RenderTemplate(ctx, w, "login", nil)
}

func (s *HttpServer) logOut(w http.ResponseWriter, r *http.Request) {
	s.SessionProvider.DeleteSession(w)
	http.Redirect(w, r, "/", 302)
}

func (s *HttpServer) registrationForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.RenderTemplate(ctx, w, "registration", nil)
}
func (s *HttpServer) userProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["user_id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user, err := s.UserService.GetUserByIdUseCase(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"User": user,
	}
	s.RenderTemplate(ctx, w, "profile", data)
}

func (s *HttpServer) registrationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &entities.User{
		Login:     r.FormValue("login"),
		Password:  r.FormValue("password"),
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		City:      r.FormValue("city"),
		Gender:    r.FormValue("gender"),
		Interests: r.FormValue("interests"),
	}
	id, err := s.UserService.AddUserUseCase(ctx, user)
	if err != nil {
		s.Logger.Error(err.Error())
		http.Redirect(w, r, "/registration/?error", 302)
		return
	}
	userSession := SessionContext{
		ID:    id,
		Login: user.Login,
	}
	err = s.SessionProvider.SetSession(w, userSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func (s *HttpServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))
	user, err := s.UserService.CheckAuthUseCase(ctx, login, password)
	if err != nil {
		http.Redirect(w, r, "/login?error", 302)
		return
	}
	userSession := SessionContext{
		ID:    user.ID,
		Login: user.Login,
	}
	err = s.SessionProvider.SetSession(w, userSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 302)
}
