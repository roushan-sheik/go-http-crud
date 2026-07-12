package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", rootHandler)

	fmt.Println("Server is running on port 5000")
	err:=http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("Server Error",err)
	 
	}
}


func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Welcome to go server")
}