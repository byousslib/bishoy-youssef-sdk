# Designing the SDK:

- Most of the code was written using TDD.
- The SDK is designed with a config struct to allow the user to configure any instance of
  the API as well as the API key.
- To make it easier for the user to reach the default API, I added a defaultServer file
  with the default URL. This also allows for easy modification of the default URL
- The main file is only concerned with the initialization of the SDK and main client
- for each of the APIs, a new file is to be added. In this example, I added the movie.go
  file for all the movie related Data Structure and functions.
- httpWrapperClient.go is a wrapper for the go http client that automatically uses the
  base url as well as the token

## Testing the SDK:

The tests are split into 3 categories:
- SDK tests: the basic unit tests of the SDK, assuming a mock API
- API tests: tests to check the integrity of the API and assert the expected outputs
- Integration tests: tests that run the SDK as a package and make sure it is working on an
  integration level.

## Documentation:

The documentation is done based on GoDocs, which is the Go library for documentation.

In order to read the documentation of the project you can either go to the GoDocs Website
[here](https://godocs.io/)

*OR*

Install `godoc` locally by running the command:
`go install -v golang.org/x/tools/cmd/godoc@latest`

Then run `godoc -http=:6060 &` to be able to view the documentation locally (on port 6060)

## P.S.

- In order not to clutter my personal github account, I have created this new account for the purpose of the task
- I had to squash all the commits into one as the initial commits (when the repo was private) had the API hard coded
