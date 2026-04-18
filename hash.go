package hash

import (
	"syscall/js"

	"github.com/ohzqq/jserr"
	"github.com/ohzqq/tinydom"
)

func OnChange(onChange Handler) {
	defer jserr.Recover()
	js.Global().Get("window").Call("addEventListener", "hashchange", WrapHandler(onChange))
}

func Get() string {
	return Location().Get("hash").String()
}

func Set(hash string) {
	Location().Set("hash", hash)
}

func Location() js.Value {
	return js.Global().Get("window").Get("location")
}

type Event struct {
	*tinydom.Event
}

func newEvent() *Event {
	v := map[string]any{
		"oldURL": Get(),
		"newURL": Get(),
	}
	return NewHashEvent(js.ValueOf(v))
}

func NewHashEvent(v js.Value) *Event {
	return &Event{tinydom.WrapEvent(v)}
}

func (h *Event) NewURL() *tinydom.URL {
	u := Get()
	if n := h.Get("newURL"); n.Truthy() {
		u = n.String()
	}
	uri, _ := tinydom.ParseURL(u)
	return uri
}

func (h *Event) OldURL() *tinydom.URL {
	u := Get()
	if n := h.Get("oldURL"); n.Truthy() {
		u = n.String()
	}
	uri, _ := tinydom.ParseURL(u)
	return uri
}

type Handler func(e *Event) error

func WrapHandler(h Handler) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jserr.Recover()
		err := h(NewHashEvent(args[0]))
		if err != nil {
			return jserr.Wrap(err).Value
		}
		return nil
	})
}
