package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	mux.HandleFunc("GET /",rootHandler)
	mux.HandleFunc("GET /health", healthChecker)
	mux.HandleFunc("POST /create-user", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	// mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)



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

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error decoding JSON", err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
 
	json.NewEncoder(w).Encode(newUser)

}

func getUserHandler (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// formatUsers, _ := json.Marshal(users)
	// w.Write(formatUsers)

	json.NewEncoder(w).Encode(users)

}

func getSingleUserHandler (w http.ResponseWriter, r *http.Request){
	idParam := r.PathValue("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid ID parameter")
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		    json.NewEncoder(w).Encode(user)
		 
		}
			
	}				
}

func updateUserHandler (w http.ResponseWriter, r *http.Request){
	idParam := r.PathValue("id")
	
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid ID parameter")
		return
	}
	updatedUser := User{}

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Req Body", err)
		return
	}

	
	for idx, user := range users {
		if user.ID == id {
		updatedUser.ID = user.ID
		users[idx] = updatedUser
					

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
		return
		 
		}
			
	}				
}