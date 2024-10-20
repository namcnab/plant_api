package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	dbConn "github.com/namcnab/plant_api/internal/database"
	"github.com/namcnab/plant_api/internal/handler"
	"github.com/namcnab/plant_api/internal/model"
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
        response := model.Response{}

        glossaryResp, err := handler.GetGlossary(db)
        if err != nil {
            response.Code = http.StatusInternalServerError
            response.Message = err.Error()
        }

        jsonResp, err := json.Marshal(glossaryResp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResp)

    })

    http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {  
        response := model.Response{}

        entry := model.Glossary{
            Term: removeNewlines(capitalizeWords(r.FormValue("term"))),
            Definition: removeNewlines(capitalizeSentence(r.FormValue("definition"))),
        }

        err := handler.CreateGlossaryEntry(db, entry)
        if err != nil {
            response.Code = http.StatusInternalServerError
            response.Message = "Failed to add glossary entry"
        } else {
            response.Code = http.StatusOK
            response.Message = "Glossary entry added successfully"
        }

        jsonResp, err := json.Marshal(response)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResp)
    })

    http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
        response := model.Response{}

        entry := model.Glossary{
            Term: removeNewlines(capitalizeWords(r.FormValue("term"))),
            Definition: removeNewlines(capitalizeSentence(r.FormValue("definition"))),
        }

        err := handler.UpdateGlossaryEntry(db, entry)
        if err != nil {
            response.Code = http.StatusInternalServerError
            response.Message = err.Error()
        } else {
            response.Code = http.StatusOK
            response.Message = entry.Term + " was updated successfully"
        }

        jsonResp, err := json.Marshal(response)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResp)
    })

    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
        response := model.Response{}


        err := handler.DeleteGlossaryTerm(db, removeNewlines(capitalizeWords(r.FormValue("term"))))
        if err != nil {
            response.Code = http.StatusInternalServerError
            response.Message = err.Error()
        } else {
            response.Code = http.StatusOK
            response.Message = "Glossary entry deleted successfully"
        }

        jsonResp, err := json.Marshal(response)

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

func capitalizeWords(s string) string {
    runes := []rune(s)
    capitalizeNext := true

    for i, r := range runes {
        if capitalizeNext {
            runes[i] = unicode.ToUpper(r)
            capitalizeNext = false
        }
        if unicode.IsSpace(r) {
            capitalizeNext = true
        }
    }

    return string(runes)
}

func capitalizeSentence(s string) string {
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}

func removeNewlines(s string) string {
    return strings.ReplaceAll(s, "\n", "")
}