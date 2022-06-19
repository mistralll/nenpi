package main

import (
	"net/http"
	"text/template"

	"github.com/mistralll/nenpi/refueling"
	"github.com/mistralll/nenpi/vihicle"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := vihicle.LoadVihicle(title)
	t, _ := template.ParseFiles("html/view.html")
	t.Execute(w, p)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[5:]
	t, _ := template.ParseFiles("html/add.html")
	p := map[string]string{"Title": title}
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	t, _ := template.ParseFiles("html/edit.html")
	p := map[string]string{"Title": title}
	t.Execute(w, p)
}

func saveHanler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	refuel := refueling.HttpReqToRefuel(r)
	refuel.SaveRefuel(title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func saveiconHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[10:]

	vihicle.SaveIcon(r, title)

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/edit/", editHandler)

	http.HandleFunc("/save/", saveHanler)
	http.HandleFunc("/saveicon/", saveiconHandler)

	http.ListenAndServe(":8080", nil)
}
