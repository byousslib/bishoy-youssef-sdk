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


## Self-Assessment
1- Have you manually tested the SDK?
  - I have created an automated test suite (the integration tests) that tests the utilization of the tool in a real end-to-end scenario.

2- Did you add a test suite? If so, how will we use it? If not, why?
  - I added 3 test suites, the instructions and uses of which are mentioned in the README file [here](https://github.com/byousslib/bishoy-youssef-sdk/blob/main/README.md#run-the-tests)

4- Do you feel this SDK makes it easier to interact with the API?
  - Yes, the SDK has very simple and intuitive methods to interact with the API. 

5- If you had more time, what else would you add?
  - If I were to do more, I would have added an optional pagination struct to the Movies Method. It would be an input to the SDK and also an output from it. When used as an input, it would instruct the SDK to bring a specific page and also the movies method would output that to show the user the pagination details.  the struct would look like:
  ```golang
    type Pagination struct {
      Limit  int
      Page   int
      Offset int
      Pages  int
    }
  ```

6- What would you change in your current SDK solution?
  - While the SDK handles the calls from the user well, it still lacks optimizations that would allow for greater scalability. That includes Async calls, keep alives, and queueing.

7 - On a scale of 1 to 10, how would you rate this SDK? (higher is better).
- Overall, I would give it 9.The code is TDD driven, and as much as possible structured.The SDK simplifies the API as much as possibleThe SDK has godocs documentation, and is available using go install.It still lacks some performance optimizations, and functionalities.

8- Anything else we should keep in mind when we evaluate the project?
  - as much as possible I have added godocs, readme instructions and design documentation. They are all included in the library. I also added the answers to these questions in the `design.md` file for easier access.


## P.S.

- In order not to clutter my personal github account, I have created this new account for the purpose of the task
- I had to squash all the commits into one as the initial commits (when the repo was private) had the API hard coded
