package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/mistralll/goSrv/row"
)

type Page struct {
	Title string
	Rows  []row.Row
}

func loadPageView(title string) (*Page, error) {
	filename := title + ".csv"
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	var rows []row.Row
	for scanner.Scan() {
		line := scanner.Text()
		r := row.StringToRow(line)
		rows = append(rows, *r)
	}

	p := &Page{Title: title, Rows: rows}

	for _, v := range p.Rows {
		fmt.Println(v)
	}
	return p, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := loadPageView(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func saveHanler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	row := row.HttprequesToRow(r)
	row.SaveRow(title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[5:]
	t, _ := template.ParseFiles("add.html")
	p := map[string]string{"Title": title}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/save/", saveHanler)
	http.ListenAndServe(":8080", nil)
}
