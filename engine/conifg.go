package engine

import (
	"github.com/Authoring/Graph/logger"
	"github.com/Authoring/Graph/pkg/fileutils"
	"gopkg.in/yaml.v2"
)

type configFinder struct {
	Name string
}

// LoadDefault load the default config
func (e *Engine) LoadDefault() {
	baseConfig, err := fileutils.ReadFileInBase("config")

	if err != nil {
		logger.L.Panicf("Seems like the database is missing, perhaps you forgot to init it? %v", err)
	}

	var base configFinder
	err = yaml.Unmarshal(baseConfig, &base)
	if err != nil {
		logger.L.Panicf("Seems like the database is corrpt")
	}

	e.LoadOrCreate(base.Name)
}

// LoadOrCreate loads or creates a new engine
func (e *Engine) LoadOrCreate(name string) (*Engine, error) {
	f, err := fileutils.LoadAndCreate(name, "core")

	if err != nil {
		logger.L.Fatalf("Unable to load %s", name)
	}

	err = yaml.Unmarshal(f, e)

	if err != nil {
		logger.L.Errorf("Corrupt file %s.yaml, %v", name, err)
		return nil, err
	}

	loadInitData(e, name)
	raw, err := yaml.Marshal(e)

	if err != nil {
		return nil, err
	}

	fileutils.WriteFile(name, "core", raw)
	base, filepath := fileutils.GetBaseAndFilePath(name, "core")
	logger.L.Infof("%s Created", base)
	logger.L.Infof("%s Created", filepath)

	initDefault(e)

	return e, nil
}

func initDefault(e *Engine) {
	var p = &configFinder{
		Name: e.Name,
	}
	o, err := yaml.Marshal(p)

	if err != nil {
		logger.L.Fatalf("Unable to save config.yml")
	}

	fileutils.WriteFileInBase("config", o)
}
