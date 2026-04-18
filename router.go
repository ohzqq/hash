package hash

import (
	"syscall/js"

	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/jserr"
)

type Router struct {
	*mux.ServeMux
	onLoad mux.Handler
}

func NewRouter() *Router {
	defer jserr.Recover()
	r := &Router{
		ServeMux: mux.NewServeMux(),
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

func (r *Router) ParseURL(uri string) (*mux.Request, error) {
	return r.NewRequest(uri)
}

func (r *Router) Serve() {
	if r.onLoad != nil {
		req, err := r.ServeMux.NewRequest(Get())
		if err != nil {
			jserr.Log(err.Error())
		}
		r.onLoad(req)
	}
	js.Global().Get("window").Call("addEventListener", "hashchange", js.FuncOf(r.routerHandlerFunc))
}

func (r *Router) routerHandlerFunc(this js.Value, args []js.Value) any {
	defer jserr.Recover()
	if !args[0].Truthy() {
		return jserr.New("no args").Value
	}
	e := NewHashEvent(args[0])
	err := r.ServeMux.Serve(e.NewURL().Hash(), e.OldURL().Hash())
	if err != nil {
		return jserr.Wrap(err).Value
	}
	return nil
}
