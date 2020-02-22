package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainPage)

	port := ":9999"
	println("Conected to port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("FATAL ERROR")
	}

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
