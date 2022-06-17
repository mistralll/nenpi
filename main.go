package main

import (
	"bufio"
	"net/http"
	"os"
	"text/template"

	"github.com/mistralll/goSrv/refueling"
	"github.com/mistralll/goSrv/vihicle"
)

type Page struct {
	Title string
	Rows  []refueling.Refueling
}

func loadPageView(title string) (*Page, error) {
	filename := title + ".csv"
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var rows []refueling.Refueling
	for scanner.Scan() {
		line := scanner.Text()
		row := refueling.StrToRefuel(line)
		rows = append(rows, *row)
	}

	p := &Page{Title: title, Rows: rows}

	return p, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := vihicle.LoadVihicle(title)
	t, _ := template.ParseFiles("html/view.html")
	t.Execute(w, p)
}

func saveHanler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	refuel := refueling.HttpReqToRefuel(r)
	refuel.SaveRefuel(title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[5:]
	t, _ := template.ParseFiles("html/add.html")
	p := map[string]string{"Title": title}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/save/", saveHanler)
	http.ListenAndServe(":8080", nil)
}
