package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type ViewDate struct {
	title string
	dt    string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := time.Now()
		tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
		tmpl.Execute(w, data)

	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		data := ViewDate{
			title: "WorldCup",
		}

		tmpl, _ := template.ParseFiles("templates/index.html")
		// http.ServeFile(w, r, "templates/index.html")
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		cardnumber := r.FormValue("number")
		cardHolder := r.FormValue("holder")
		expDate := r.FormValue("date")
		ccv := r.FormValue("ccv")
		fmt.Fprintf(w, "CardNumber: %s; cardHolder: %s; expDate: %s; ccv: %s;", cardnumber, cardHolder, expDate, ccv)

	})
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
