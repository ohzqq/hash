package hash

import (
	"syscall/js"

	"github.com/ohzqq/hash/mux"
)

func OnHashChange(router *mux.ServeMux) {
	h := js.FuncOf(func(this js.Value, args []js.Value) any {
		err := router.Serve(Get())
		if err != nil {
			return "error"
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
