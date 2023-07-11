package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]product, error) {
	query := "SELECT id, name, quantity, price FROM products"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	products := []product{}

	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {

	query := fmt.Sprintf("SELECT name, quantity, price FROM products WHERE id=%v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)

	if err != nil {
		return err
	}
	return nil
}

func (p *product) createProduct(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES('%v', '%v', '%v')", p.Name, p.Quantity, p.Price)

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)

	return nil
}
