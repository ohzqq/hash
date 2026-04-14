package hash

import (
	"syscall/js"

	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/jserr"
)

func OnHashChange(router *mux.ServeMux) {
	defer jserr.Recover()
	h := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jserr.Recover()
		jserr.Log(args[0])
		if args[0].Truthy() {
			err := router.Serve(Get())
			if err != nil {
				return jserr.Wrap(err).Value
			}
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
