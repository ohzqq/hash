package mux

import (
	"errors"

	"github.com/ohzqq/hash/url"
	"github.com/ohzqq/pretree"
)

type ServeMux struct {
	tree     *pretree.PreTree
	handlers map[string]Handler
}

type Handler func(req *Request) error

func NewServeMux() *ServeMux {
	return &ServeMux{
		tree:     pretree.New(),
		handlers: make(map[string]Handler),
	}
}

func (mux *ServeMux) Handle(rule string, cb Handler) {
	mux.tree.GET(rule)
	mux.handlers[rule] = cb
}

func (mux *ServeMux) NewRequest(uri string) (*Request, error) {
	ok, rule, vars := mux.Query(uri)
	if !ok {
		return nil, errors.New("error parsing request url")
	}
	req := NewRequest(uri)
	req.rule = rule
	req.PathVars = vars
	return req, nil
}

func (mux *ServeMux) Serve(uri string) error {
	req, err := mux.NewRequest(uri)
	if h, ok := mux.handlers[req.rule]; ok {
		return h(req)
	}
	return errors.New("serve error")
}

type Request struct {
	*url.URL
	rule     string
	PathVars map[string]string
}

func NewRequest(uri string) *Request {
	return &Request{
		URL: url.Parse(uri),
	}
}
