package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

type errorResp struct {
	Statuscode int
	Message    string
}

var (
	users = map[string]User{}
)

func main() {
	http.HandleFunc("/createuser", adduser)
	http.HandleFunc("/returnuser", getusers)
	fmt.Println("Users are: ", users)
	fmt.Println("Server started")
	log.Fatalf("Server is not getting active, err:%v\n", http.ListenAndServe(":8000", nil))
}

func adduser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		errr := errorResp{
			Statuscode: http.StatusBadRequest,
			Message:    "Focus on error, err= " + err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	users[user.Name] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	fmt.Println("Users are: ", users)
	return
}

func getusers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		errr := errorResp{
			Statuscode: http.StatusBadRequest,
			Message:    "Payload could not make it, err= " + err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	return
}
