package models

import (
	"time"
	"database/sql"
	"github.com/go-xorm/xorm"
)

type Todo struct {
	ID int `json:"id" xorm:"'id'"`
	Name string `json:"name" xorm:"'name'"`
	Completed bool `json:"completed" xorm:"'completed'"`
	Due time.Time `json:"due" xorm:"'due'"`
	Created_at time.Time `json:"created_at" xorm:"'created_at'"`
}

type TodoRepository struct {}

type Todos []Todo

var engine *xorm.Engine

func NewTodoRepository() TodoRepository {
	return TodoRepository{}
}

// Get todo by using ID
func (m TodoRepository) GetTodoByID(id int) *Todo {
	var todo  = Todo{ID: id}
	has, _ := engine.Get(&todo)
	if has {
		return &todo
	}
	return nil
}

// Get all todos
func (m TodoRepository) GetAllTodos() Todos {
	var todo = Todo{}
	var todos = Todos{}
	rows, err := engine.Rows(todo)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&todo)
		todos = append(todos, todo)
	}
	return todos
}

//DBにtodoを追加する
func (t *Todo) Insert(db *sql.DB) (sql.Result, error) {
	stmt, err := db.Prepare(`INSERT INTO todos(name, completed, due, created_at) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Name, t.Completed, t.Due, t.Created_at)
}

func (m TodoRepository) Create(name string, completed bool, due time.Time, created_at time.Time) {
	var todo = Todo{Name: name, Completed: completed, Due: due, Created_at: created_at}
	engine.Insert(&todo)
}

// 全てのtodowを取得する
//func GetAllTodos(db *sql.DB) ([]Todo, error) {
//	results := make([]Todo, 0, 16)
//	var err error
//	rows, err := db.Query(`SELECT * FROM todos`)
//	if err != nil {
//		panic(err)
//	}
//	for rows.Next() {
//		var t Todo
//		if err = rows.Scan(
//			&t.Name,
//			&t.Completed,
//			&t.Due,
//			&t.Created_at,
//		); err != nil {
//			return nil, err
//		}
//		results = append(results, t)
//	}
//	return results, nil
//}
