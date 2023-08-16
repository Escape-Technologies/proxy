package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	username := "admin"
	password := "admin"

	start("0.0.0.0", "8080", hash(fmt.Sprintf("%s:%s", username, password)))
}

func hash(data string) []byte {
	bytes := []byte(data)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash
}

func compare(hash []byte, data string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(data))
	if err != nil {
		return false
	}
	return true
}

func start(addr string, port string, expected_hash []byte) {
	total := fmt.Sprintf("%s:%s", addr, port)
	proxy := goproxy.NewProxyHttpServer()

	auth.ProxyBasic(proxy, "realm", func(user, pass string) bool {
		return compare(expected_hash, fmt.Sprintf("%s:%s", user, pass))
	})

	log.Printf("Listening on %s", total)
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(total, proxy))
}
