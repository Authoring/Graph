package router

import (
	"github.com/Authoring/Graph/api/server/httputils"
	"github.com/Authoring/Graph/daemon"
)

// Options defines router options
type Options struct {
	Backend *daemon.Daemon
	Verbose bool
}

// Router defines a group of routes
type Router interface {
	Routes() []Route
}

// Route defines an individual API endpoint
type Route interface {
	Handler() httputils.APIFunc
	Method() string
	Path() string
}
