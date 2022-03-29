package server

import (
	"fmt"
	"net/http"
)

//viet custom handler
type handle struct {
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home Page"))

	}else if r.URL.Path == "/about"{
		w.Write([]byte("About page"))
	}else{
	w.Write([]byte("Custom handler"))
}
}

func DemoCustomServer() {
	handle := new(handle)
	fmt.Println(http.ListenAndServe(":3000", handle))
}
