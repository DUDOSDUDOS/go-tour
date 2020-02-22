package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)

	port := ":9999"
	println("SRABOTALO ZBS:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("VI OBOSRALIS`")
	}

}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	//user := User{"Кирилл", "Яйцеглист"}
	//js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func usersPage(w http.ResponseWriter, r *http.Request) {
	users := []User{User{"Кирилл", "Яйцеглист"}, User{"Антон", "Гнида"}}
	//js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
