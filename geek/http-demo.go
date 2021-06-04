package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request)  {
	// r.ParseForm()
	fmt.Println("Form:", r.Form)
	fmt.Println("Url:", r.URL.Path)
	fmt.Fprintf(w, "It works, Path:%s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", testHandle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("", err)
	}
}



