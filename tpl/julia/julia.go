package julia

import (
	"fmt"

	"github.com/PietroCarrara/gulia"
	"github.com/gohugoio/hugo/deps"
	"github.com/spf13/cast"
	"golang.design/x/thread"
)

type Namespace struct {
	deps *deps.Deps
}

var th thread.Thread

func New(d *deps.Deps) *Namespace {
	return &Namespace{
		deps: d,
	}
}

func (n *Namespace) Eval(code interface{}) (gulia.JuliaValue, error) {
	str, err := cast.ToStringE(code)
	if err != nil {
		return nil, err
	}

	var res gulia.JuliaValue
	th.Call(func() {
		res, err = gulia.EvalString(str)
	})

	return res, err
}

func (n *Namespace) Call(name interface{}, args ...interface{}) (gulia.JuliaValue, error) {
	str, err := cast.ToStringE(name)
	if err != nil {
		return nil, err
	}

	var f gulia.JuliaFunction
	th.Call(func() {
		f, err = gulia.GetFunction(str)
	})
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, fmt.Errorf(`function "%s" not found`, str)
	}

	var res gulia.JuliaValue
	th.Call(func() {
		res, err = f.Call(args...)
	})

	return res, err
}
