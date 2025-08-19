package migrate

import (
	"fmt"
	"log"

	"job-hunting-service-management-backend/app/infrastructure/db"
	"job-hunting-service-management-backend/app/internal/entity"

	"gorm.io/gorm"
)

func Run() error {
	// DB接続
	database, err := db.NewDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer func() {
		if err := db.Close(database); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	log.Println("Starting database migration...")

	entities := []interface{}{
		&entity.SampleUser{},
		&entity.User{},
		// 新しいエンティティはここに追加
	}

	// 各エンティティのマイグレーション状況をチェック
	var migrateTargets []interface{}
	for _, ent := range entities {
		if !database.Migrator().HasTable(ent) {
			migrateTargets = append(migrateTargets, ent)
			log.Printf("📝 New table detected for migration: %T", ent)
		}
	}

	if len(migrateTargets) == 0 {
		log.Println("📋 No migration targets found. All tables are up to date.")
		return nil
	}

	// マイグレーション実行
	log.Printf("🚀 Migrating %d entities...", len(migrateTargets))
	err = database.AutoMigrate(migrateTargets...)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully!")

	if err := verifyMigration(database, entities); err != nil {
		log.Printf("Migration verification failed: %v", err)
		return err
	} else {
		log.Println("Migration verification passed!")
	}

	return nil
}

func verifyMigration(database *gorm.DB, entities []interface{}) error {
	for _, ent := range entities {
		if !database.Migrator().HasTable(ent) {
			return fmt.Errorf("table does not exist for entity: %T", ent)
		}
		log.Printf("✅ Table verified: %T", ent)
	}

	return nil
}
