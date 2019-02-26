package engine

import (
	"github.com/Authoring/Graph/logger"
	"github.com/Authoring/Graph/pkg/fileutils"
	"gopkg.in/yaml.v2"
)

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

	return e, nil
}
