package pkg

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// Config provides options to establish connection to MySQL db
type Config struct {
	ConnectionURL string `json:"connectionUrl"`
}

// SQLiteClient a way to work with MySQL database
type SQLiteClient struct {
	Config *Config
	Db     *sqlx.DB
}

// NewSQLiteClient - initialize new struct with config
func NewSQLiteClient(config *Config) *SQLiteClient {
	return &SQLiteClient{
		Config: config,
	}
}

// Open new MySQL connection using passed to New func config
func (m *SQLiteClient) Open() error {
	db, err := sqlx.Connect("sqlite", m.Config.ConnectionURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	m.Db = db

	return nil
}

// Close current MySQL connection
func (m *SQLiteClient) Close() error {
	return m.Db.Close()
}
