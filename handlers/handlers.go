package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prueba/db"
	"prueba/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitHandler(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	SendData(rw, users)
}

func getUserByID(r *http.Request) (models.User, *gorm.DB) {
	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	if err := db.Database.First(&user, userId); err.Error != nil {

		return user, err
	} else {
		return user, nil

	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByID(r); err != nil {
		fmt.Println("Error", err)
		SendNotfound(rw)
	} else {
		SendData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		SendUnprocessableEntity(rw)
	} else {
		db.Database.Save(&user)
		SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//obtener id
	var userId int64

	if user_ant, err := getUserByID(r); err != nil {
		SendNotfound(rw)
	} else {
		userId = user_ant.Id
		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			SendUnprocessableEntity(rw)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			SendData(rw, user)
		}

	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByID(r); err != nil {
		SendNotfound(rw)
	} else {
		db.Database.Delete(&user)
		SendData(rw, user)
	}

}
