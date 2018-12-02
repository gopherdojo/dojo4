package main

import (
	"flag"
	"fmt"
	"net/http"

	"tomoyukikobayashi/handler"
)

const defaultPort = "8080"

var servePort = flag.String("port", defaultPort, "service port")

func main() {
	flag.Parse()

	fmt.Printf("serve port : %v", *servePort)

	http.HandleFunc("/", handler.Fortune)
	http.ListenAndServe(":"+*servePort, nil)
}
