package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	custSql "github.com/volodymyrzuyev/marketizer/internal/database/sql"
)

// Service represents a service that interacts with a database.
type Service interface {
	GetUser(email string) (custSql.User, error)
	AddUser(email, password, name string) error
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	q  *custSql.Queries
	db *sql.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	dbInstance = &service{
		q:  custSql.New(db),
		db: db,
	}
	dbInstance.q.Create_table1(context.TODO())
	dbInstance.q.Create_table2(context.TODO())
	dbInstance.q.Create_table3(context.TODO())
	return dbInstance
}

func (s *service) GetUser(email string) (custSql.User, error) {
	return s.q.Get_User(context.TODO(), email)
}

func (s *service) AddUser(email, password, name string) error {
	args := custSql.Add_UserParams{
		Name:     sql.NullString{name, name != ""},
		Email:    email,
		Password: sql.NullString{password, password != ""},
	}
	return s.q.Add_User(context.TODO(), args)
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}
