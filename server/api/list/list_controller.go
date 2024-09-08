package list

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/internal/response"
	"github.com/gorilla/mux"
)

type ListController struct {
	listService *ListService
}

func NewListController(listService *ListService) *ListController {
	return &ListController{listService}
}

func (c *ListController) CreateList(w http.ResponseWriter, r *http.Request) {
	var list db.List
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		fmt.Println(err)
		response.JSON(
			w, 
			http.StatusBadRequest, 
			map[string]string{"error": "Invalid Request payload"},
		)
		return
	}
	list, err := c.listService.CreateList(list)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot create a List"},
		)
		return
	}
	response.JSON(w, http.StatusCreated, list)
}

func (c *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	lists, err := c.listService.GetAllLists()
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot retreive a list"},
		)
		return
	}
	response.JSON(w, http.StatusOK, lists)
}

func (c *ListController) DeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Invalid id of List"},
		)
		return
	}
	if err := c.listService.DeleteList(uint(id)); err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot delete a list"},
		)
		return
	}
	response.JSON(w, http.StatusNoContent, map[string]string{})
}