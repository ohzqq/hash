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
	parts := strings.Split(uri, "?")
	u := &URL{}
	if len(parts) > 0 {
		u.Path = parts[0]
	}
	if len(parts) > 1 {
		u.RawQuery = parts[1]
	}
	return u
}

func (u *URL) Query() query.Values {
	v, _ := query.ParseQuery(u.RawQuery)
	return v
}

func (u *URL) String() string {
	return u.Path + `?` + u.RawQuery
}
