package mux

import (
	"fmt"
	"testing"
)

var data = [][]string{
	{"/pet/findByStatus", "/pet/findByStatus?param1=toot&param2=poot&param1=root"},
	{"/pet/{petId}", "/pet/113?param1=toot&param2=poot&param1=root"},
	{"/pet/{petId}/info", "/pet/12121/info?param1=toot&param2=poot&param1=root"},
	{"/store/inventory", "/store/inventory?param1=toot&param2=poot&param1=root"},
	{"/store/order/{orderId}", "/store/order/939?param1=toot&param2=poot&param1=root"},
	{"/user/{username}", "/user/1002"},
	{"/user/login", "/user/login?param1=toot&param2=poot&param1=root"},
	{"/user/logout", "/user/logout"},
}

func Test_Match(t *testing.T) {
	// 测试数据data包括 http请求方法，路由规则，客户端请求路径

	mux := NewServeMux()
	for _, v := range data {
		sourceRule := v[0]
		h := func(req *Request) error {
			fmt.Printf("%#v\n", req.URL)
			//println(req.URL.String())
			return nil
		}
		mux.Handle(sourceRule, h)
	}

	for _, v := range data {
		urlPath := v[1]
		//sourceRule := v[0]
		err := mux.Serve(urlPath)
		if err != nil {
			t.Fatalf("route %s, %s\n", urlPath, err.Error())
		}
		//ok, rule, vars := mux.Query(urlPath, method)
		//if ok && rule == sourceRule {
		//  t.Logf("urlPath:%s match rule:%s result: %t vars: %s", urlPath, rule, ok, vars)
		//} else {
		//  t.Errorf("method: %s urlPath:%s match rule:%s result: %t", method, urlPath, sourceRule, ok)
		//}

	}
}
