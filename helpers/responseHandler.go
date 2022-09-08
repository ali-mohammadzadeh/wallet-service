package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MetaStruct struct {
	Time        time.Time `json:"time"`
	PackageName string    `json:"packageName"`
}
type ResponseStruct struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}
type ErrorStruct struct {
	Message string `json:"message"`
}

var errorMessage ErrorStruct

func SuccessResponse(writer http.ResponseWriter, instance interface{}, statusCode int) {
	var response ResponseStruct
	var meta MetaStruct
	meta.Time = time.Now()
	meta.PackageName = "serviceWallet"
	response.Data = instance
	response.Meta = meta
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(response)
}

func ErrorResponse(writer http.ResponseWriter, message string, statusCode int) {
	errorMessage.Message = message
	var response ResponseStruct
	var meta MetaStruct
	meta.Time = time.Now()
	meta.PackageName = "serviceWallet"
	response.Data = errorMessage
	response.Meta = meta
	fmt.Println(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(response)
}
