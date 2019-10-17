package webserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"html/template"
	"net/http"
	"social/internal/config"
	"social/internal/domain/usecase"
	gw "social/internal/grpcserver"
	"social/internal/webserver/middleware"
)

type HttpServer struct {
	UserService usecase.UserService
	HttpConfig  *config.HttpConf
	GrpcConfig  *config.GrpcConf
	Logger      *zap.Logger
	Templates   map[string]*template.Template
}

func NewHttpServer(userService usecase.UserService, httpConfig *config.HttpConf, grpcConfig *config.GrpcConf, logger *zap.Logger) *HttpServer {
	templates:= NewTemplates()
	return &HttpServer{UserService: userService, HttpConfig: httpConfig, GrpcConfig: grpcConfig, Logger: logger, Templates: templates}
}

func (s *HttpServer) RenderTemplate(ctx context.Context, w http.ResponseWriter, templateName string, date interface{}) {
	tmpl, ok := s.Templates[templateName]
	if !ok {
		http.Error(w, "The html does not exist.", http.StatusInternalServerError)
		return
	}
	err := tmpl.ExecuteTemplate(w ,"base",date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type tokenAuth struct {
	Token string
}

func (t *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.Token,
	}, nil
}

func (t *tokenAuth) RequireTransportSecurity() bool {
	return false
}

// GrpcHandler connect to grpc server
func (s *HttpServer) GrpcHandler(ctx context.Context) (http.Handler, error) {
	addressRpc := fmt.Sprintf("%s:%d", s.GrpcConfig.GrpcHost, s.GrpcConfig.GrpcPort)
	option := grpc.WithPerRPCCredentials(&tokenAuth{"Bearer secret"})
	conn, err := grpc.Dial(addressRpc, option, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	rpcGWMux := runtime.NewServeMux()

	err = gw.RegisterUsersHandler(ctx, rpcGWMux, conn)
	if err != nil {
		return nil, err
	}
	return rpcGWMux, nil
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
	//ctx := context.Background()
	//rpcHandler, err := s.GrpcHandler(ctx)
	//if err != nil {
	//	s.Logger.Fatal(err.Error())
	//}

	router := s.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.SessionMiddleware)
	//router.PathPrefix("/v1").Handler(rpcHandler)
	httpServer := http.Server{
		Addr:    dsn,
		Handler: router,
	}
	s.Logger.Info("Starting web server", zap.String("address", dsn))
	if err := httpServer.ListenAndServe(); err != nil {
		s.Logger.Fatal(err.Error())
	}
}


