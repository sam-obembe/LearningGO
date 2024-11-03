package main

import (
	"errors"
	"fmt"
	"gowiki/models"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var title = r.PathValue("title")
	var page, err = models.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "views/view.html", page)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	var title = r.PathValue("title")
	var page, err = models.LoadPage(title)
	if err != nil {
		page = &models.Page{Title: title}
	}

	renderTemplate(w, "views/edit.html", page)
}

func renderTemplate(w http.ResponseWriter, tem string, p *models.Page) {
	//var t, err = template.ParseFiles(tem)
	//w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, tem, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	var title = r.PathValue("title")
	var body = r.FormValue("body")
	var page = &models.Page{Title: title, Body: []byte(body)}
	err := page.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/{title}", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// var p1 = &models.Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.Save()

	// var p2, _ = models.LoadPage("TestPage")
	// fmt.Println(string(p2.Body))
}
