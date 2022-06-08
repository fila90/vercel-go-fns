package handler

import (
	"fmt"
	"net/http"
)

func Addone(x int) int {
	return x + 1
}

func Date(w http.ResponseWriter, r *http.Request) {
	num := Addone(1)
	fmt.Fprintf(w, "%d", num)
}
