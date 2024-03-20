package main

import (
	"encoding/json"
	"net/http"
)

type Animal struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}


func HandleListAnimals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	animals := []Animal{
		{Name: "Dog", ImageUrl: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcQgByBT5IiAT_a2x9pUVb4VMoOrlzHH7Jrzj-HB5jzHlR4lNLMS"},
		{Name: "Cat", ImageUrl: "https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg"},
		{Name: "Bird", ImageUrl: "https://www.birds.cornell.edu/home/wp-content/uploads/2023/09/334289821-Baltimore_Oriole-Matthew_Plante.jpg"},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(animals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/api/animals", HandleListAnimals)

	http.ListenAndServe(":6969", nil)
}

