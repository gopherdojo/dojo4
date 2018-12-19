package main

import (
	"net/http"
	"time"

	"github.com/gopherdojo/dojo4/kadai4/uobikiemukot/fortune"
)

func main() {
	clock := fortune.ClockFunc(func() time.Time {
		return time.Now()
	})
	f := fortune.Fortune{Clock: clock}

	http.HandleFunc("/", f.Handler)
	http.ListenAndServe(":8080", nil)
}
