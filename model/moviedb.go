package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"movieProject/config"
)

var ConnStr = config.ConnectionString()

func ListMoviesHandler() []Movie {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM movies")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer results.Close()

	var movies []Movie
	for results.Next() {
		var mov Movie
		err = results.Scan(&mov.Title, &mov.Director, &mov.ID, &mov.UserID)
		if err != nil {
			fmt.Println("Scan Err", err.Error())
			return nil
		}
		movies = append(movies, mov)
	}
	return movies
}

func CreateMoviesHandler(movie Movie) {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO movies (id, title, director, user_id) VALUES ($1, $2, $3, $4)", movie.Title, movie.Director, movie.ID, movie.UserID)
	if err != nil {
		fmt.Println("Insert Err", err.Error())
	}
}

func GetMoviesbyID(id string) *Movie {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	mov := &Movie{}
	err = db.QueryRow("SELECT id, title, director, user_id FROM movies WHERE id = $1", id).Scan(&mov.Title, &mov.Director, &mov.ID, &mov.UserID)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	return mov
}
