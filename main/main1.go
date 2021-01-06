package main

import (
	"fmt"
	"net/http"
	"openapm-monitoring"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func init() {

}
func main() {
	openapm.Exec()
	serverMuxA := http.NewServeMux()
	//serverMuxA.HandleFunc("/hello", hello)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/world", homeLink)

	go func() {
		http.ListenAndServe("localhost:8081", serverMuxA)
	}()

	http.ListenAndServe("localhost:8082", serverMuxB)

	//log.Fatal(http.ListenAndServe(":8081", nil))

}
