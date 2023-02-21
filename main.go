package lotrApiSdk

// This type is main object that is returned by the SDK, and it encapsulates
// All APIs that are currently available.
type Client struct {
	config      Config
	httpWrapper httpClient
}

// The Config struct is used to configure the SDK to talk to any instance of the API
// it allows the user to choose between the default API or their own instance of the API
type Config struct {
	// The Base URL of the LOTR API
	// If this field is left empty, the default URL of the API will be used
	// "https://the-one-api.dev/v2/"
	ApiUrl string

	// The API key for the API
	// This field is mandatory since all the calls to the API are
	// Authenticated
	ApiKey string
}

// This Function initializes the connectection to API based on the passed config struct
// And returns a client object to be used to communicate with the API
func Init(c Config) Client {
	if c.ApiUrl == "" {
		c.ApiUrl = defaultServerApi
	}
	wrapper := httpClient{
		apiToken: c.ApiKey,
		baseUrl:  c.ApiUrl,
	}

	return Client{config: c, httpWrapper: wrapper}
}
