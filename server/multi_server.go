package server

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world at %d", port)
	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return &server
}

func DemoMutipleServer() {
	wg := new(sync.WaitGroup)
	//add number server
	wg.Add(3)
	//go routing
	go func() {
		server := createServer(3000)
		fmt.Println(server.ListenAndServe())
		fmt.Println("1")
		wg.Done() // moi lan done la tru di 1
		
	}()
	//server 2
	go func() {
		server := createServer(3001)
		fmt.Println(server.ListenAndServe())
		fmt.Println(2)
		wg.Done()
	}()
	//server 3
	go func() {
		server := createServer(3002)
		fmt.Println(server.ListenAndServe())
		fmt.Println(3)
		wg.Done()
	}()
	wg.Wait()

}
// cach nay khong kha dung
func DemoMutipleServerDefault() {
	fmt.Println("start")
	server := createServer(3000)
	fmt.Println(server.ListenAndServe())
	fmt.Println("done")

	server1 := createServer(3001)
	fmt.Println(server1.ListenAndServe())

	server2 := createServer(3002)
	fmt.Println(server2.ListenAndServe())
}
