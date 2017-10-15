package controllers

import (
	"localhost/go-todo/src/models"
	"time"
)

type Todo struct {
}

func NewTodo() Todo {
	return Todo{}
}

func (c Todo) Create(name string, completed bool, due time.Time, created_at time.Time)  {
	repo := models.NewTodoRepository()
	repo.Create(name, completed, due, created_at)
}

func (c Todo) Get(n int) interface{} {
	repo := models.NewTodoRepository()
	todo := repo.GetTodoByID(n)
	return todo
}

func (c Todo) GetAll() interface{} {
	repo := models.NewTodoRepository()
	todos := repo.GetAllTodos()
	return todos
}

/*
func (t *Todo) Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println(r)
	todos, err := models.GetAllTodos(t.DB)
	if err = json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func (t *Todo) CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := t.Model.Insert(t.DB)
	if err != nil {
		panic(err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, fmt.Sprintf("/"), 301)
}
*/

//func (t *Todo) Save(w http.ResponseWriter, r *http.Request) error {
//	var todo models.Todo
//}

/*
func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Todo Index!")
}

func (t *Todo) TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}

func (t *Todo) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db := t.DB
	fmt.Println(params)
	//if params.ByName("name") == "" {
	//	http.Error(w, "name can't be blank", 403)
	//}
	//
	//if params.ByName("due") == "" {
	//	http.Error(w, "due can't be blank", 403)
	//}

	fmt.Println("params[:dur]", params.ByName("due"))
	stmt, err := db.Prepare("INSERT todos SET name=?, due=?")
	if err != nil {
		panic(err)
	}

	fmt.Println("params[:dur]", params.ByName("due"))
	res, err := stmt.Exec(params.ByName("name"), params.ByName("due"))
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
 */

