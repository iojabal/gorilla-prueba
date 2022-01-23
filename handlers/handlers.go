package handlers

import (
	"encoding/json"
	"net/http"
	"prueba/models"
	"strconv"

	"github.com/gorilla/mux"
)

func InitHandler(rw http.ResponseWriter, r *http.Request) {

	if users, err := models.ListUsers(); err != nil {
		models.SendNotfound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotfound(rw)
	} else {
		models.SendData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//obtener id
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotfound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotfound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}

}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
