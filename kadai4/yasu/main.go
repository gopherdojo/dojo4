package main

import (
	"dojo4/kadai4/yasu/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/rottely", controllers.DevineFortune)

	http.ListenAndServe(":8080", nil)
}
