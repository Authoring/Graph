package engine

import (
	"time"

	"github.com/Authoring/Graph/pkg/fileutils"
)

// CurrentEngineVersion defines the current version fo the engine
var CurrentEngineVersion = "1.0.0"

// Init defines the methods for the engine interface
type Init interface {
	LoadOrCreate(string) (*Engine, error)
	LoadDefault()
}

// Engine defines the engine object
type Engine struct {
	Created time.Time
	Version string
	Name    string
	Path    string
}

func loadInitData(e *Engine, name string) {
	if len(e.Version) == 0 {
		e.Version = CurrentEngineVersion
		e.Created = time.Now()
		e.Name = name
	}
	p, _ := fileutils.GetBaseAndFilePath(name, "core")
	e.Path = p
}
