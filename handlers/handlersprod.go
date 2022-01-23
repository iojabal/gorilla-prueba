package handlers

import (
	"encoding/json"
	"net/http"
	"prueba/models"
)

func CreateProducts(rw http.ResponseWriter, r *http.Request) {
	product := models.Article{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		product.Save()
		models.SendData(rw, product)
	}
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	if product, err := models.ListArticle(); err != nil {
		models.SendNotfound(rw)
	} else {
		models.SendData(rw, product)
	}
}
