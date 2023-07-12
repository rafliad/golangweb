package handler

import (
	"golangWeb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error sedang terjadi", http.StatusInternalServerError)
		return
	}
	data := []entity.Product{
		{ID: 1, Name: "Xiaomi", Price: 3000000, Stock: 2},
		{ID: 2, Name: "Huawei", Price: 9000000, Stock: 7},
		{ID: 3, Name: "Samsung", Price: 12000000, Stock: 14},
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error sedang terjadi", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, belajar golang"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idAngka, err := strconv.Atoi(id)
	if err != nil || idAngka < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error sedang terjadi", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"content": idAngka,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error sedang terjadi", http.StatusInternalServerError)
		return
	}
}
