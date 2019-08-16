package main

func SaveCustomer(username string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO customer (username) VALUES($1)", username)
	return err
}
