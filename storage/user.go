package storage

import (
	"urlshortner/db"
)

func SaveUser(username string, hashPassword string) error {
	_, err := db.DB.Query(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		username,
		hashPassword,
	)
	return err
}

func GetUserId(username string) (int, error) {
	var id int
	err := db.DB.QueryRow(
		"SELECT id FROM users WHERE username=$1",
		username,
	).Scan(&id)
	return id, err
}
