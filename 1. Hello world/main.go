package main

import (
	"fmt"
	"net/http"
)

func Halo_web(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!, from web")
}

func main() {
	http.HandleFunc("/", Halo_web)

	fmt.Println("Server is on, please check http://localhost:7000")

	if err := http.ListenAndServe(":7000", nil); err != nil {
		fmt.Println(err)
	}
}
