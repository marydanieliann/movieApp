package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"log"
	"movieProject/model"
	"net/http"
)

func ListmoviesHandler(c *gin.Context) {
	movies1 := model.ListMoviesHandler()
	if movies1 == nil || len(movies1) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, movies1)
	}
}

func CreatemoviesHandler(c *gin.Context) {
	var mov model.Movie

	if err := c.BindJSON(&mov); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		model.CreateMoviesHandler(mov)
		c.IndentedJSON(http.StatusCreated, mov)
	}
}

func Getmoviesbyid(c *gin.Context) {
	id := c.Param("id")

	mov := model.GetMoviesbyID(id)

	if mov == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, mov)
	}
}

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/register", model.RegisterHandler)
	http.HandleFunc("/login", model.LoginHandler)
	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		movies := model.ListMoviesHandler()
		if movies != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movies)
		} else {
			http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		}
	})

	r.POST("/movies", ListmoviesHandler)
	r.GET("/movies/:id", Getmoviesbyid)
	log.Println("Server is starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
