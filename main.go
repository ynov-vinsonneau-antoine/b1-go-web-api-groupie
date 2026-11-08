package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Pokemon struct {
	ID   int
	Name string `json:"name"`
}

type ListePokemon struct {
	Results []Pokemon `json:"results"`
}

func getPokemons() []Pokemon {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=151")
	if err != nil {
		fmt.Println("Erreur :", err)
		return nil
	}
	defer resp.Body.Close()

	var liste ListePokemon
	json.NewDecoder(resp.Body).Decode(&liste)

	for i := range liste.Results {
		liste.Results[i].ID = i + 1
	}
	return liste.Results
}

func accueil(w http.ResponseWriter, r *http.Request) {
	pokemons := getPokemons()

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, pokemons)
}

func main() {
	http.HandleFunc("/", accueil)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
