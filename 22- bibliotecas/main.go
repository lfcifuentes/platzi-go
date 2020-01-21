package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Iniciando servidor")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello	")
}
