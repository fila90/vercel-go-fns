package handler

import (
	"fmt"
	"net/http"
)

func Notdate(w http.ResponseWriter, r *http.Request) {
	num := Addone(2)
	fmt.Fprintf(w, "%d", num)
}
