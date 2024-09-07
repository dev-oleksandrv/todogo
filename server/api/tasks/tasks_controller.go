package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/internal/response"
	"github.com/gorilla/mux"
)

type TaskController struct {
	taskService *TaskService
}

func NewTaskController(taskService *TaskService) *TaskController {
	return &TaskController{taskService}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Println(err)
		response.JSON(
			w, 
			http.StatusBadRequest, 
			map[string]string{"error": "Invalid Request payload"},
		)
		return
	}
	task, err := c.taskService.CreateTask(task)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot create a Task"},
		)
		return
	}
	response.JSON(w, http.StatusCreated, task)
}

func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.taskService.GetAllTasks()
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot retreive a task"},
		)
		return
	}
	response.JSON(w, http.StatusOK, tasks)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Invalid id of Task"},
		)
		return
	}
	if err := c.taskService.DeleteTask(uint(id)); err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot delete a task"},
		)
		return
	}
	response.JSON(w, http.StatusNoContent, map[string]string{})
}