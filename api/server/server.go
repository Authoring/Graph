package server

import (
	"fmt"
	"net/http"

	"github.com/Authoring/Graph/api/server/httputils"
	"github.com/Authoring/Graph/api/server/router"
	"github.com/Authoring/Graph/api/server/router/status"
	"github.com/Authoring/Graph/logger"
	"github.com/gorilla/mux"
)

// Server defines a new server
type Server struct {
	routers []router.Router
	opts    *router.Options
}

// InitServer initialize the server
func InitServer(opts *router.Options) {
	if opts.Verbose {
		logger.L.Infof("Initilizing the server")
	}

	s := Server{
		routers: initRouter(opts),
		opts:    opts,
	}

	s.serve()
}

func initRouter(opts *router.Options) []router.Router {
	return []router.Router{
		status.NewRouter(opts),
	}
}

func (s *Server) serve() {

	if s.opts.Verbose {
		logger.L.Infof("Listening on 0.0.0.0:%d", s.opts.Backend.Port)
	}

	r := s.createMux()
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.opts.Backend.Port), r)
	if err != nil {
		logger.L.Errorf("Unable to start server: %v", err)
	}
}

// createMux create a new mux router
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()

	if s.opts.Verbose {
		logger.L.Infof("Registering routes")
	}

	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {

			if s.opts.Verbose {
				logger.L.Infof("Registering route: [%s] %s", r.Method(), r.Path())
			}

			f := s.makeHTTPHandler(r.Handler())
			m.Path(r.Path()).Methods(r.Method()).Handler(f)
		}
	}

	return m
}

// makeHTTPHandler make a http handler from the APIFunc
func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		if s.opts.Verbose {
			logger.L.Infof("[%s] %s", r.Method, r.RequestURI)
		}

		if vars == nil {
			vars = make(map[string]string)
		}

		vals := r.URL.Query()
		for k, v := range vals {
			if len(v) > 0 {
				vars[k] = v[0]
			}
		}

		err := handler(ctx, w, r, vars)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
