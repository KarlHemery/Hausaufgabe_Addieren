package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"Hausaufgabe_Addieren/models"
	"Hausaufgabe_Addieren/services"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, "Error loading HTML file", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Failed to decode request: %v", err)
		return
	}

	db := services.NewDB(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	sum, err := db.CalculateSum(req.Num1, req.Num2)
	if err != nil {
		http.Error(w, "Failed to calculate sum", http.StatusInternalServerError)
		log.Printf("Failed to calculate sum: %v", err)
		return
	}

	res := models.AddResponse{Sum: sum}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}
