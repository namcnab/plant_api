package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbConn "github.com/namcnab/plant_api/internal/database"
	"github.com/namcnab/plant_api/internal/handler"
)

func main() {
    db, err  := dbConn.InitializeDB()

    if err != nil {
        fmt.Printf("Error initializing database: %s\n", err)
        return
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the Plant API!")
    })

    http.HandleFunc("/glossary", func(w http.ResponseWriter, r *http.Request) {
        
        glossaryResp, err := handler.GetGlossary(db)
        if err != nil {
            fmt.Fprintf(w, "Error getting glossary: %s", err)
            return
        }

        jsonResp, err := json.Marshal(glossaryResp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResp)

    })

    fmt.Println("Server starting on localhost:8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}