package storage

import "urlshortner/db"

func SaveUrl(code string, url string) {
	_, err := db.DB.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
		code,
		url,
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
