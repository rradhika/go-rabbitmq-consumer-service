package utils

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("PG_DSN")
	if dsn == "" {
		log.Fatal("❌ PG_DSN is empty. Check your .env file")
	}

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	var dbname string
	err = DB.Get(&dbname, "SELECT current_database()")
	if err != nil {
		log.Fatalf("❌ Could not fetch current DB: %v", err)
	}
	log.Println("📌 Go is connected to database:", dbname)

	log.Println("✅ Connected to PostgreSQL")
	createSchema()
}

func createSchema() {
	schema := `
CREATE TABLE IF NOT EXISTS public.messages (
	id SERIAL PRIMARY KEY,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	log.Println("📐 Creating table with SQL:")
	log.Println(schema)

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("❌ Failed to create messages table: %v", err)
	} else {
		log.Println("✅ messages table is ready in public schema")
	}
}
