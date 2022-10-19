package main

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// Save page locally
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// Load page from local directory
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

// Render html template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	temp, _ := template.ParseFiles(tmpl + ".html")
	err := temp.Execute(w, p)
	if err != nil {
		w.Write([]byte("Can't render html"))
	}
}

// handles /view request
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, _ := loadPage(title)
	renderTemplate(w, "view", page)
}

// handles /edit request
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: "Edit"}
	}
	renderTemplate(w, "edit", page)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}
