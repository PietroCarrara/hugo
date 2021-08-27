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

func (n *Namespace) AnkoEnv() *env.Env {
	env := env.NewEnv()
	return core.Import(env)
}

func (n *Namespace) AnkoSet(_env, _name, value interface{}) error {
	env, ok := _env.(*env.Env)
	if !ok {
		return fmt.Errorf("can't use object \"%v\" as an environment", _env)
	}

	name, err := cast.ToStringE(_name)
	if err != nil {
		return err
	}

	return env.Define(name, value)
}

func (n *Namespace) Anko(code interface{}, args ...interface{}) (interface{}, error) {
	out := ""

	var e *env.Env
	var ok bool

	if len(args) == 0 {
		e = n.AnkoEnv()
	} else {
		e, ok = args[0].(*env.Env)
		if !ok {
			return nil, fmt.Errorf("can't use object \"%v\" as an environment", args[0])
		}
	}

	// Redirect the output
	e.Define("print", func(a ...interface{}) {
		out += fmt.Sprint(a...)
	})
	e.Define("println", func(a ...interface{}) {
		out += fmt.Sprintln(a...)
	})
	e.Define("printf", func(format string, a ...interface{}) {
		out += fmt.Sprintf(format, a...)
	})

	codeStr, err := cast.ToStringE(code)
	if err != nil {
		return nil, err
	}

	result, err := vm.Execute(e, nil, codeStr)
	if out != "" {
		return out, err
	}

	return result, err
}
