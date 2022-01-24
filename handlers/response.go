package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

func CreatDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

func (resp *Response) Send() {
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintf(resp.respWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreatDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) Notfound() {
	resp.Status = http.StatusNotFound
	resp.Message = "resource not found"
}

func SendNotfound(rw http.ResponseWriter) {
	response := CreatDefaultResponse(rw)
	response.Notfound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprosesableEntity not found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreatDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
