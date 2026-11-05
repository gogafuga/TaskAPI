package main

import (
	"TaskAPI/models"
	"TaskAPI/postgre"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := models.Tasks{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	task.Title = strings.Trim(task.Title, " ")
	if task.Title == "" {
		http.Error(w, "not title", http.StatusBadRequest)
		return
	}

	id, err := postgre.CreateTask(ctx, conn, task.Title)
	if err != nil {
		http.Error(w, "OOOPs", http.StatusInternalServerError)
		return
	}
	fmt.Println("ID созданной задачи: ", id)
	task.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("Неверный метод ! Ожидалось Get")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	getTasks, err := postgre.GetTasks(r.Context(), conn)
	if err != nil {
		http.Error(w, "OOOPs", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(getTasks)

}

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "not id true", http.StatusBadRequest)
		return
	}
	postgre.DoneTask(ctx, conn, id)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "not id true", http.StatusBadRequest)
		return
	}

	if err := postgre.DeleteTask(ctx, conn, id); err != nil {
		http.Error(w, "OOOPs", http.StatusInternalServerError)
		return
	}
}
