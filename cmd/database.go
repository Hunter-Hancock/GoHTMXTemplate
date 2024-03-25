package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type DBConfig struct {
	user     string
	dbname   string
	password string
	dbtype   string
	host     string
}

func NewDBConfig(host, dbtype, dbname, user, password string) *DBConfig {
	return &DBConfig{user: user, dbname: dbname, password: password, dbtype: dbtype, host: host}
}

type DB interface {
	GetAllOrders() ([]*Order, error)
	CreateOrder(string) error
	DeleteOrder(int) error
}

type Store struct {
	db *sql.DB
}

func InitDB(config *DBConfig) error {
	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", config.host, config.dbname, config.user, config.password)
	db, err := sql.Open(config.dbtype, connectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store = &Store{db: db}

	return nil
}

func (s *Store) GetAllOrders() ([]*Order, error) {
	orders := []*Order{}

	rows, err := s.db.Query("SELECT * FROM Orders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		order := new(Order)
		if err := rows.Scan(&order.Method, &order.Id); err != nil {
			return nil, fmt.Errorf("Error scanning: %s", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Store) CreateOrder(method string) error {
	_, err := s.db.Exec("INSERT INTO Orders (Method) VALUES (@method)", sql.Named("Method", method))
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteOrder(id int) error {
	_, err := s.db.Exec("DELETE FROM Orders WHERE id = @id", sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}
