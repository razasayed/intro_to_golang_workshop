package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//jsonHandler returns http respone in JSON format.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{Id: 1,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
		Phone: "000099999"}
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world!")
	})

	http.HandleFunc("/json", jsonHandler)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
