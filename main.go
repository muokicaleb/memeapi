package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var listAllMemes []string = listMemes()

var memeServerBaseURl string = loadServerUrl()

func main() {
	router := gin.Default()
	router.GET("/api/v3/meme", getMeme)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getMeme(c *gin.Context) {

	memeUrl := memeServerBaseURl + listAllMemes[randomNumber()]
	c.String(http.StatusOK, string(memeUrl))
}

func loadServerUrl() string {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv("MEME_SERVER")

}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(listAllMemes)
	return rand.Intn(max-min+1) + min

}
func listMemes() []string {
	var files []string

	fileInfo, err := ioutil.ReadDir("/memestore")

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files

}
