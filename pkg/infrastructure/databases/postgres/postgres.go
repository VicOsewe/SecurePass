package postgres

import "database/sql"

type SecurePassDB struct {
	DB *sql.DB
}

func NewSecurePassDB() *SecurePassDB {

}
