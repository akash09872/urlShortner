package storage

import (
	"time"
	"urlshortner/db"
	"urlshortner/model"
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
func SaveUrl(code string, url string, userId int) {
	expiry_time := time.Now().Add(48 * time.Hour)
	_, err := db.DB.Exec(
		"INSERT INTO urls (short_code, original_url,expires_at,user_id) VALUES ($1, $2, $3, $4)",
		code,
		url,
		expiry_time,
		userId,
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
		return ""
	}
	return url
}

func GetUserURLs(userId int) []model.URLResponse {
	rows, _ := db.DB.Query(
		"SELECT short_code,original_url FROM urls WHERE user_id=$1",
		userId,
	)
	defer rows.Close()
	var urls []model.URLResponse
	for rows.Next() {
		var url model.URLResponse
		err := rows.Scan(
			&url.ShortCode,
			&url.OriginalURL,
		)
		if err != nil {
			return nil
		}
		urls = append(urls, url)
	}
	return urls
}
