package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config represents database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Connection represents database connection
type Connection struct {
	DB *sql.DB
}

// NewConnection creates a new database connection
func NewConnection(config Config) (*Connection, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Connection{DB: db}, nil
}

// Close closes the database connection
func (c *Connection) Close() error {
	return c.DB.Close()
}

// Migrate runs database migrations
func (c *Connection) Migrate() error {
	// Add your entities here for auto-migration
	// return c.DB.AutoMigrate(&entity.User{})
	return nil
}
