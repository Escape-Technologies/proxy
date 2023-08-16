package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"golang.org/x/crypto/bcrypt"
)

var UUID = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var PORT = regexp.MustCompile(`^[0-9]{1,5}$`)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	if ! PORT.MatchString(port) {
		log.Printf("PORT must be a number between 0 and 65535")
		os.Exit(1)
	}

	username := os.Getenv("ESCAPE_ORGANIZATION_ID")
	if ! UUID.MatchString(username) {
		log.Printf("ESCAPE_ORGANIZATION_ID must be a UUID in lowercase")
		log.Printf("To get your organization id, go to https://app.escape.tech/organization/settings/")
		os.Exit(1)
	}
	password := os.Getenv("ESCAPE_API_KEY")
	if ! UUID.MatchString(password) {
		log.Printf("ESCAPE_API_KEY must be a UUID in lowercase")
		log.Printf("To get your API key, go to https://app.escape.tech/user/profile/")
		os.Exit(1)
	}

	start("0.0.0.0", port, hash(username + password))
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
	proxy := goproxy.NewProxyHttpServer()

	auth.ProxyBasic(proxy, "realm", func(user, pass string) bool {
		return compare(expected_hash, user + pass)
	})

	total := fmt.Sprintf("%s:%s", addr, port)
	log.Printf("Listening on %s", total)
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(total, proxy))
}
