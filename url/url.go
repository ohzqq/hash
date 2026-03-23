package url

import (
	"strings"

	"github.com/ohzqq/query"
)

type URL struct {
	Path     string
	RawQuery string
}

func Parse(uri string) *URL {
	p, q := strings.Split(uri, "?")
	return &URL{
		Path:     p,
		RawQuery: q,
	}
}

func (u *URL) Query() query.Values {
	v, _ := query.ParseQuery(u.RawQuery)
	return v
}

func (u *URL) String() string {
	return u.Path + `?` + u.RawQuery
}
