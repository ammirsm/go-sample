package main

import (
	"fmt"
	"net/http"
)


type JsonResponse struct {
	Code	int
	Description	string
}


func mainPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world, this is wealth ethical API main page :)")
}

func main() {
	fmt.Println("Server is up")
	initialMigration()
	handleRequests()
}