package main

import (
	"TaskAPI/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTasks(t *testing.T) {
	task1 := models.Tasks{
		Id:    1,
		Title: "Изучить GO",
	}

	task2 := models.Tasks{
		Id:    2,
		Title: "Изучить C++",
	}
	tasks = []models.Tasks{task1, task2}

	localTask := make([]models.Tasks, 0)
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	GetTaskHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Error("WAIT STATUS ", http.StatusOK, " BUT COMING ", rec.Code)
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &localTask); err != nil {
		t.Fatal(err)
	}
	if len(localTask) != 2 {
		t.Fatal("No true len")
	}
	if localTask[0].Title != "Изучить GO" || localTask[1].Title != "Изучить C++" {
		t.Error("No true data")
	}
}

func TestCreateTasks(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`
	{
	"title":"Изучить GO",
	"completed":false
	}
	`))

	trueID := id
	trueTasks := tasks
	tasks = []models.Tasks{}
	id = 0
	t.Cleanup(func() {
		id = trueID
		tasks = trueTasks
	})

	CreateTaskHandler(rec, req)
	if rec.Code != http.StatusCreated {
		t.Error("WAIT STATUS ", http.StatusCreated, " BUT COMING ", rec.Code)
	}
	localTask := models.Tasks{}

	if err := json.Unmarshal(rec.Body.Bytes(), &localTask); err != nil {
		t.Fatal(err)
	}

	if localTask.Title != "Изучить GO" || localTask.Id != 1 {
		t.Fatal("NO TRUE DATA")
	}
}
