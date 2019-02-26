package status

import "github.com/Authoring/Graph/api/server/router"

type statusRouter struct {
	backend Backend
	routes  []router.Route
}

// NewRouter initializes a new status router
func NewRouter(opts *router.Options) router.Router {
	r := &statusRouter{
		backend: opts.Backend,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routers to the status controller
func (r *statusRouter) Routes() []router.Route {
	return r.routes
}

func (r *statusRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/status", r.getStatus),
	}
}
