package lotrApiSdk

import (
	"encoding/json"
	"io/ioutil"
)

// use this stuct to be able to interact with the API, or use it as the API return.
// The fields that are present here are:
type Movie struct {
	// The Database ID of the movie, can be used to retrieve movies using the movie endpoint
	// (Still to be implemented)
	ID string `json:"_id"`

	// This represents the name of the movie
	Name string `json:"name"`

	// How Long is the movie, in minutes
	RuntimeInMinutes int `json:"runtimeInMinutes"`

	// The Movie budget in millions
	BudgetInMillions int `json:"budgetInMillions"`

	// How Much money did the movie make in millions
	BoxOfficeRevenueInMillions float32 `json:"boxOfficeRevenueInMillions"`

	// How many Academy Award Nominations did the movie get
	AcademyAwardNominations int `json:"academyAwardNominations"`

	// Out of the nominations, how many did the movie win
	AcademyAwardWins int `json:"academyAwardWins"`

	// The movie score on Rotten Tomatoes
	RottenTomatoesScore float32 `json:"rottenTomatoesScore"`
}

type apiMovieResult struct {
	Movies []Movie `json:"docs"`
}

// This is the main API of the SDK, it return all the movies that are available in the database.
// The return is an array of Movies struct
func (c *Client) Movies() ([]Movie, error) {
	resp, err := c.httpWrapper.Get("/movie")
	if err != nil {
		return []Movie{}, err
	}

	defer resp.Body.Close()
	var resultMovies apiMovieResult
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Movie{}, err
	}
	if err := json.Unmarshal(body, &resultMovies); err != nil {
		return []Movie{}, err
	}
	return resultMovies.Movies, nil
}
