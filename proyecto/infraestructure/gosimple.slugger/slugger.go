package slugger

import (
	"example.com/mittaus/ddd-example/application"
	"github.com/gosimple/slug"
)

type slugger struct{}

func New() application.Slugger {
	return slugger{}
}

func (slugger) NewSlug(initial string) string {
	return slug.Make(initial)
}
