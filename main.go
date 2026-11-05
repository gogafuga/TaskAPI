package main

import (
	"TaskAPI/models"
	"TaskAPI/postgre"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5"
)

var mtx sync.RWMutex
var conn *pgx.Conn
var ctx = context.Background()
var tasks = make([]models.Tasks, 0)
var id int

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTaskHandler(w, r)
	case http.MethodPost:
		CreateTaskHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	var err error
	conn, err = postgre.ConnectionPSQl(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	fmt.Println("Подключение к базе данных прошло успешно!")

	if err := postgre.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("PATCH /tasks/{id}", DoneTaskHandler)
	http.HandleFunc("DELETE /tasks/{id}", DeleteTaskHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
