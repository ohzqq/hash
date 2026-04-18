package hash

import (
	"syscall/js"

	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/jserr"
	"github.com/ohzqq/tinydom"
)

type Router struct {
	*mux.ServeMux
	onLoad   mux.Handler
	onChange func(e *Event) error
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

func (r *Router) OnChange(h func(e *Event) error) {
	r.onChange = h
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
	js.Global().Get("window").Call("addEventListener", "hashchange", r.serveHandlerFunc())
}

func (r *Router) serveHandlerFunc() js.Func {
	if r.onChange != nil {
		return js.FuncOf(r.hashChangeFunc)
	}
	return js.FuncOf(r.routerHandlerFunc)
}

func (r *Router) hashChangeFunc(this js.Value, args []js.Value) any {
	defer jserr.Recover()
	e := NewHashEvent(args[0])
	err := r.onChange(e)
	if err != nil {
		return jserr.Wrap(err).Value
	}
	return nil
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

type Event struct {
	*tinydom.Event
}

func NewHashEvent(v js.Value) *Event {
	return &Event{tinydom.WrapEvent(v)}
}

func (h *Event) NewURL() *tinydom.URL {
	u := Get()
	if n := h.Get("newURL"); !n.Truthy() {
		u = n.String()
	}
	uri, _ := tinydom.ParseURL(u)
	return uri
}

func (h *Event) OldURL() *tinydom.URL {
	u := Get()
	if n := h.Get("oldURL"); !n.Truthy() {
		u = n.String()
	}
	uri, _ := tinydom.ParseURL(u)
	return uri
}
