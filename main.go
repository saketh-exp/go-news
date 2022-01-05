package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")
	page := params.Get("page")

	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is : ", searchQuery)
	fmt.Println("Page is: ", page)
}
func main() {
	fmt.Println("Look into port 3000!")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env File.")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.HandleFunc("/search", searchHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
