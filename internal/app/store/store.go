package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // ...
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	if len(s.config.DatabaseURL) < 5 {
		log.Println("empry config")
		return nil
	}
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=restapi_test sslmode=disable")
	if err != nil {
		log.Println(err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {

	s.db.Close()

}

func (s *Store) User() *UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository

}
