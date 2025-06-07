package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type HealthData struct {
	ServerID    string `json:"server_id"`
	CPUUsage    int    `json:"cpu_usage"`
	MemoryUsage int    `json:"memory_usage"`
	DiskUsage   int    `json:"disk_usage"`
	Timestamp   string `json:"timestamp"`
}

func collectData(serverID string) HealthData {
	cpuPercent, _ := cpu.Percent(0, false)
	memStats, _ := mem.VirtualMemory()
	diskStats, _ := disk.Usage("/")

	data := HealthData{
		ServerID:    serverID,
		CPUUsage:    int(cpuPercent[0]),
		MemoryUsage: int(memStats.UsedPercent),
		DiskUsage:   int(diskStats.UsedPercent),
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	log.Printf("Collected Data - CPU: %d%%, Mem: %d%%, Disk: %d%%", data.CPUUsage, data.MemoryUsage, data.DiskUsage)
	return data
}

func sendData(apiURL string, data HealthData) error {
	payload, _ := json.Marshal(data)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error sending data: %v", err)
		return err
	}
	defer resp.Body.Close()

	log.Printf("Data sent. Status: %s", resp.Status)
	return nil
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func main() {
	loadEnv()

	apiURL := os.Getenv("API_URL")
	serverID := os.Getenv("SERVER_ID")

	if apiURL == "" || serverID == "" {
		log.Fatal("API_URL or SERVER_ID not set")
	}

	for {
		data := collectData(serverID)
		sendData(apiURL, data)
		time.Sleep(30 * time.Second)
	}
}
