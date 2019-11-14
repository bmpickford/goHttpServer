package main

import "database/sql"

type datastore struct {
	db *sql.DB
}

func (d *datastore) SaveCustomer(username string) error {
	_, err := d.db.Exec("INSERT INTO customer (username) VALUES($1)", username)
	return err
}

func (d *datastore) GetCustomer(username string) (*customer, error) {
	var c *customer
	rows, err := d.db.Query("SELECT id, username FROM customer WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	err = rows.Scan(&c)
	return c, err
}
