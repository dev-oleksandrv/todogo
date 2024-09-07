package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dev-oleksandrv/internal/response"
)

type TaskController struct {
}

type Task interface{}

var localTasks []Task

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Println(err)
		response.JSON(
			w, 
			http.StatusBadRequest, 
			map[string]string{"error": "Invalid Request payload"},
		)
		return
	}
	localTasks = append(localTasks, task)
	response.JSON(w, http.StatusCreated, task)
}

func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, localTasks)
}