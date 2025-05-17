package main

import (
	"fmt"
	"encoding/json"
	"github.com/scientificworld/frpsp"
)

func login(o *frpsp.LoginObject) {
	req, _ := json.Marshal(o.Req)
	fmt.Println("Login", string(req))
	o.Resolve(true)
}

func newproxy(o *frpsp.NewProxyObject) {
	req, _ := json.Marshal(o.Req)
	fmt.Println("NewProxy", string(req))
	if o.Req.User.Metas["token"] != "114514" {
		o.Reject("no privilege")
		return
	}
	o.Resolve(true)
}

func closeproxy(o *frpsp.CloseProxyObject) {
	req, _ := json.Marshal(o.Req)
	fmt.Println("CloseProxy", string(req))
	o.Resolve(true)
}

func ping(o *frpsp.PingObject) {
	req, _ := json.Marshal(o.Req)
	fmt.Println("Ping", string(req))
	o.Resolve(true)
}

func workconn(o *frpsp.NewWorkConnObject) {
	req, _ := json.Marshal(o.Req)
	fmt.Println("NewWorkConn", string(req))
	o.Resolve(true)
}

func userconn(o *frpsp.NewUserConnObject) {
	req, _ := json.Marshal(o.Req)
	if o.Req.ProxyType == frpsp.TCP {
		fmt.Println("NewUserConn (TCP)")
	} else {
		fmt.Println("NewUserConn", string(req))
	}
	o.Resolve(true)
}

func main() {
	frpsp.HandleLoginFunc("/login", login)
	frpsp.HandleNewProxyFunc("/newproxy", newproxy)
	frpsp.HandleCloseProxyFunc("/closeproxy", closeproxy)
	frpsp.HandlePingFunc("/ping", ping)
	frpsp.HandleNewWorkConnFunc("/workconn", workconn)
	frpsp.HandleNewUserConnFunc("/userconn", userconn)
	frpsp.ListenAndServe(":8000", nil)
}
