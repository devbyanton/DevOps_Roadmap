package handlers

import (
	"HealthCheckerAPI/pkg/db"
	"encoding/json"
	"net/http"
)

type AgentHealthData struct {
	ServerID    string `json:"server_id"`
	CPUUsage    int    `json:"cpu_usage"`
	MemoryUsage int    `json:"memory_usage"`
	DiskUsage   int    `json:"disk_usage"`
	Timestamp   string `json:"timestamp"`
}

func PostHealthStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data AgentHealthData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if data.ServerID == "" || data.Timestamp == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	dbData := db.AgentHealthData{
		ServerID:    data.ServerID,
		CPUUsage:    data.CPUUsage,
		MemoryUsage: data.MemoryUsage,
		DiskUsage:   data.DiskUsage,
		Timestamp:   data.Timestamp,
	}
	err = db.StoreHealthCheckData(dbData)
	if err != nil {
		http.Error(w, "Failed to store health data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health data received and stored"))
}
