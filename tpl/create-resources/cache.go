package createResources

import (
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/resources/resource_factories/create"
	"github.com/spf13/cast"
)

type Namespace struct {
	create *create.Client
}

func New(d *deps.Deps) *Namespace {
	return &Namespace{
		create: create.New(d.ResourceSpec),
	}
}

func (n *Namespace) FromString(destination, content interface{}) (resource.Resource, error) {
	d, err := cast.ToStringE(destination)
	if err != nil {
		return nil, err
	}

	c, err := cast.ToStringE(content)
	if err != nil {
		return nil, err
	}

	return n.create.FromString(d, c)
}

func (n *Namespace) FromFile(destination, source string) (resource.Resource, error) {
	d, err := cast.ToStringE(destination)
	if err != nil {
		return nil, err
	}

	c, err := cast.ToStringE(source)
	if err != nil {
		return nil, err
	}

	return n.create.FromFile(d, c)
}
