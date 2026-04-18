package hash

import (
	"syscall/js"

	"github.com/ohzqq/jserr"
)

func OnHashChange(onChange func(e *Event) error) {
	defer jserr.Recover()
	h := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jserr.Recover()
		err := onChange(NewHashEvent(args[0]))
		if err != nil {
			return jserr.Wrap(err).Value
		}
		return nil
	})
	js.Global().Get("window").Call("addEventListener", "hashchange", h)
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
