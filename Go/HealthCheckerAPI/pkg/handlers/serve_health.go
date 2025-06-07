package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"HealthCheckerAPI/pkg/db"
)

func GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	healthStatus, err := db.GetServerHealthStatus()
	if err != nil {
		http.Error(w, "Failed to retrieve health status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStatus)
}

func GetHealthStatusByServerID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Server ID is missing", http.StatusBadRequest)
		return
	}
	serverID := parts[2]

	healthStatus, err := db.GetServerHealthStatusByID(serverID)
	if err != nil {
		log.Printf("Error fetching health status for server %s: %v", serverID, err)
		http.Error(w, "Failed to retrieve health status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStatus)
}
