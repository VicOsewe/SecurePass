package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VicOsewe/secure-pass/pkg/config"
)

type SecurePassDB struct {
	DB *sql.DB
}

func NewSecurePassDB(env *config.Env) *SecurePassDB {
	s := &SecurePassDB{
		DB: Init(env),
	}
	s.checkPreconditions()
	return s
}

func (db *SecurePassDB) checkPreconditions() {
	if db.DB == nil {
		log.Panicf("database has not been initialized")
	}
}

func Init(env *config.Env) *sql.DB {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	sqlDB, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open postgres connection: %v", err)
	}
	return sqlDB
}
