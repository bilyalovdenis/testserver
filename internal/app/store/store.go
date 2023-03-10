package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


type Store struct{
	config *Config
	db *sql.DB
	productRepository *ProductRepository
}

func New (config *Config) *Store{
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error{
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil{
		return err
	}
	if err := db.Ping(); err != nil{
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close(){
	s.db.Close() 
}
func (s *Store) Product() *ProductRepository{
	if (s.productRepository != nil){
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}
	return s.productRepository
}