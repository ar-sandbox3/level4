package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ar-sandbox3/level4/sum"
)

func main() {
	http.ListenAndServe(":8080", handler())
}

func handler() http.Handler {
	s := http.NewServeMux()
	s.HandleFunc("/sum", sumHandler)
	return s
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("a")
	b := r.FormValue("b")
	va, _ := strconv.Atoi(a)
	vb, _ := strconv.Atoi(b)
	s := sum.Ints(va, vb)
	w.Write([]byte(fmt.Sprintf("%d", s)))
}
