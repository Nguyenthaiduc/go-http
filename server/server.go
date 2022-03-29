package server

import (
	"fmt"
	"net/http"
)
//controlller home
func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}
//controller about
func aboutPage(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("About Page"))
}
//server using http
func DefaultServer(){
	http.HandleFunc("/",homeHandle)
	fmt.Println("Server listening on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000",nil))
}
//server using mux

func DefaultServerMux(){
	mux := http.NewServeMux()
	mux.HandleFunc("/about",aboutPage)

	fmt.Println("Server listening on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000",mux))

}