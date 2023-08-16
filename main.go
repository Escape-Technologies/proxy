package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

func main() {
	start("0.0.0.0", "8080", "user", "test")
}

func start(addr string, port string, user string, password string) {
	total := fmt.Sprintf("%s:%s", addr, port)
	proxy := goproxy.NewProxyHttpServer()

	auth.ProxyBasic(proxy, "realm", func(user, pwd string) bool {
		return user == "user" && password == pwd
	})

	log.Printf("Listening on %s", total)
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(total, proxy))
}
