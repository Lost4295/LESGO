package web

import (
	"html/template"
	"log"
	"net/http"
	// "os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

// func (p *Page) save() error {
// 	filename := "web/data"+ p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }


func homeHandler(w http.ResponseWriter, r *http.Request, title string) {
	// p, err := loadPage(title)
	// if err != nil {
	// 	http.Redirect(w, r, "/notfound/", http.StatusNotFound)
	// 	return
	// }
	renderTemplate(w, "home", nil)
}

// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return
// 	}
// 	renderTemplate(w, "view", p)
// }

// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	renderTemplate(w, "edit", p)
// }

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	body := r.FormValue("body")
// 	p := &Page{Title: title, Body: []byte(body)}
// 	err := p.save()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

var templates = template.Must(template.ParseFiles(
	// "web/edit.html", "web/view.html",
	"web/home.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/home/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[0])
	}
}

func Main() {
	http.HandleFunc("/home/", makeHandler(homeHandler))

	go log.Fatal(http.ListenAndServe(":8000", nil))
}
