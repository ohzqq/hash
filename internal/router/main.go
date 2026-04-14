package main

import (
	"syscall/js"

	"github.com/ohzqq/hash"
	"github.com/ohzqq/hash/mux"
	"github.com/ohzqq/tinydom"
)

var data = [][]string{
	{"#/pet/findByStatus", "/pet/findByStatus?param1=toot&param2=poot&param1=root"},
	{"#/pet/{petId}", "/pet/113?param1=toot&param2=poot&param1=root"},
	{"#/pet/{petId}/info", "/pet/12121/info?param1=toot&param2=poot&param1=root"},
	{"#/store/inventory", "/store/inventory?param1=toot&param2=poot&param1=root"},
	{"#/store/order/{orderId}", "/store/order/939?param1=toot&param2=poot&param1=root"},
	{"#/user/{username}", "/user/1002"},
	{"#/user/login", "/user/login?param1=toot&param2=poot&param1=root"},
	{"#/user/logout", "/user/logout"},
}

func main() {
	rmux := hash.NewRouter()
	for _, v := range data {
		sourceRule := v[0]
		rmux.AddRoute(sourceRule, wrapHandler())
	}
	rmux.OnLoad(wrapHandler())
	rmux.Serve()
	//hash.OnHashChange(rmux)
	select {}
}

func wrapHandler() mux.Handler {
	return func(req *mux.Request) error {
		js.Global().Get("console").Call("log", "hashchange", req.URL.String())
		div := tinydom.GetDocument().GetElementById("test")
		div.SetTextContent(req.Rule)
		return nil
	}
}
