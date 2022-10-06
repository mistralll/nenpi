package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/mistralll/nenpi/refueling"
	"github.com/mistralll/nenpi/vehicle"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := vehicle.LoadVehicleInf(title)
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

func listHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/list.html")
	list, err := vehicle.GetVehicleList()
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, list)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	refuel := refueling.HttpReqToRefuel(r)
	refuel.SaveRefuel(title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func saveIconHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[10:]
	vehicle.SaveIcon(r, title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getVehicleInfHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[15:]
	fmt.Println(title)
}

func main() {
	http.HandleFunc("/view/", viewHandler) // 車体ごとの情報
	http.HandleFunc("/add/", addHandler)   // 給油の追加ページ
	http.HandleFunc("/edit/", editHandler) // 画像をアップロードするページ
	http.HandleFunc("/list/", listHandler) // 車体一覧を表示

	http.HandleFunc("/save/", saveHandler)          // 給油情報の保存
	http.HandleFunc("/saveicon/", saveIconHandler) // 画像の保存

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("data/"))))

	http.ListenAndServe(":8080", nil)
}
