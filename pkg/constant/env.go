package constant

import (
	"go-restful/pkg/util"
	"os"
	"path/filepath"
)

var (
	basePath, _ = os.Getwd()
	Env, _      = util.NewEnv(filepath.Join(basePath, ".env"))
)
