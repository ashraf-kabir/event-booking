package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

// InitMysqlConn returns a singleton DB connection
func InitMysqlConnection() *sql.DB {
	dbOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println(".env file not found. Using system environment variables.")
		}

		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		name := os.Getenv("DB_NAME")

		if user == "" || pass == "" || host == "" || name == "" {
			log.Fatal("Missing database environment variables")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, name)
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}

		if err := conn.Ping(); err != nil {
			log.Fatalf("Failed to ping DB: %v", err)
		}

		log.Println("MySQL connection established")
		db = conn
	})

	return db
}

func CloseMysqlConnection() {
	if db != nil {
		_ = db.Close()
		log.Println("MySQL connection closed")
	}
}

// CreateMysqlTables initializes the database schema
func CreateMysqlTables() {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
        user_id INTEGER,
        name VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        location VARCHAR(255) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    )`

    createRegTable := `
    CREATE TABLE IF NOT EXISTS registrations (
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
        event_id INTEGER,
        user_id INTEGER,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY(event_id) REFERENCES events(id) ON DELETE CASCADE,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    )`

    _, err := db.Exec(createUsersTable)
    if err != nil {
        log.Fatalf("Could not create users table: %v", err)
    }

    _, err = db.Exec(createEventsTable)
    if err != nil {
        log.Fatalf("Could not create events table: %v", err)
    }

    _, err = db.Exec(createRegTable)
    if err != nil {
        log.Fatalf("Could not create registrations table: %v", err)
    }

    log.Println("Database tables created successfully")
}