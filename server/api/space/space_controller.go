package space

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-oleksandrv/api/auth"
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/internal/response"
	"github.com/gorilla/mux"
)

type SpaceController struct {
	spaceService *SpaceService
}

func NewSpaceController(spaceService *SpaceService) *SpaceController {
	return &SpaceController{spaceService}
}

func (c *SpaceController) CreateSpace(w http.ResponseWriter, r *http.Request) {
	var space db.Space
	if err := json.NewDecoder(r.Body).Decode(&space); err != nil {
		fmt.Println(err)
		response.JSON(
			w, 
			http.StatusBadRequest, 
			map[string]string{"error": "Invalid Request payload"},
		)
		return
	}
	userID := r.Context().Value(auth.GetUserIDContextKey()).(int)
	space, err := c.spaceService.CreateSpace(userID, space)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot create a Space"},
		)
		return
	}
	response.JSON(w, http.StatusCreated, space)
}

func (c *SpaceController) GetSpaces(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.GetUserIDContextKey()).(int)
	spaces, err := c.spaceService.GetAllSpacesByUserID(userID)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot retreive a space"},
		)
		return
	}
	response.JSON(w, http.StatusOK, spaces)
}

func (c *SpaceController) DeleteSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Invalid id of Space"},
		)
		return
	}
	userID := r.Context().Value(auth.GetUserIDContextKey()).(int)
	if err := c.spaceService.DeleteSpace(userID, int(id)); err != nil {
		response.JSON(
			w, 
			http.StatusInternalServerError, 
			map[string]string{"error": "Cannot delete a space"},
		)
		return
	}
	response.JSON(w, http.StatusNoContent, map[string]string{})
}