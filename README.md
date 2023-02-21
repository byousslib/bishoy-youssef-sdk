# The Lord of the Rings Helper SDK

The lotr SDK lets you connect to your favorite LOTR "One API Rule them ALL" API.

for more information about the API, please visit: [The One
API](https://the-one-api.dev/documentation)

## Installation:

The recommended installation method is using [Go Modules](https://go.dev/ref/mod#go-get)

In order to initialize a Go Application, make sure to install the latest version of [GO](https://go.dev/).

Then, run the below command to start a new project:
`go mod init [YOUR_PROJECT_NAME]`

Once you have an initialized Go Application, you add the SDK to your go modules.
`go get github.com/byousslib/bishoy-youssef-sdk`

Make sure to obtain an API Key by signing up [here](https://the-one-api.dev/sign-up)

These are all the steps needed to use the SDK.

Here's a sample code for getting all the available movies in the database:

```golang
package main

import (
	"fmt"
	"log"

	lotrSdk "github.com/byousslib/bishoy-youssef-sdk"
)

func main() {
	config := lotrSdk.Config{ApiKey: token}

	lotrClient := lotrSdk.Init(config)

	movies, err := lotrClient.Movies()
	if err != nil {
		log.Fatal(err)
	}

	//this will print all the movies in the database
	fmt.Println(movies)
}
```
## Read the Docs:

In order to read the documentation of the project you can either go to the GoDocs Website
[here](https://godocs.io/)

*OR*

Install `godoc` locally by running the command:
`go install -v golang.org/x/tools/cmd/godoc@latest`

Then run `godoc -http=:6060 &` to be able to view the documentation locally (on port 6060)

## Run the tests:

The simplest way to run all the tests for the SDK is to run `TOKEN=[YOUR_TOKEN] go test -v
./...` in the root folder of the repository

The tests are split into 3 categories:
- SDK tests: the basic unit tests of the SDK, assuming a mock API
- API tests: tests to check the integrity of the API and assert the expected outputs
- Integration tests: tests that run the SDK as a package and make sure it is working on an
  integration level.










TODO:
- [x] Implement http test to make sure the API is working as expected
- [x] Implement SDK import test
- [x] Implement SDK initialization test
- [x] Implement Auth
- [ ] Implement SDK getMovie
  - currently need to make sure the SDK is compatible with the real api

Features:
- Config:
  - API KEY
  - URL
- Login //not needed right now //can be just for hand shakes and keep alive
- GO Docs for easy usage
- Readme file to show easy steps to get started
- Movie Struct
- Movies Struct ???
- quotes struct
