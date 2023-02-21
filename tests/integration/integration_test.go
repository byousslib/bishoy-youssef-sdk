package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	lotrSdk "github.com/byousslib/bishoy-youssef-sdk"
	"github.com/stretchr/testify/assert"
)

var (
	api_url = "https://the-one-api.dev/v2/"
	token   string
)

type TotalMovies struct {
	Total int `json:"total"`
}

func init() {
	token = os.Getenv("TOKEN")
}

func TestMain(m *testing.M) {
	if token == "" {
		log.Fatal("token not defined")
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func countCurrentMovies() int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", api_url+"movie", nil)
	req.Header.Add("Authorization", "bearer "+token)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result TotalMovies
	json.Unmarshal(body, &result)

	return result.Total
}

func TestHelloWorld(t *testing.T) {
	config := lotrSdk.Config{ApiKey: token}

	lotrClient := lotrSdk.Init(config)

	movies, err := lotrClient.Movies()

	assert.Nil(t, err, "SDK should not return an error")
	assert.NotEmpty(t, movies, "The SDK should return data")
	assert.Equal(t, countCurrentMovies(), len(movies))
}
