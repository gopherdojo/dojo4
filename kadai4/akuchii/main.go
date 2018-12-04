package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gopherdojo/dojo4/kadai4/akuchii/fortune"
)

// servert port
var port int

func init() {
	flag.IntVar(&port, "p", 8080, "server port")
}

func main() {
	flag.Parse()
	f := fortune.NewFortune(fortune.DefaultClock{})
	http.HandleFunc("/", f.Handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
