package handlers

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Returns all users.
func AllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := pool.Dispatch(pool.UsersAll, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users)
}

// Returns user with given id.
// Path param: "id" - user id.
func GetUserById(w http.ResponseWriter, req *http.Request) {
	id := mux.Params(req).PathParams["id"]
	user, err := pool.Dispatch(pool.UserFindById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user)
}

// Returns all projects of given user
func GetAllProjectsFromUser(w http.ResponseWriter, req *http.Request) {
	id, _ := mux.Params(req).PathParams["id"]

	modelId, err := models.NewRequiredId(id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	projects, err := pool.Dispatch(pool.UserAllProjects, modelId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, projects)
}
