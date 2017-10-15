package go_todo

import (
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db *sql.DB
	mux *http.ServeMux
}

// new server object
func New() *Server {
	return &Server{}
}

// close database connection
func (s *Server) Close() error {
	return s.db.Close()
}

func (s *Server) Init() {
	db, err := sql.Open("mysql", "root@/go_todo_development")
	if err != nil {
		panic(err)
	}
	s.db = db
	//s.Router()
}

func (s *Server) Router() {
	//router := httprouter.New()

	//todo := &controllers.Todo{DB: s.db}

	//router.GET("/", todo.Root)
	//router.GET("/todos/:todoId", todo.TodoShow)
	//router.POST("/todos", todo.Create)
	//router.POST("/todos/new", todo.CreateTodo)

	//http.ListenAndServe(":8080", router)
}
