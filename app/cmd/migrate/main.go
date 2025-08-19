package main

import (
	"log"

	"job-hunting-service-management-backend/app/infrastructure/migrate"
)

func main() {
	log.Println("Running database migration...")

	if err := migrate.Run(); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully!")
}
