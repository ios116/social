package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
	"strconv"
	"strings"
)

// Index - main page
func (s *HttpServer) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := s.UserService.GetUsersWithLimitAndOffset(ctx, 200, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := map[string]interface{}{
		"Users":  users,
		"Errors": "",
	}
	s.RenderTemplate(ctx, w, "index", data)
}

// Search by first name and last name with limit
func (s *HttpServer) Search(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := r.URL.Query().Get("id")
	if id == "" {
		id="0"
	}
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, exceptions.IntegerRequired.Error(), 500)
		return
	}

	direction := r.URL.Query().Get("direction")

	query := r.FormValue("query")
	users, err := s.UserService.FindByNameUC(ctx, query, id64, 21, direction)
	var firstID, lastID int64

	data := map[string]interface{}{
		"Users":  users,
		"Errors": "",
		"Query":  query,
	}

	if err != nil {
		data["Errors"] = err.Error()
		s.RenderTemplate(ctx, w, "index", data)
		return
	}

	if len(users) > 0 {
		lastID = users[len(users)-1].ID
		firstID = users[0].ID
	}

	data["FirstID"] = firstID
	data["LastID"] = lastID

	s.RenderTemplate(ctx, w, "index", data)

}

// loginForm user enter credentials
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
