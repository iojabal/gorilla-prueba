package handlers

import (
	"encoding/json"
	"net/http"
	"prueba/db"
	"prueba/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateProducts(rw http.ResponseWriter, r *http.Request) {
	product := models.Article{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		SendUnprocessableEntity(rw)
	} else {
		db.Database.Save(&product)
		SendData(rw, product)
	}
}

func GetProducts(rw http.ResponseWriter, r *http.Request) {
	product := models.Articles{}
	db.Database.Preload("User").Find(&product)
	//db.Database.Find(&product)
	SendData(rw, product)
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	if art, err := getArtiByID(r); err != nil {
		SendNotfound(rw)
	} else {
		SendData(rw, art)
	}
}

func getArtiByID(r *http.Request) (models.Article, *gorm.DB) {
	//Obtener ID
	vars := mux.Vars(r)
	artId, _ := strconv.Atoi(vars["id"])

	art := models.Article{}
	if err := db.Database.Preload("User").First(&art, artId); err.Error != nil {

		return art, err
	} else {
		return art, nil

	}
}

func UpdateArti(rw http.ResponseWriter, r *http.Request) {
	var artId int64
	if art_ant, err := getArtiByID(r); err != nil {
		SendNotfound(rw)
	} else {
		artId = art_ant.Id
		product := models.Article{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&product); err != nil {
			SendUnprocessableEntity(rw)
		} else {
			product.Id = artId
			db.Database.Preload("User").Find(&product)
			SendData(rw, product)
		}
	}

}
