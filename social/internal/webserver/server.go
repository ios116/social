package webserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"social/internal/config"
	"social/internal/domain/usecase"
	"time"
)

type HttpServer struct {
	UserService     usecase.UserService
	HttpConfig      *config.HttpConf
	GrpcConfig      *config.GrpcConf
	Logger          *zap.Logger
	Templates       map[string]*template.Template
	SessionProvider SessionProvider
}

func NewHttpServer(userService usecase.UserService, httpConfig *config.HttpConf, grpcConfig *config.GrpcConf, logger *zap.Logger) *HttpServer {
	templates := NewTemplates()
	sessionProvider := NewSessionManager(httpConfig.ContextKey, time.Duration(httpConfig.SessionTime))
	return &HttpServer{UserService: userService, HttpConfig: httpConfig, GrpcConfig: grpcConfig, Logger: logger, Templates: templates, SessionProvider: sessionProvider}
}

func (s *HttpServer) RenderTemplate(ctx context.Context, w http.ResponseWriter, templateName string, date map[string]interface{}) {
	tmpl, ok := s.Templates[templateName]
	if !ok {
		http.Error(w, "The html does not exist.", http.StatusInternalServerError)
		return
	}
	//tmpl.Funcs(template.FuncMap{
	//	"User": func() SessionContext { return ctx.Value(s.HttpConfig.ContextKey).(SessionContext)},
	//})
	if date == nil {
		date = make(map[string]interface{})
	}
	date["Session"], ok = ctx.Value(s.HttpConfig.ContextKey).(SessionContext)
	if !ok {
		s.Logger.Error("interface {} is not SessionContext")
	}

	err := tmpl.ExecuteTemplate(w, "base", date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *HttpServer) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/code/assets"))))

	for _, route := range s.Routing() {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func (s *HttpServer) Run() {

	dsn := fmt.Sprintf("%s:%d", s.HttpConfig.Host, s.HttpConfig.Port)
	router := s.NewRouter()
	router.Use(s.Log)
	router.Use(s.SessionMiddleware)
	httpServer := http.Server{
		Addr:    dsn,
		Handler: router,
	}
	s.Logger.Info("Starting web server", zap.String("address", dsn))
	if err := httpServer.ListenAndServe(); err != nil {
		s.Logger.Fatal(err.Error())
	}
}
