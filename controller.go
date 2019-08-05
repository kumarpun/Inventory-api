package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var devices = []Device{}

type Response struct {
	Method  string `json:"method"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(method, message string, status int) Response {

	return Response{Method: method, Message: message, Status: status}

}

func HttpInfo(r *http.Request) {

	fmt.Printf("%s\t %s\t %s%s\r\n", r.Method, r.Proto, r.Host, r.URL)

}
func setJsonHeader(w http.ResponseWriter) {

	w.Header().Set("Content-type", "application/json")

}

func getDevices(w http.ResponseWriter, r *http.Request) {
	setJsonHeader(w)

	HttpInfo(r)
	var devices []Device
	if err := db.Find(&devices).Error; err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(devices)

	}
}

func getDevice(w http.ResponseWriter, r *http.Request) {

	setJsonHeader(w)

	HttpInfo(r)

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var device Device
	if err := db.Where("id = ?", id).First(&device).Error; err != nil {

		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(device)
	return
}

func postDevice(w http.ResponseWriter, r *http.Request) {

	setJsonHeader(w)

	HttpInfo(r)

	body, _ := ioutil.ReadAll(r.Body)

	var device Device

	err := json.Unmarshal(body, &device)

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(NewResponse(r.Method, "failed", 400))
		return

	}

	devices = append(devices, device)
	device.Status = "pending"
	db.Create(&device)

	json.NewEncoder(w).Encode(NewResponse(r.Method, "success", 201))

}

func putDevice(w http.ResponseWriter, r *http.Request) {

	setJsonHeader(w)
	HttpInfo(r)

	body, _ := ioutil.ReadAll(r.Body)

	var device Device

	err := json.Unmarshal(body, &device)

	if err != nil {

		json.NewEncoder(w).Encode(NewResponse(r.Method, "failed", 400))
		return

	}

	devices = append(devices, device)
	db.Save(&device)
	json.NewEncoder(w).Encode(NewResponse(r.Method, "success", 201))
	return
}

func deleteDevice(w http.ResponseWriter, r *http.Request) {

	setJsonHeader(w)

	HttpInfo(r)

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	var device Device
	d := db.Where("id = ?", id).Delete(&device)
	fmt.Println(d)

	for index, _ := range devices {

		if devices[index].Id == id {

			devices = append(devices[:index], devices[index+1:]...)

			return

		}
		json.NewEncoder(w).Encode(NewResponse(r.Method, "sucess", 200))
		return
	}

}

// Lists

func approveDevice(w http.ResponseWriter, r *http.Request) {

	setJsonHeader(w)

	HttpInfo(r)

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var device Device
	if err := db.Where("id = ?", id).First(&device).Error; err != nil {

		fmt.Println(err)
	}
	deviceDetail := db.Where("id = ?", id).Save(&device)
	// deviceDetailencoded := json.NewEncoder(w).Encode(deviceDetail)

	fmt.Println(deviceDetail)

	json.NewEncoder(w).Encode(device)
	device.Status = "approved"
	db.Save(&device)
	fmt.Println(device)
	return
}
