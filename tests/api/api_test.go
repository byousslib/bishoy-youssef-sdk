package api_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	sdk "github.com/byousslib/bishoy-youssef-sdk"
	httpexpect "github.com/gavv/httpexpect/v2"
)

var (
	api_url = "https://the-one-api.dev/v2/"
	token   string
)

type testMovies struct {
	Docs   []sdk.Movie `json:"docs"`
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Page   int         `json:"page"`
	Pages  int         `json:"pages"`
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

func TestApiUnauthorizedWithoutAPIKey(t *testing.T) {
	e := httpexpect.Default(t, api_url)

	e.GET("/movie").
		Expect().
		Status(http.StatusUnauthorized)
}

func TestAPIAuthentication(t *testing.T) {
	e := httpexpect.Default(t, api_url)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	auth.GET("/movie").
		Expect().
		Status(http.StatusOK)
}

func TestMovieAPIStructure(t *testing.T) {
	schema := `{
		"type": "object",
		"properties": {
			"docs": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"_id": {
							"type": "string"
						},
						"academyAwardNominations": {
							"type": "number"
						},
						"academyAwardWins": {
							"type": "number"
						},
						"boxOfficeRevenueInMillions": {
							"type": "number"
						},
						"budgetInMillions": {
							"type": "number"
						},
						"name": {
							"type": "string"
						},
						"rottenTomatoesScore": {
							"type": "number"
						},
						"runtimeInMinutes": {
							"type": "number"
						}
					}
				}
			}
		}
	}`
	e := httpexpect.Default(t, api_url)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	movies := auth.GET("/movie").
		Expect().
		Status(http.StatusOK).
		JSON()

	movies.Schema(schema)
}

func TestMovieCompatible(t *testing.T) {
	e := httpexpect.Default(t, api_url)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	var testMoviesInstance testMovies
	auth.GET("/movie").
		Expect().
		Status(http.StatusOK).
		JSON().
		Decode(&testMoviesInstance)
}
