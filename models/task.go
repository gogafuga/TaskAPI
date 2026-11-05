package models

type Tasks struct {
	Id        int    `json:"id"`        //уникальный идентификатор
	Title     string `json:"title"`     //название задачи
	Completed bool   `json:"completed"` //выполнена ли задача?
}
