CREATE TABLE IF NOT EXISTS server_health (
	id SERIAL PRIMARY KEY,
	server_id VARCHAR(255) NOT NULL,
	cpu_usage INT,
	memory_usage INT,
	disk_usage INT,
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
