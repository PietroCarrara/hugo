package createResources

import (
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/internal"
)

const name = "createResource"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New(d)

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) (interface{}, error) { return ctx, nil },
		}

		ns.AddMethodMapping(
			ctx.FromString,
			[]string{"fromString"},
			[][2]string{
				{`{{ fromString "0.12342" }}`, ""},
			},
		)

		ns.AddMethodMapping(
			ctx.FromFile,
			[]string{"fromFile"},
			[][2]string{
				{`{{ fromFile "/tmp/generated_img.png" }}`, ""},
			},
		)

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
