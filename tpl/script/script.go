package script

import (
	"fmt"

	"github.com/gohugoio/hugo/deps"
	"github.com/mattn/anko/core"
	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
	"github.com/spf13/cast"
)

type Namespace struct {
	deps *deps.Deps
}

func New(d *deps.Deps) *Namespace {
	return &Namespace{
		deps: d,
	}
}

func (n *Namespace) Anko(code interface{}) (interface{}, error) {
	out := ""

	env := env.NewEnv()
	core.Import(env)

	// Redirect the output
	env.Define("print", func(a ...interface{}) {
		out += fmt.Sprint(a...)
	})
	env.Define("println", func(a ...interface{}) {
		out += fmt.Sprintln(a...)
	})
	env.Define("printf", func(format string, a ...interface{}) {
		out += fmt.Sprintf(format, a...)
	})

	codeStr, err := cast.ToStringE(code)
	if err != nil {
		return nil, err
	}

	result, err := vm.Execute(env, nil, codeStr)
	if out != "" {
		return out, err
	}

	return result, err
}
