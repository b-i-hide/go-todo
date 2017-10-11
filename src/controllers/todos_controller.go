package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"localhost/go-todo/src/models"
	"encoding/json"
)

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := models.Todos{
		models.Todo{Name: "Write presentation"},
		models.Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
