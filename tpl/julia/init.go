package julia

import (
	"github.com/PietroCarrara/gulia"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/internal"
	"golang.design/x/thread"
)

const name = "julia"

func init() {
	th = thread.New()
	th.Call(func() {
		gulia.Open()
	})

	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New(d)

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) (interface{}, error) { return ctx, nil },
		}

		ns.AddMethodMapping(
			ctx.Eval,
			[]string{"juliaEval"},
			[][2]string{
				{`{{ (juliaEval "1 + 2 + 3").GetValue }}`, "6"},
				{`{{ (juliaEval "using Plots; plot(sin, -π:0.1:π").GetType }}`, "Plot"},
			},
		)

		ns.AddMethodMapping(
			ctx.Call,
			[]string{"juliaCall"},
			[][2]string{
				{`{{ juliaCall "sin" 1.2 }}`, "0.9320390859672263"},
				{`{{ juliaCall "savefig" $plot "my-plog.png" }}`, ""},
			},
		)

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
