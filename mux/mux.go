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
	req := NewRequest(uri)
	ok, rule, vars := mux.tree.Query(req.Path)
	if !ok {
		return nil, errors.New("error parsing request url " + uri + " rule " + rule)
	}
	req.Rule = rule
	req.PathVars = vars
	return req, nil
}

func (mux *ServeMux) Serve(uri string) error {
	req, err := mux.NewRequest(uri)
	if err != nil {
		return err
	}
	if h, ok := mux.handlers[req.Rule]; ok {
		return h(req)
	}
	return errors.New("serve error")
}

type Request struct {
	*url.URL
	Rule     string
	PathVars map[string]string
}

func NewRequest(uri string) *Request {
	return &Request{
		URL: url.Parse(uri),
	}
}

func (req *Request) QueryValue(k string) string {
	return req.Query().Get(k)
}

func (req *Request) PathValue(k string) string {
	if len(req.PathVars) < 1 {
		return ""
	}
	if val, ok := req.PathVars[k]; ok {
		return val
	}
	return ""
}
