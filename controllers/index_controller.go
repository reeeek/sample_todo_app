package controllers

import (
	"encoding/json"
	"go_test_sample/sample_gorilla/infrastructure/repository/task"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	ContentTypeJson = "application/json"
)

type HTTPResponse struct {
	contentType string
	statusCode  int
	data        interface{}
}

func NewHTTPResponse(code int, data interface{}) HTTPResponse {
	return HTTPResponse{
		contentType: ContentTypeJson,
		statusCode:  code,
		data:        data,
	}
}

func WriteJSON(r HTTPResponse, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", r.contentType)
	w.WriteHeader(r.statusCode)
	return json.NewEncoder(w).Encode(r.data)
}

type IndexController struct {
}

func (c *IndexController) Get(w http.ResponseWriter, r *http.Request) {
	repo := task.NewTaskRepository()
	tasks, err := repo.FindAll()
	if err != nil {
		// error
	}
	response := NewHTTPResponse(http.StatusOK, tasks)
	resErr := WriteJSON(response, w)
	if resErr != nil {
		// 500
		response := HTTPResponse{
			contentType: "application/json",
			statusCode:  500,
			data:        resErr,
		}
		WriteJSON(response, w)
	}
}

func (c *IndexController) Post(w http.ResponseWriter, r *http.Request) {
	repo := task.NewTaskRepository()
	ent := &task.Task{
		Status:    1,
		Title:     r.FormValue("title"),
		Body:      r.FormValue("data"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := repo.Store(ent)
	if err != nil {
		// error
	}

	tasks, err := repo.FindAll()
	if err != nil {
		// error
	}
	response := HTTPResponse{
		contentType: "application/json",
		statusCode:  201,
		data:        tasks,
	}
	err = WriteJSON(response, w)
	if err != nil {
		// 500
		response := HTTPResponse{
			contentType: "application/json",
			statusCode:  500,
			data:        err,
		}
		WriteJSON(response, w)
	}
}

func (c *IndexController) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["task_id"]
	repo := task.NewTaskRepository()
	i, _ := strconv.Atoi(taskID)
	ent := &task.Task{}
	ent.Title = vars["title"]
	ent.Body = vars["data"]
	err := repo.Update(i, ent)
	if err != nil {
		// error
	}
}
