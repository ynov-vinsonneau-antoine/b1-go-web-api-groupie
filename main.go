package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	Members      []string `json:"members"`
}

func getArtists() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur :", err)
		return nil
	}
	defer resp.Body.Close()

	var artists []Artist
	json.NewDecoder(resp.Body).Decode(&artists)
	return artists
}

func accueil(w http.ResponseWriter, r *http.Request) {
	artists := getArtists()

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
func main() {
	http.HandleFunc("/", accueil)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}