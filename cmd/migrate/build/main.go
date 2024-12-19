package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"         // PostgreSQL драйвер // Импорты миграций (из папки migrations)
	"github.com/pressly/goose/v3" // Goose
)

func main() {
	// Подключаем флаги
	upFlag := flag.Bool("up", false, "Apply all migrations (up)")
	downFlag := flag.Bool("down", false, "Rollback the last migration (down)")
	statusFlag := flag.Bool("status", false, "Print the status of migrations")
	flag.Parse()

	// Подключаемся к базе данных
	dsn := "postgres://goapi:secret@localhost:5432/goapi?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Устанавливаем dialect
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	// Путь к папке с миграциями
	migrationsDir := "./migrations"

	// Выполняем действие в зависимости от флага
	switch {
	case *upFlag:
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully!")
	case *downFlag:
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
		log.Println("Migration rolled back successfully!")
	case *statusFlag:
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatalf("Failed to get migration status: %v", err)
		}
	default:
		log.Println("Please specify an action: --up, --down, or --status")
		os.Exit(1)
	}
}
