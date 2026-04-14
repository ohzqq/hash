package hash

import (
	"syscall/js"

	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/jserr"
	"github.com/ohzqq/tinydom"
)

type Router struct {
	*mux.ServeMux
	HashChangeHandler func(r *Router) js.Func
	onLoad            mux.Handler
}

func NewRouter() *Router {
	defer jserr.Recover()
	r := &Router{
		ServeMux:          mux.NewServeMux(),
		HashChangeHandler: DefaultHashChangeHandler,
	}
	return r
}

func (r *Router) AddRoute(rule string, h mux.Handler) *Router {
	r.Handle(rule, h)
	return r
}

func (r *Router) OnLoad(h mux.Handler) *Router {
	r.onLoad = h
	return r
}

func (r *Router) SetHashChangeHandler(h func(r *Router) js.Func) *Router {
	r.HashChangeHandler = h
	return r
}

func (r *Router) ParseURL(uri string) (*mux.Request, error) {
	return r.NewRequest(uri)
}

func (r *Router) Serve() {
	if r.onLoad != nil {
		hash := js.Global().Get("window").Get("location").Get("hash").String()
		req, err := r.ServeMux.NewRequest(hash)
		if err != nil {
			jserr.Log(err.Error())
		}
		r.onLoad(req)
	}
	js.Global().Get("window").Call("addEventListener", "hashchange", r.HashChangeHandler(r))
}

func (r *Router) WrapHashEvent(e js.Value) (*mux.Request, error) {
	n, _ := tinydom.ParseURL(e.Get("newURL").String())
	o, _ := tinydom.ParseURL(e.Get("oldURL").String())
	req, err := r.ServeMux.NewRequest(n.Hash(), o.Hash())
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DefaultHashChangeHandler(r *Router) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jserr.Recover()
		//jserr.Log(args[0])
		if args[0].Truthy() {
			req, err := r.WrapHashEvent(args[0])
			if err != nil {
				return jserr.Wrap(err).Value
			}
			jserr.Log("new", req.String(), "old", req.OldURL.String())
			err = r.ServeMux.HandleRequest(req)
			if err != nil {
				return jserr.Wrap(err).Value
			}
		}
		return nil
	})
}
