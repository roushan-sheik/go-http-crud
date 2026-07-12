package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{
		ID:    1,
		Name:  "Roushan",
		Email: "roushansheik@gmail.com",
	},
	{
		ID:    2,
		Name:  "Rohim",
		Email: "rohim@gmail.com",
	},
	{
		ID:    3,
		Name:  "Karim",
		Email: "karim@gmail.com",
	},
}  

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/",rootHandler)
	mux.HandleFunc("/health", healthChecker)
	mux.HandleFunc("POST /create-user", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)


	fmt.Println("Server is running on port 5000")
	err:=http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server Error",err)
	}
}


func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Welcome to go server")
}

func healthChecker (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Server health is good")
}

func createUserHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "User Created Successfully")
}

func getUserHandler (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// formatUsers, _ := json.Marshal(users)
	// w.Write(formatUsers)

	encoder := json.NewEncoder(w)
	encoder.Encode(users)

}