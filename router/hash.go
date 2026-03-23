package router

import (
	"strings"
	"syscall/js"

	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/tinydom"
)

func OnHashChange(router *mux.ServeMux) {
	h := js.FuncOf(func(this js.Value, args []js.Value) any {
		return router.Serve(strings.TrimPrefix(GetHash(), "#"))
	})
	tinydom.GetWindow().Call("addEventListener", "hashchange", h)
}

func GetHash() string {
	return Location().Get("hash").String()
}

func SetHash(hash string) {
	Location().Set("hash", hash)
}

func Location() *tinydom.Location {
	return tinydom.GetWindow().Location()
}
