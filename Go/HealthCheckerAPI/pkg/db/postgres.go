package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var connStr = "host=localhost user=postgres dbname=health_checker_api password=mypassword sslmode=disable"

func InitDB() error {

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Successfully connected to PostgreSQL")
	return nil
}

func GetServerHealthStatus() ([]HealthStatus, error) {
	query := `
		SELECT server_id, cpu_usage, memory_usage, disk_usage, timestamp
		FROM server_health 
		ORDER BY timestamp DESC 
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var status []HealthStatus
	for rows.Next() {
		var s HealthStatus
		if err := rows.Scan(&s.ServerID, &s.CPUUsage, &s.MemoryUsage, &s.DiskUsage, &s.Timestamp); err != nil {
			return nil, err
		}
		status = append(status, s)
	}
	return status, nil
}

func GetServerHealthStatusByID(serverID string) ([]HealthStatus, error) {
	query := `
		SELECT server_id, cpu_usage, memory_usage, disk_usage, timestamp
		FROM server_health 
		WHERE id = $1
		ORDER BY timestamp DESC 
		LIMIT 1
	`
	rows, err := DB.Query(query, serverID)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var status []HealthStatus

	for rows.Next() {
		var s HealthStatus
		if err := rows.Scan(&s.ServerID, &s.CPUUsage, &s.MemoryUsage, &s.DiskUsage, &s.Timestamp); err != nil {
			return nil, err
		}
		status = append(status, s)
	}

	if len(status) == 0 {
		return nil, fmt.Errorf("server with ID %s not found", serverID)
	}

	return status, nil
}

func StoreHealthCheckData(data AgentHealthData) error {
	query := `
		INSERT INTO server_health (server_id, cpu_usage, memory_usage, disk_usage, timestamp)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := DB.Exec(query, data.ServerID, data.CPUUsage, data.MemoryUsage, data.DiskUsage, data.Timestamp)
	if err != nil {
		log.Printf("Error storing health data: %v", err)
		return err
	}

	return nil
}

type HealthStatus struct {
	ServerID    string
	CPUUsage    int
	MemoryUsage int
	DiskUsage   int
	Timestamp   string
}

type AgentHealthData struct {
	ServerID    string `json:"server_id"`
	CPUUsage    int    `json:"cpu_usage"`
	MemoryUsage int    `json:"memory_usage"`
	DiskUsage   int    `json:"disk_usage"`
	Timestamp   string `json:"timestamp"`
}
