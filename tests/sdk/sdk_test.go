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
	mockMovies []sdk.Movie = []sdk.Movie{{ID: "some-id", Name: "Movie-name"}}
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

func movieHandler(t *testing.T) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/movie/"+mockMovies[0].ID, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer test", "Authorization Header should be added based on the config")

		MoviesReturn := apiReturn{mockMovies}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(MoviesReturn)
	})

	return mux
}
func mockApiServer(handler http.Handler) *httptest.Server {
	server := httptest.NewServer(handler)

	return server
}

func TestSDKSingleMovie(t *testing.T) {
	handler := movieHandler(t)
	server := mockApiServer(handler)

	defer server.Close()
	config := sdk.Config{
		ApiUrl: server.URL,
		ApiKey: "test",
	}
	client := sdk.Init(config)
	movie, err := client.Movie(mockMovies[0].ID)
	assert.Nil(t, err, "No Errors should be returned")
	assert.Equal(t, mockMovies[0].Name, movie.Name, "The SDK should return the same movies as the API")
}

func TestSDKMovies(t *testing.T) {
	handler := moviesHandler(t)
	server := mockApiServer(handler)

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
