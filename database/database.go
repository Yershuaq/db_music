package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "yourpassword"
		dbname   = "musicstreaming"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных")
}

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS playlists (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		user_id INT REFERENCES users(id),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS songs (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		artist VARCHAR(255),
		album VARCHAR(255),
		genre VARCHAR(255)
	);
	CREATE TABLE IF NOT EXISTS playlist_songs (
		playlist_id INT REFERENCES playlists(id),
		song_id INT REFERENCES songs(id),
		PRIMARY KEY (playlist_id, song_id)
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	fmt.Println("Миграция базы данных завершена")
}
