package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// "Person type" (tipo um objeto)
type Company struct {
    ID        string   `json:"id,omitempty"`
    Name string   `json:"Name,omitempty"`
    Cnpj  string   `json:"cnpj,omitempty"`
}

var company []Company

// GetPeople mostra todos os contatos da variável people
func GetCampany(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(company)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range company {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Company{})
}

// função principal para executar a api
func main() {
    router := mux.NewRouter()
    company = append(company, Company{ID: "1", Name: "Google", Cnpj: "00.000.0001/20"})
    company = append(company, Company{ID: "2", Name: "IBM", Cnpj: "00.000.0001/30"})

    router.HandleFunc("/company", GetCampany).Methods("GET")
    router.HandleFunc("/company/{id}", GetPerson).Methods("GET")
    log.Fatal(http.ListenAndServe(":1337", router))
}