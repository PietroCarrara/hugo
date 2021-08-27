package script

import (
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/internal"
)

const name = "script"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New(d)

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) (interface{}, error) { return ctx, nil },
		}

		ns.AddMethodMapping(
			ctx.AnkoEnv,
			[]string{"ankoEnv"},
			[][2]string{},
		)

		ns.AddMethodMapping(
			ctx.AnkoSet,
			[]string{"ankoSet"},
			[][2]string{
				{`{{ ankoSet $env "variable" 1 }}`, ""},
				{`{{ ankoSet $env "msg" "Hello, World!" }}`, ""},
			},
		)

		ns.AddMethodMapping(
			ctx.Anko,
			[]string{"anko"},
			[][2]string{
				{`{{ anko "1 + 1" }}`, "1"},
				{`{{ anko "print(\"https://github.com/mattn/anko\")" $env }}`, "1"},
			},
		)

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
