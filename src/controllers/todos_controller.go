package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
