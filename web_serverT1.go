package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/User/", UserHandler)
	handler.HandleFunc("/Users/", UsersHandler)

	s := http.Server{
		Addr:           ":8081",
		Handler:        handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		IdleTimeout:    100 * time.Second,
		MaxHeaderBytes: 1 << 40,
	}
	log.Fatal(s.ListenAndServe())
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"user Name"`
	Lastname string `json:"last Name"`
}
type UserStore struct {
	users []User
}
type Status struct {
	Message interface{}
	Error   string
}

var userStore = UserStore{
	[]User{
		{"1", "Dima", "Neverov"},
		{"2", "Sasha", "Zaporozhcev"},
		{"3", "Ivan", "Ivanov"},
	}}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		handleGetUser(w, r)
	} else if r.Method == http.MethodPost {
		handleAddUser(w, r)
	} else if r.Method == http.MethodPost {
		handleAddUser(w, r)
	} else if r.Method == http.MethodDelete {
		handleDeleteUser(w, r)
	} else if r.Method == http.MethodPut {
		handleUpdateUser(w, r)
	}
}
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/User/", "", 1)
	decoder := json.NewDecoder(r.Body)
	var user User
	var status Status
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status.Error = err.Error()
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	user.Id = id
	err = userStore.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status.Error = err.Error()
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	status.Message = user
	statusJson, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)

	w.Write(statusJson)
}
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/User/", "", 1)

	var status Status

	err := userStore.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status.Error = err.Error()
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	UsersHandler(w, r)
}

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var user User
	var status Status
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status.Error = err.Error()
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	err = userStore.AddUsers(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		status.Error = err.Error()
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	UsersHandler(w, r)

}
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		handleGetUser(w, r)
	}
	status := Status{
		Message: userStore.GetUsers(),
	}
	usersJson, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/User/", "", 1)
	user := userStore.FindUserByID(id)
	var status Status
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		status.Error = fmt.Sprintf("")
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
		return
	}
	status.Message = user
	statusJson, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)

	w.Write(statusJson)
}

func (s UserStore) FindUserByID(id string) *User {
	for _, user := range s.users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}
func (s UserStore) GetUsers() []User {
	return s.users
}
func (s *UserStore) AddUsers(user User) error {
	for _, us := range s.users {
		if us.Id == user.Id {

			return errors.New(fmt.Sprintf("User with id %s not found", user.Id))
		}
	}
	s.users = append(s.users, user)
	return nil
}
func (s *UserStore) UpdateUser(user User) error {
	for i, is := range s.users {
		if is.Id == user.Id {
			s.users[i] = user
			return nil
		}
	}
	return errors.New(fmt.Sprintf("User with id %s not found", user.Id))
}
func (s *UserStore) DeleteUser(id string) error {
	for i, is := range s.users {
		if is.Id == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}

	}
	return errors.New(fmt.Sprintf("User with id %s not found", id))
}
