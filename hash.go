package hash

import (
	"strings"
	"syscall/js"

	"github.com/ohzqq/hash/mux"
)

func OnHashChange(router *mux.ServeMux) {
	h := js.FuncOf(func(this js.Value, args []js.Value) any {
		return router.Serve(strings.TrimPrefix(GetHash(), "#"))
	})
	js.Global().Get("window").Call("addEventListener", "hashchange", h)
}

func GetHash() string {
	return Location().Get("hash").String()
}

func SetHash(hash string) {
	Location().Set("hash", hash)
}

func Location() js.Value {
	return js.Global().Get("window").Get("location")
}
