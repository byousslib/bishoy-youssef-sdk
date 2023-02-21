package sdk_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	sdk "github.com/byousslib/bishoy-youssef-sdk"
)

var (
	mockMovies []sdk.Movie = []sdk.Movie{{Name: "Movie-name"}}
)

type apiReturn struct {
	Docs []sdk.Movie
}

func moviesHandler(t *testing.T) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer test", "Authorization Header should be added based on the config")

		MoviesReturn := apiReturn{mockMovies}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(MoviesReturn)
	})

	return mux
}

func mockApiServer(t *testing.T) *httptest.Server {
	handler := moviesHandler(t)
	server := httptest.NewServer(handler)

	return server
}

func TestSDKImport(t *testing.T) {
	server := mockApiServer(t)

	defer server.Close()
	config := sdk.Config{
		ApiUrl: server.URL,
		ApiKey: "test",
	}
	client := sdk.Init(config)
	movies, err := client.Movies()
	assert.Nil(t, err, "No Errors should be returned")
	assert.Equal(t, mockMovies, movies, "The SDK should return the same movies as the API")
}
