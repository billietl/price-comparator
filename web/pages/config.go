package pages

import (
	"fmt"
	"path"
	"runtime"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	opts := renderer.Options{
		ParseGlobPattern: fmt.Sprintf("%s/tpl/*.html", path.Dir(filename)),
	}

	rnd = renderer.New(opts)
}
