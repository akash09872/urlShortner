package storage

import (
	"time"
	"urlshortner/db"
)

func Clean() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		db.DB.Exec(
			"DELETE FROM urls WHERE expires_at < $1",
			time.Now(),
		)
	}
}
func SaveUrl(code string, url string) {
	expiry_time := time.Now().Add(6 * time.Hour)
	_, err := db.DB.Exec(
		"INSERT INTO urls (short_code, original_url,expires_at) VALUES ($1, $2, $3)",
		code,
		url,
		expiry_time,
	)
	if err != nil {
		panic(err)
	}
}

func GetURL(code string) string {
	var url string
	err := db.DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code= $1",
		code,
	).Scan(&url)
	if err != nil {
		panic(err)
	}
	return url
}
